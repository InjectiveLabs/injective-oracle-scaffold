package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Parameter keys
var (
	KeyOracleRelayers = []byte("OracleRelayers")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		Relayers: []string{},
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	// TODO: @albert, add the rest of the parameters
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyOracleRelayers, &p.Relayers, validateRelayers),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{}
}

// Validate performs basic validation on auction parameters.
func (p Params) Validate() error {
	if err := validateRelayers(p.Relayers); err != nil {
		return err
	}

	return nil
}

func validateRelayers(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, relayer := range v {
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return err
		}
	}

	return nil
}
