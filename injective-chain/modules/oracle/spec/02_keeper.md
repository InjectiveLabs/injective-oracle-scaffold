# Keepers

- The stdreference module provides three different exported keeper interfaces which can be passed to other modules which need to read price data. The price data should only be updated by the relayer.

## Types

### Ref

A raw data from BandChain, which has a requestID inside it. The requestID will allow anyone to check that the price data actually exists in the BandChain.

```go
type Ref struct {
    Rate        uint64
    ResolveTime uint64 // unix timestamp
    RequestID   uint64
}
```

### ReferenceData

An output that derived from the raw data or Ref. The Rate of ReferenceData will calculated by dividing of 2 Rates from the 2 Refs (base and quote) and then keep the timestamp of them but cut out the reqeustID.

```go
type ReferenceData struct {
    Rate            sdk.Int
    LastUpdateBase  uint64 // unix timestamp
    LastUpdateQuote uint64 // unix timestamp
}
```

For example

```go
// We use normal int in this example just for simplicity.
// Normally we would use uint64 for Rate and any time value (unix timestamp).
base := Ref {Rate: 2, ResolveTime: 4, RequestID: 6}
quote := Ref {Rate: 1, ResolveTime: 3, RequestID: 5}

referenceData := ReferenceData {Rate: base.Rate * uint64(1e9) / quote.Rate, LastUpdateBase: base.ResolveTime, LastUpdateQuote: quote.ResolveTime}

fmt.Println(base, quote, referenceData)

// An outout should be "{2 4 6} {1 3 5} {2 4 3}"
```

## Keeper

```go
// Keeper defines a module interface that facilitates the transfer of coins
// between accounts.
type Keeper interface {
    // read the stored
    GetRef(ctx sdk.Context, symbol string) types.Ref
    // can only be called by Relay function
	SetRef(ctx sdk.Context, symbol string, refData types.Ref)

    // derive a ReferenceData from base Ref and quote Ref
    GetReferenceData(ctx sdk.Context, base string, quote string) types.ReferenceData
    GetReferenceDataBulk(ctx sdk.Context, bases []string, quotes []string) []types.ReferenceData

    GetParams(ctx sdk.Context) types.Params
    // can only be set by a proposal
	SetParams(ctx sdk.Context, params types.Params)

    // check that the relayer has been authorized
    IsRelayer(ctx sdk.Context, relayer sdk.AccAddress) bool
}
```

### Example Implementation

#### GetRef

```go
func (k Keeper) GetRef(ctx sdk.Context, symbol string) types.Ref {
	var ref Ref
	bz := ctx.KVStore(k.storeKey).Get(types.GetRefStoreKey(symbol))
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ref)
	return ref
}
```

#### SetRef

```go
func (k Keeper) SetRef(ctx sdk.Context, symbol string, refData types.Ref) {
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(refData)
	ctx.KVStore(k.storeKey).Set(types.GetRefStoreKey(symbol), bz)

    // emit EventTypeSetRef event
    ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetRef,
        sdk.NewAttribute(types.AttributeKeyRelayer, fmt.Sprintf("%s", m.Relayer)),
		sdk.NewAttribute(types.AttributeKeySymbol, fmt.Sprintf("%s", m.Symbol)),
        sdk.NewAttribute(types.AttributeKeyRate, fmt.Sprintf("%d", m.Ref.Rate)),
        sdk.NewAttribute(types.AttributeKeyResolveTime, fmt.Sprintf("%d", m.Ref.ResolveTime)),
        sdk.NewAttribute(types.AttributeKeyRequestID, fmt.Sprintf("%d", m.Ref.RequestID)),
	))
}
```

#### GetReferenceData

```go
func (k Keeper) GetReferenceData(ctx sdk.Context, base string, quote string) types.ReferenceData {
    // query ref by using GetRef
	baseRef := k.GetRef(ctx,base)
    quoteRef := k.GetRef(ctx,quote)

    // convert from uint64 to sdk.Int
    baseRate := sdk.NewInt(baseRef.Rate)
    quoteRate := sdk.NewInt(quoteRef.Rate)
    e9 := sdk.NewInt(1000000000)

    // create reference data from base ref and quote ref
	referenceData := ReferenceData {
        Rate: baseRate.Mul(e9).Quo(quoteRate)
        LastUpdateBase: baseRef.ResolveTime,
        LastUpdateQuote: quoteRef.ResolveTime,
    }

    // emit event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeGetReferenceData,
        sdk.NewAttribute(types.AttributeKeyBaseQuoteSymbol, fmt.Sprintf("%s/%s", base, quote)),
        sdk.NewAttribute(types.AttributeKeyBaseQuoteRate, referenceData.Rate.String()),
        sdk.NewAttribute(types.AttributeKeyLastUpdateBase, fmt.Sprintf("%d", baseRef.ResolveTime)),
        sdk.NewAttribute(types.AttributeKeyLastUpdateQuote, fmt.Sprintf("%d", quoteRef.ResolveTime)),
	))

	return referenceData
}
```

#### GetReferenceDataBulk

```go
func (k Keeper) GetReferenceDataBulk(ctx sdk.Context, bases []string, quotes []string) []types.ReferenceData {
    referenceDataSlice := []ReferenceData{}
    if len(bases) == len(quotes) {
        for i := 0; i < len(bases); i++ {
            referenceDataSlice = append(referenceDataSlice, k.GetReferenceData(ctx,bases[i],quotes[i]))
        }
    }
    return referenceDataSlice
}
```

#### GetParams

```go
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}
```

#### SetParams

```go
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
```

#### IsRelayer

```go
func (k Keeper) IsRelayer(ctx sdk.Context, relayer sdk.AccAddress) bool
	params := k.GetParams(ctx)
    for _, existedRelayer := range params.Relayers {
        if relayer == existedRelayer {
            return true
        }
    }
    return false
}
```
