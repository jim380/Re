package keeper

import (
	"testing"

	"github.com/jim380/Re/x/tokenregistry/keeper"
	"github.com/jim380/Re/x/tokenregistry/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// tokenregistryChannelKeeper is a stub of cosmosibckeeper.ChannelKeeper.
type tokenregistryChannelKeeper struct{}

func (tokenregistryChannelKeeper) GetChannel(ctx sdk.Context, portID, channelID string) (channeltypes.Channel, bool) {
	return channeltypes.Channel{}, false
}

func (tokenregistryChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return 0, false
}

func (tokenregistryChannelKeeper) SendPacket(
    ctx sdk.Context,
    channelCap *capabilitytypes.Capability,
    sourcePort string,
    sourceChannel string,
    timeoutHeight clienttypes.Height,
    timeoutTimestamp uint64,
    data []byte,
) (uint64, error) {
    return 0, nil
}

func (tokenregistryChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	return nil
}

// tokenregistryportKeeper is a stub of cosmosibckeeper.PortKeeper
type tokenregistryPortKeeper struct{}

func (tokenregistryPortKeeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return &capabilitytypes.Capability{}
}



func TokenregistryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"TokenregistryParams",
	)
	k := keeper.NewKeeper(
        appCodec,
        storeKey,
        memStoreKey,
        paramsSubspace,
        tokenregistryChannelKeeper{},
        tokenregistryPortKeeper{},
        capabilityKeeper.ScopeToModule("TokenregistryScopedKeeper"),
    )

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
