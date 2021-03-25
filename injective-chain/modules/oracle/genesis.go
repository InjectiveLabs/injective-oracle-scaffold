package oracle

import (
	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/keeper"
	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) {
	// TODO: @Albert initialize genesis properly
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// TODO: @Albert export genesis properly
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
