package keeper

import (
	"context"

	"github.com/babylonchain/babylon/x/btccheckpoint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	k Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

// TODO at some point add proper logging of error
// TODO emit some events for external consumers. Those should be probably emited
// at EndBlockerCallback
func (m msgServer) InsertBTCSpvProof(ctx context.Context, req *types.MsgInsertBTCSpvProof) (*types.MsgInsertBTCSpvProofResponse, error) {
	rawSubmission, err := types.ParseSubmission(req, m.k.GetPowLimit(), m.k.GetExpectedTag())

	if err != nil {
		return nil, types.ErrInvalidCheckpointProof.Wrap(err.Error())
	}

	// Get the SDK wrapped context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	submissionKey := rawSubmission.GetSubmissionKey()

	if m.k.SubmissionExists(sdkCtx, submissionKey) {
		return nil, types.ErrDuplicatedSubmission
	}

	newSubmissionOldestHeaderDepth, err := m.k.GetSubmissionBtcInfo(sdkCtx, submissionKey)

	if err != nil {
		return nil, types.ErrInvalidHeader.Wrap(err.Error())
	}

	// At this point:
	// - every proof of inclusion is valid i.e every transaction is proved to be
	// part of provided block and contains some OP_RETURN data
	// - header is proved to be part of the chain we know about through BTCLightClient
	// - this is new checkpoint submission
	// Verify if this is expected checkpoint
	err = m.k.checkpointingKeeper.VerifyCheckpoint(sdkCtx, rawSubmission.CheckpointData)

	if err != nil {
		return nil, err
	}

	// At this point we know this is a valid checkpoint for this epoch as this was validated
	// by checkpointing module
	epochNum := rawSubmission.CheckpointData.Epoch

	err = m.k.checkAncestors(sdkCtx, epochNum, newSubmissionOldestHeaderDepth)

	if err != nil {
		return nil, err
	}

	// construct TransactionInfo pair and the submission data
	txsInfo := make([]*types.TransactionInfo, len(submissionKey.Key))
	for i := range submissionKey.Key {
		// creating a per-iteration `txKey` variable rather than assigning it in the `for` statement
		// in order to prevent overwriting previous `txKey`
		// see https://github.com/golang/go/discussions/56010
		txKey := submissionKey.Key[i]
		txsInfo[i] = types.NewTransactionInfo(txKey, req.Proofs[i].BtcTransaction, req.Proofs[i].MerkleNodes)
	}
	submissionData := rawSubmission.GetSubmissionData(epochNum, txsInfo)

	// Everything is fine, save new checkpoint and update Epoch data
	err = m.k.addEpochSubmission(
		sdkCtx,
		epochNum,
		submissionKey,
		submissionData,
	)

	if err != nil {
		return nil, err
	}

	return &types.MsgInsertBTCSpvProofResponse{}, nil
}
