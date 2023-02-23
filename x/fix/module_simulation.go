package fix

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jim380/Re/testutil/sample"
	fixsimulation "github.com/jim380/Re/x/fix/simulation"
	"github.com/jim380/Re/x/fix/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = fixsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAccount = "op_weight_msg_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAccount int = 100

	opWeightMsgUpdateAccount = "op_weight_msg_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAccount int = 100

	opWeightMsgDeleteAccount = "op_weight_msg_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAccount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	fixGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AccountList: []types.Account{
			{
				Did:     "",
				Creator: sample.AccAddress(),
			},
			{
				Did:     "",
				Creator: sample.AccAddress(),
			},
		},
		AccountCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&fixGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAccount, &weightMsgCreateAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAccount = defaultWeightMsgCreateAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAccount,
		fixsimulation.SimulateMsgCreateAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAccount, &weightMsgUpdateAccount, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAccount = defaultWeightMsgUpdateAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAccount,
		fixsimulation.SimulateMsgUpdateAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAccount, &weightMsgDeleteAccount, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAccount = defaultWeightMsgDeleteAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAccount,
		fixsimulation.SimulateMsgDeleteAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
