package app

import (
	bbn "github.com/babylonchain/babylon/types"
	btccheckpointtypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btclightclient "github.com/babylonchain/babylon/x/btclightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BtcValidationDecorator struct {
	BtcCfg bbn.BtcConfig
}

func NewBtcValidationDecorator(cfg bbn.BtcConfig) BtcValidationDecorator {
	return BtcValidationDecorator{
		BtcCfg: cfg,
	}
}

func (bvd BtcValidationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {

	// only do this validation when handling mempool addition. During DeliverTx they
	// should be performed by btclightclient and btccheckpoint modules
	if ctx.IsCheckTx() || ctx.IsReCheckTx() {
		for _, m := range tx.GetMsgs() {
			switch msg := m.(type) {
			case *btccheckpointtypes.MsgInsertBTCSpvProof:
				powLimit := bvd.BtcCfg.PowLimit()
				_, err := btccheckpointtypes.ParseSubmission(msg, &powLimit, bvd.BtcCfg.CheckpointTag())

				if err != nil {
					return ctx, btccheckpointtypes.ErrInvalidCheckpointProof.Wrap(err.Error())
				}

			case *btclightclient.MsgInsertHeader:
				powLimit := bvd.BtcCfg.PowLimit()
				err := msg.ValidateHeader(&powLimit)

				if err != nil {
					return ctx, btclightclient.ErrInvalidProofOfWOrk
				}
			default:
				// NOOP in case of other messages
			}
		}
	}
	return next(ctx, tx, simulate)
}
