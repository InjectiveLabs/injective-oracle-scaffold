package keeper

import (
	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/types"
	"github.com/InjectiveLabs/injective-oracle-scaffold/metrics"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	log "github.com/xlab/suplog"
)

// Keeper defines a module interface that facilitates the getting and setting of oracle reference data
type Keeper interface {
	// read the stored
	GetRef(ctx sdk.Context, symbol string) types.Ref
	// can only be called by Relay function
	SetRef(ctx sdk.Context, symbol string, refData types.Ref)

	// derive a ReferenceData from base Ref and quote Ref
	GetReferenceData(ctx sdk.Context, base string, quote string) *types.ReferenceData
	GetReferenceDataBulk(ctx sdk.Context, bases []string, quotes []string) []*types.ReferenceData

	GetParams(ctx sdk.Context) types.Params
	// can only be set by a proposal
	SetParams(ctx sdk.Context, params types.Params)

	// check that the relayer has been authorized
	IsRelayer(ctx sdk.Context, relayer sdk.AccAddress) bool

	types.QueryServer
}

// BaseKeeper of this module maintains collections of oracle.
type BaseKeeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	paramSpace paramtypes.Subspace

	accountKeeper authkeeper.AccountKeeper
	bankKeeper    types.BankKeeper

	logger  log.Logger
	svcTags metrics.Tags
}

// NewKeeper creates new instances of the oracle BaseKeeper
func NewKeeper(
	cdc codec.BinaryMarshaler,
	storeKey sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	ak authkeeper.AccountKeeper,
	bk types.BankKeeper,
) BaseKeeper {

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return BaseKeeper{
		svcTags: metrics.Tags{
			"svc": "oracle_k",
		},
		paramSpace: paramSpace,

		storeKey:      storeKey,
		cdc:           cdc,
		accountKeeper: ak,
		bankKeeper:    bk,
		logger:        log.WithField("module", types.ModuleName),
	}
}

func (k BaseKeeper) GetStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}
