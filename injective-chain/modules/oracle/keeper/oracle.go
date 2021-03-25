package keeper

import (
	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/types"
	"github.com/InjectiveLabs/injective-oracle-scaffold/metrics"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

func (k BaseKeeper) GetRef(ctx sdk.Context, symbol string) types.Ref {
	// metric collector
	metrics.ReportFuncCall(k.svcTags)
	doneFn := metrics.ReportFuncTiming(k.svcTags)
	defer doneFn()

	var ref types.Ref
	bz := k.GetStore(ctx).Get(types.GetRefStoreKey(symbol))
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ref)
	return ref
}

func (k BaseKeeper) SetRef(ctx sdk.Context, symbol string, refData types.Ref) {
	// metric collector
	metrics.ReportFuncCall(k.svcTags)
	doneFn := metrics.ReportFuncTiming(k.svcTags)
	defer doneFn()

	bz := k.cdc.MustMarshalBinaryBare(&refData)
	k.GetStore(ctx).Set(types.GetRefStoreKey(symbol), bz)
}

func (k BaseKeeper) GetReferenceData(ctx sdk.Context, base string, quote string) *types.ReferenceData {
	metrics.ReportFuncCall(k.svcTags)
	doneFn := metrics.ReportFuncTiming(k.svcTags)
	defer doneFn()

	// query ref by using GetRef
	baseRef := k.GetRef(ctx, base)
	quoteRef := k.GetRef(ctx, quote)

	// convert from uint64 to sdk.Int
	baseRate := sdk.NewIntFromUint64(baseRef.Rate)
	quoteRate := sdk.NewIntFromUint64(quoteRef.Rate)
	e9 := sdk.NewInt(1000000000)

	// create reference data from base ref and quote ref
	referenceData := &types.ReferenceData{
		Rate:            baseRate.Mul(e9).Quo(quoteRate),
		LastUpdateBase:  baseRef.ResolveTime,
		LastUpdateQuote: quoteRef.ResolveTime,
	}

	// emit GetReferenceData event
	ctx.EventManager().EmitTypedEvent(&types.GetReferenceData{
		BaseQuoteSymbol: base + "/" + quote,
		BaseQuoteRate:   referenceData.Rate.String(),
		LastUpdateBase:  strconv.FormatUint(baseRef.ResolveTime, 10),
		LastUpdateQuote: strconv.FormatUint(quoteRef.ResolveTime, 10),
	})

	return referenceData
}

func (k BaseKeeper) GetReferenceDataBulk(ctx sdk.Context, bases []string, quotes []string) []*types.ReferenceData {
	metrics.ReportFuncCall(k.svcTags)
	doneFn := metrics.ReportFuncTiming(k.svcTags)
	defer doneFn()

	if len(bases) != len(quotes) {
		return nil
	}

	referenceDataSlice := make([]*types.ReferenceData, len(bases))

	for idx := range bases {
		referenceDataSlice[idx] = k.GetReferenceData(ctx, bases[idx], quotes[idx])
	}
	return referenceDataSlice
}

func (k BaseKeeper) IsRelayer(ctx sdk.Context, relayer sdk.AccAddress) bool {
	params := k.GetParams(ctx)
	for idx := range params.Relayers {
		if relayer.String() == params.Relayers[idx] {
			return true
		}
	}
	return false
}
