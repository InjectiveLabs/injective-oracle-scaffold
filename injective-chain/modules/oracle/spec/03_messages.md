# Messages

## MsgRelay

### struct

```go
type MsgRelay struct {
    Relayer      sdk.AccAddress
	Symbols      []string
	Rates        []uint64
    ResolveTimes []uint64
    RequestIDs   []uint64
}

// ValidateBasic implements the sdk.Msg interface for MsgRelay.
func (msg MsgRelay) ValidateBasic() error {
    // note that unmarshaling from bech32 ensures either empty or valid
    if msg.Relayer.Empty() {
        return types.ErrEmptyRelayerAddr
    }

    // check that the sizes of symbols,rates,resolveTimes,requestIDs are equal
    symbolsCount := len(msg.Symbols)
    if len(msg.Rates) !== symbolsCount {
        return nil, types.ErrBadRatesCount
    }
    if len(msg.ResolveTimes) !== symbolsCount || {
        return nil, types.ErrBadResolveTimesCount
    }
    if len(msg.RequestIDs) !== symbolsCount {
        return nil, types.ErrBadRequestIDsCount
    }

    return nil
}
```

### handleMsgRelay

```go
func handleMsgRelay(ctx sdk.Context, k Keeper, m MsgRelay) (*sdk.Result, error) {
	if !k.IsRelayer(ctx, m.Relayer) {
		return nil, types.ErrRelayerNotAuthorized
	}

    // loop SetRef for all symbols
    for i := 0; i < len(m.Symbols); i++ {
        k.SetRef(
            ctx,
            m.Symbols[i],
            types.Ref.NewRef(m.Rates[i], m.ResolveTimes[i], m.RequestIDs[i])
        )
    }

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
```
