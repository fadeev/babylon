package keeper_test

import (
	"math/rand"
	"testing"

	"github.com/babylonchain/babylon/app"
	"github.com/babylonchain/babylon/testutil/datagen"
	checkpointingkeeper "github.com/babylonchain/babylon/x/checkpointing/keeper"
	"github.com/babylonchain/babylon/x/checkpointing/types"
	"github.com/babylonchain/babylon/x/epoching/testepoching"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func FuzzGetValidatorBlsKeySet(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		rand.Seed(seed)
		// a genesis validator is generated for setup
		helper := testepoching.NewHelper(t)
		ek := helper.EpochingKeeper
		ck := helper.App.CheckpointingKeeper
		querier := checkpointingkeeper.Querier{Keeper: ck}
		queryHelper := baseapp.NewQueryServerTestHelper(helper.Ctx, helper.App.InterfaceRegistry())
		types.RegisterQueryServer(queryHelper, querier)
		msgServer := checkpointingkeeper.NewMsgServerImpl(ck)
		genesisVal := ek.GetValidatorSet(helper.Ctx, 0)[0]
		genesisBLSPubkey, err := ck.GetBlsPubKey(helper.Ctx, genesisVal.Addr)
		require.NoError(t, err)

		// BeginBlock of block 1, and thus entering epoch 1
		ctx := helper.BeginBlock()
		epoch := ek.GetEpoch(ctx)
		require.Equal(t, uint64(1), epoch.EpochNumber)

		// 1. get validator BLS set when there's only a genesis validator
		valBlsSet := ck.GetValidatorBlsKeySet(ctx, epoch.EpochNumber)
		require.Equal(t, genesisVal.GetValAddressStr(), valBlsSet.ValSet[0].ValidatorAddress)
		require.True(t, genesisBLSPubkey.Equal(valBlsSet.ValSet[0].BlsPubKey))
		require.Equal(t, uint64(genesisVal.Power), valBlsSet.ValSet[0].VotingPower)

		// add n new validators via MsgWrappedCreateValidator
		n := rand.Intn(10) + 1
		addrs := app.AddTestAddrs(helper.App, helper.Ctx, n, sdk.NewInt(100000000))

		wcvMsgs := make([]*types.MsgWrappedCreateValidator, n)
		for i := 0; i < n; i++ {
			msg, err := buildMsgWrappedCreateValidator(addrs[i])
			require.NoError(t, err)
			wcvMsgs[i] = msg
			_, err = msgServer.WrappedCreateValidator(ctx, msg)
			require.NoError(t, err)
		}

		// EndBlock of block 1
		ctx = helper.EndBlock()

		// go to BeginBlock of block 11, and thus entering epoch 2
		for i := uint64(0); i < ek.GetParams(ctx).EpochInterval; i++ {
			ctx = helper.GenAndApplyEmptyBlock()
		}
		epoch = ek.GetEpoch(ctx)
		require.Equal(t, uint64(2), epoch.EpochNumber)

		// 2. get validator BLS set when there are n+1 validators
		epochNum := uint64(2)
		valBlsSet2 := ck.GetValidatorBlsKeySet(ctx, epochNum)
		expectedValSet := ek.GetValidatorSet(ctx, 2)
		for i, expectedVal := range expectedValSet {
			expectedBlsPubkey, err := ck.GetBlsPubKey(ctx, expectedVal.Addr)
			require.NoError(t, err)
			require.Equal(t, expectedVal.GetValAddressStr(), valBlsSet2.ValSet[i].ValidatorAddress)
			require.True(t, expectedBlsPubkey.Equal(valBlsSet2.ValSet[i].BlsPubKey))
			require.Equal(t, uint64(expectedVal.Power), valBlsSet2.ValSet[i].VotingPower)
		}
	})
}
