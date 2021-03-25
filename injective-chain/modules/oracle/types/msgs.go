package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgSetPrice{}
	_ sdk.Msg = &MsgRelay{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgSetPrice) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgSetPrice) Type() string { return "msgSetPrice" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgSetPrice) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgSetPrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgSetPrice) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRelay) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRelay) Type() string { return "msgRelay" }

// ValidateBasic implements the sdk.Msg interface for MsgRelay.
func (msg MsgRelay) ValidateBasic() error {
	if msg.Relayer == "" {
		return ErrEmptyRelayerAddr
	}

	// check that the sizes of symbols,rates,resolveTimes,requestIDs are equal
	symbolsCount := len(msg.Symbols)
	if len(msg.Rates) != symbolsCount {
		return ErrBadRatesCount
	}
	if len(msg.ResolveTimes) != symbolsCount {
		return ErrBadResolveTimesCount
	}
	if len(msg.RequestIDs) != symbolsCount {
		return ErrBadRequestIDsCount
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRelay) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRelay) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
