package testutil

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	ibctesting "github.com/cosmos/interchain-security/legacy_ibc_testing/testing"
	icssimapp "github.com/cosmos/interchain-security/testutil/ibc_testing"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/jim380/Re/app"
)

// NewCoordinator initializes Coordinator with interchain security dummy provider and Re consumer chain
func NewProviderConsumerCoordinator(t *testing.T) (*ibctesting.Coordinator, *ibctesting.TestChain, *ibctesting.TestChain) {
	coordinator := ibctesting.NewCoordinator(t, 1)
	chainID := ibctesting.GetChainID(1)
	coordinator.Chains[chainID] = ibctesting.NewTestChain(t, coordinator, icssimapp.ProviderAppIniter, chainID)
	providerChain := coordinator.GetChain(chainID)
	chainID = ibctesting.GetChainID(2)
	coordinator.Chains[chainID] = ibctesting.NewTestChainWithValSet(t, coordinator,
		SetupTestingApp, chainID, providerChain.Vals, providerChain.Signers)
	consumerChain := coordinator.GetChain(chainID)
	return coordinator, providerChain, consumerChain
}

// SetupTestingApp initializes the IBC-go testing application
func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	encoding := app.MakeEncodingConfig()
	db := dbm.NewMemDB()
	testApp := app.NewReApp(
		log.NewNopLogger(),
		db,
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
	return testApp, app.NewDefaultGenesisState(testApp.AppCodec())
}
