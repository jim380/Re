package ics_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	appProvider "github.com/cosmos/interchain-security/app/provider"
	"github.com/cosmos/interchain-security/tests/e2e"
	e2etestutil "github.com/cosmos/interchain-security/testutil/e2e"
	icssimapp "github.com/cosmos/interchain-security/testutil/simapp"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/jim380/Re/app"
)

// Executes the standard group of ccv tests against a consumer and provider app.go implementation.
func TestCCVTestSuite(t *testing.T) {
	ccvSuite := e2e.NewCCVTestSuite(
		func(t *testing.T) (
			*ibctesting.Coordinator,
			*ibctesting.TestChain,
			*ibctesting.TestChain,
			e2etestutil.ProviderApp,
			e2etestutil.ConsumerApp,
		) {
			// Here we pass the concrete types that must implement the necessary interfaces
			// to be ran with e2e tests.
			coord, prov, cons := NewProviderConsumerCoordinator(t)
			return coord, prov, cons, prov.App.(*appProvider.App), cons.App.(*app.ReApp)
		},
	)
	suite.Run(t, ccvSuite)
}

// NewCoordinator initializes Coordinator with interchain security dummy provider and Re consumer chain
func NewProviderConsumerCoordinator(t *testing.T) (*ibctesting.Coordinator, *ibctesting.TestChain, *ibctesting.TestChain) {
	coordinator := icssimapp.NewBasicCoordinator(t)
	chainID := ibctesting.GetChainID(1)
	coordinator.Chains[chainID] = ibctesting.NewTestChain(t, coordinator, icssimapp.SetupTestingappProvider, chainID)
	providerChain := coordinator.GetChain(chainID)
	chainID = ibctesting.GetChainID(2)
	coordinator.Chains[chainID] = ibctesting.NewTestChainWithValSet(t, coordinator,
		SetupTestingAppConsumer, chainID, providerChain.Vals, providerChain.Signers)
	consumerChain := coordinator.GetChain(chainID)
	return coordinator, providerChain, consumerChain
}

func SetupTestingAppConsumer() (ibctesting.TestingApp, map[string]json.RawMessage) {
	encoding := app.MakeEncodingConfig()
	testApp := app.NewReApp(
		log.NewNopLogger(),
		dbm.NewMemDB(),
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
