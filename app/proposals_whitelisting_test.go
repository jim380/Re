package app_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	icssimapp "github.com/cosmos/interchain-security/testutil/simapp"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	"github.com/jim380/Re/app"
	//"github.com/jim380/Re/cmd"
)

func TestConsumerWhitelistingKeys(t *testing.T) {
	chain := ibctesting.NewTestChain(t, icssimapp.NewBasicCoordinator(t), SetupTestingAppConsumer, "test")
	paramKeeper := chain.App.(*app.ReApp).ParamsKeeper
	for paramKey := range app.WhitelistedParams {
		ss, ok := paramKeeper.GetSubspace(paramKey.Subspace)
		require.True(t, ok, "Unknown subspace %s", paramKey.Subspace)
		hasKey := ss.Has(chain.GetContext(), []byte(paramKey.Key))
		require.True(t, hasKey, "Invalid key %s for subspace %s", paramKey.Key, paramKey.Subspace)
	}
}

func SetupTestingAppConsumer() (ibctesting.TestingApp, map[string]json.RawMessage) {
	encoding := app.MakeEncodingConfig()
	testApp := app.NewReApp(
		log.NewNopLogger(),
		tmdb.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		0,
		encoding,
		app.GetEnabledProposals(),
		simapp.EmptyAppOptions{},
		nil,
	)

	return testApp, app.NewDefaultGenesisState(encoding.Marshaler)
}
