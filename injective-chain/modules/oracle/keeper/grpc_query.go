package keeper

import (
	"context"
	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = BaseKeeper{}

func (k BaseKeeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := k.GetParams(ctx)

	res := &types.QueryParamsResponse{
		Params: params,
	}
	return res, nil
}
