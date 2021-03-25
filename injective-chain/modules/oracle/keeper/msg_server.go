package keeper

import (
	"context"
	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/types"
	"github.com/InjectiveLabs/injective-oracle-scaffold/metrics"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

type msgServer struct {
	Keeper
	svcTags metrics.Tags
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided BaseKeeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		Keeper: keeper,
		svcTags: metrics.Tags{
			"svc": "oracle_h",
		},
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Relay(goCtx context.Context, msg *types.MsgRelay) (*types.MsgRelayResponse, error) {
	// metric collector
	metrics.ReportFuncCall(k.svcTags)
	doneFn := metrics.ReportFuncTiming(k.svcTags)
	defer doneFn()

	// prepare context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check params to verify that msg.Relayer is an authorized relayer
	relayer, _ := sdk.AccAddressFromBech32(msg.Relayer)

	if !k.IsRelayer(ctx, relayer) {
		return nil, types.ErrRelayerNotAuthorized
	}

	// loop SetRef for all symbols
	for idx := range msg.Symbols {

		symbol := msg.Symbols[idx]
		rate := msg.Rates[idx]
		resolveTime := msg.ResolveTimes[idx]
		requestID := msg.RequestIDs[idx]

		k.SetRef(
			ctx,
			symbol,
			types.Ref{Rate: rate, ResolveTime: resolveTime, Request_ID: requestID},
		)

		// emit EventTypeSetRef event
		ctx.EventManager().EmitTypedEvent(&types.SetRef{
			Relayer:     msg.Relayer,
			Symbol:      symbol,
			Rate:        strconv.FormatUint(rate, 10),
			ResolveTime: strconv.FormatUint(resolveTime, 10),
			RequestId:   strconv.FormatUint(requestID, 10),
		})
	}

	return &types.MsgRelayResponse{}, nil
}

func (k msgServer) SetPrice(goCtx context.Context, msg *types.MsgSetPrice) (*types.MsgSetPriceResponse, error) {
	// metric collector
	metrics.ReportFuncCall(k.svcTags)
	doneFn := metrics.ReportFuncTiming(k.svcTags)
	defer doneFn()

	// prepare context
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx
	return &types.MsgSetPriceResponse{}, nil
}
