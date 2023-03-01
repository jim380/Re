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

	opWeightMsgLogonInitiator = "op_weight_msg_logon_initiator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogonInitiator int = 100

	opWeightMsgLogonAcceptor = "op_weight_msg_logon_acceptor"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogonAcceptor int = 100

	opWeightMsgTerminateLogon = "op_weight_msg_terminate_logon"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTerminateLogon int = 100

	opWeightMsgLogonReject = "op_weight_msg_logon_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogonReject int = 100

	opWeightMsgLogoutInitiator = "op_weight_msg_logout_initiator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogoutInitiator int = 100

	opWeightMsgLogoutAcceptor = "op_weight_msg_logout_acceptor"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogoutAcceptor int = 100

	opWeightMsgNewOrderSingle = "op_weight_msg_new_order_single"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNewOrderSingle int = 100

	opWeightMsgOrderCancelRequest = "op_weight_msg_order_cancel_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOrderCancelRequest int = 100

	opWeightMsgOrderCancelReject = "op_weight_msg_order_cancel_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOrderCancelReject int = 100

	opWeightMsgOrderExecutionReport = "op_weight_msg_order_execution_report"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOrderExecutionReport int = 100

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

	var weightMsgLogonInitiator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLogonInitiator, &weightMsgLogonInitiator, nil,
		func(_ *rand.Rand) {
			weightMsgLogonInitiator = defaultWeightMsgLogonInitiator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogonInitiator,
		fixsimulation.SimulateMsgLogonInitiator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLogonAcceptor int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLogonAcceptor, &weightMsgLogonAcceptor, nil,
		func(_ *rand.Rand) {
			weightMsgLogonAcceptor = defaultWeightMsgLogonAcceptor
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogonAcceptor,
		fixsimulation.SimulateMsgLogonAcceptor(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTerminateLogon int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTerminateLogon, &weightMsgTerminateLogon, nil,
		func(_ *rand.Rand) {
			weightMsgTerminateLogon = defaultWeightMsgTerminateLogon
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTerminateLogon,
		fixsimulation.SimulateMsgTerminateLogon(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLogonReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLogonReject, &weightMsgLogonReject, nil,
		func(_ *rand.Rand) {
			weightMsgLogonReject = defaultWeightMsgLogonReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogonReject,
		fixsimulation.SimulateMsgLogonReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLogoutInitiator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLogoutInitiator, &weightMsgLogoutInitiator, nil,
		func(_ *rand.Rand) {
			weightMsgLogoutInitiator = defaultWeightMsgLogoutInitiator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogoutInitiator,
		fixsimulation.SimulateMsgLogoutInitiator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLogoutAcceptor int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLogoutAcceptor, &weightMsgLogoutAcceptor, nil,
		func(_ *rand.Rand) {
			weightMsgLogoutAcceptor = defaultWeightMsgLogoutAcceptor
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogoutAcceptor,
		fixsimulation.SimulateMsgLogoutAcceptor(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgNewOrderSingle int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNewOrderSingle, &weightMsgNewOrderSingle, nil,
		func(_ *rand.Rand) {
			weightMsgNewOrderSingle = defaultWeightMsgNewOrderSingle
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNewOrderSingle,
		fixsimulation.SimulateMsgNewOrderSingle(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderCancelRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderCancelRequest, &weightMsgOrderCancelRequest, nil,
		func(_ *rand.Rand) {
			weightMsgOrderCancelRequest = defaultWeightMsgOrderCancelRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderCancelRequest,
		fixsimulation.SimulateMsgOrderCancelRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderCancelReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderCancelReject, &weightMsgOrderCancelReject, nil,
		func(_ *rand.Rand) {
			weightMsgOrderCancelReject = defaultWeightMsgOrderCancelReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderCancelReject,
		fixsimulation.SimulateMsgOrderCancelReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderExecutionReport int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderExecutionReport, &weightMsgOrderExecutionReport, nil,
		func(_ *rand.Rand) {
			weightMsgOrderExecutionReport = defaultWeightMsgOrderExecutionReport
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderExecutionReport,
		fixsimulation.SimulateMsgOrderExecutionReport(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
