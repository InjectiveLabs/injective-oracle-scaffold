package keeper

import (
	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the total set of oracle parameters.
func (k BaseKeeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k BaseKeeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
