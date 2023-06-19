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

	opWeightMsgQuoteRequest = "op_weight_msg_quote_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgQuoteRequest int = 100

	opWeightMsgQuoteAcknowledgement = "op_weight_msg_quote_acknowledgement"
	// TODO: Determine the simulation weight value
	defaultWeightMsgQuoteAcknowledgement int = 100

	opWeightMsgQuoteRequestReject = "op_weight_msg_quote_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgQuoteRequestReject int = 100

	opWeightMsgSecurityDefinitionRequest = "op_weight_msg_security_definition_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityDefinitionRequest int = 100

	opWeightMsgSecurityDefinition = "op_weight_msg_security_definition"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityDefinition int = 100

	opWeightMsgSecurityDefinitionRequestReject = "op_weight_msg_security_definition_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityDefinitionRequestReject int = 100

	opWeightMsgOrderMassStatusRequest = "op_weight_msg_order_mass_status_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOrderMassStatusRequest int = 100

	opWeightMsgOrderMassStatusReport = "op_weight_msg_order_mass_status_report"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOrderMassStatusReport int = 100

	opWeightMsgOrderMassStatusRequestReject = "op_weight_msg_order_mass_status_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOrderMassStatusRequestReject int = 100

	opWeightMsgTradingSessionStatusRequest = "op_weight_msg_trading_session_status_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTradingSessionStatusRequest int = 100

	opWeightMsgTradingSessionStatus = "op_weight_msg_trading_session_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTradingSessionStatus int = 100

	opWeightMsgTradingSessionStatusRequestReject = "op_weight_msg_trading_session_status_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTradingSessionStatusRequestReject int = 100

	opWeightMsgTradingSessionListRequest = "op_weight_msg_trading_session_list_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTradingSessionListRequest int = 100

	opWeightMsgTradingSessionListResponse = "op_weight_msg_trading_session_list_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTradingSessionListResponse int = 100

	opWeightMsgTradingSessionListRequestReject = "op_weight_msg_trading_session_list_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTradingSessionListRequestReject int = 100

	opWeightMsgSecurityListRequest = "op_weight_msg_security_list_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityListRequest int = 100

	opWeightMsgSecurityListResponse = "op_weight_msg_security_list_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityListResponse int = 100

	opWeightMsgSecurityListRequestReject = "op_weight_msg_security_list_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityListRequestReject int = 100

	opWeightMsgSecurityStatusRequest = "op_weight_msg_security_status_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityStatusRequest int = 100

	opWeightMsgSecurityStatusResponse = "op_weight_msg_security_status_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityStatusResponse int = 100

	opWeightMsgSecurityStatusRequestReject = "op_weight_msg_security_status_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityStatusRequestReject int = 100

	opWeightMsgSecurityTypesRequest = "op_weight_msg_security_types_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityTypesRequest int = 100

	opWeightMsgSecurityTypesResponse = "op_weight_msg_security_types_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityTypesResponse int = 100

	opWeightMsgSecurityTypesRequestReject = "op_weight_msg_security_types_request_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSecurityTypesRequestReject int = 100

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
				Address: sample.AccAddress(),
			},
			{
				Address: sample.AccAddress(),
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

	var weightMsgQuoteRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgQuoteRequest, &weightMsgQuoteRequest, nil,
		func(_ *rand.Rand) {
			weightMsgQuoteRequest = defaultWeightMsgQuoteRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgQuoteRequest,
		fixsimulation.SimulateMsgQuoteRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgQuoteAcknowledgement int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgQuoteAcknowledgement, &weightMsgQuoteAcknowledgement, nil,
		func(_ *rand.Rand) {
			weightMsgQuoteAcknowledgement = defaultWeightMsgQuoteAcknowledgement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgQuoteAcknowledgement,
		fixsimulation.SimulateMsgQuoteAcknowledgement(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgQuoteRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgQuoteRequestReject, &weightMsgQuoteRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgQuoteRequestReject = defaultWeightMsgQuoteRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgQuoteRequestReject,
		fixsimulation.SimulateMsgQuoteRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityDefinitionRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityDefinitionRequest, &weightMsgSecurityDefinitionRequest, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityDefinitionRequest = defaultWeightMsgSecurityDefinitionRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityDefinitionRequest,
		fixsimulation.SimulateMsgSecurityDefinitionRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityDefinition int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityDefinition, &weightMsgSecurityDefinition, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityDefinition = defaultWeightMsgSecurityDefinition
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityDefinition,
		fixsimulation.SimulateMsgSecurityDefinition(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityDefinitionRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityDefinitionRequestReject, &weightMsgSecurityDefinitionRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityDefinitionRequestReject = defaultWeightMsgSecurityDefinitionRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityDefinitionRequestReject,
		fixsimulation.SimulateMsgSecurityDefinitionRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderMassStatusRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderMassStatusRequest, &weightMsgOrderMassStatusRequest, nil,
		func(_ *rand.Rand) {
			weightMsgOrderMassStatusRequest = defaultWeightMsgOrderMassStatusRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderMassStatusRequest,
		fixsimulation.SimulateMsgOrderMassStatusRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderMassStatusReport int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderMassStatusReport, &weightMsgOrderMassStatusReport, nil,
		func(_ *rand.Rand) {
			weightMsgOrderMassStatusReport = defaultWeightMsgOrderMassStatusReport
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderMassStatusReport,
		fixsimulation.SimulateMsgOrderMassStatusReport(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderMassStatusRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderMassStatusRequestReject, &weightMsgOrderMassStatusRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgOrderMassStatusRequestReject = defaultWeightMsgOrderMassStatusRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderMassStatusRequestReject,
		fixsimulation.SimulateMsgOrderMassStatusRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTradingSessionStatusRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTradingSessionStatusRequest, &weightMsgTradingSessionStatusRequest, nil,
		func(_ *rand.Rand) {
			weightMsgTradingSessionStatusRequest = defaultWeightMsgTradingSessionStatusRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTradingSessionStatusRequest,
		fixsimulation.SimulateMsgTradingSessionStatusRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTradingSessionStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTradingSessionStatus, &weightMsgTradingSessionStatus, nil,
		func(_ *rand.Rand) {
			weightMsgTradingSessionStatus = defaultWeightMsgTradingSessionStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTradingSessionStatus,
		fixsimulation.SimulateMsgTradingSessionStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTradingSessionStatusRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTradingSessionStatusRequestReject, &weightMsgTradingSessionStatusRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgTradingSessionStatusRequestReject = defaultWeightMsgTradingSessionStatusRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTradingSessionStatusRequestReject,
		fixsimulation.SimulateMsgTradingSessionStatusRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTradingSessionListRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTradingSessionListRequest, &weightMsgTradingSessionListRequest, nil,
		func(_ *rand.Rand) {
			weightMsgTradingSessionListRequest = defaultWeightMsgTradingSessionListRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTradingSessionListRequest,
		fixsimulation.SimulateMsgTradingSessionListRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTradingSessionListResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTradingSessionListResponse, &weightMsgTradingSessionListResponse, nil,
		func(_ *rand.Rand) {
			weightMsgTradingSessionListResponse = defaultWeightMsgTradingSessionListResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTradingSessionListResponse,
		fixsimulation.SimulateMsgTradingSessionListResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTradingSessionListRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTradingSessionListRequestReject, &weightMsgTradingSessionListRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgTradingSessionListRequestReject = defaultWeightMsgTradingSessionListRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTradingSessionListRequestReject,
		fixsimulation.SimulateMsgTradingSessionListRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityListRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityListRequest, &weightMsgSecurityListRequest, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityListRequest = defaultWeightMsgSecurityListRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityListRequest,
		fixsimulation.SimulateMsgSecurityListRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityListResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityListResponse, &weightMsgSecurityListResponse, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityListResponse = defaultWeightMsgSecurityListResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityListResponse,
		fixsimulation.SimulateMsgSecurityListResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityListRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityListRequestReject, &weightMsgSecurityListRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityListRequestReject = defaultWeightMsgSecurityListRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityListRequestReject,
		fixsimulation.SimulateMsgSecurityListRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityStatusRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityStatusRequest, &weightMsgSecurityStatusRequest, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityStatusRequest = defaultWeightMsgSecurityStatusRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityStatusRequest,
		fixsimulation.SimulateMsgSecurityStatusRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityStatusResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityStatusResponse, &weightMsgSecurityStatusResponse, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityStatusResponse = defaultWeightMsgSecurityStatusResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityStatusResponse,
		fixsimulation.SimulateMsgSecurityStatusResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityStatusRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityStatusRequestReject, &weightMsgSecurityStatusRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityStatusRequestReject = defaultWeightMsgSecurityStatusRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityStatusRequestReject,
		fixsimulation.SimulateMsgSecurityStatusRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityTypesRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityTypesRequest, &weightMsgSecurityTypesRequest, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityTypesRequest = defaultWeightMsgSecurityTypesRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityTypesRequest,
		fixsimulation.SimulateMsgSecurityTypesRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityTypesResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityTypesResponse, &weightMsgSecurityTypesResponse, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityTypesResponse = defaultWeightMsgSecurityTypesResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityTypesResponse,
		fixsimulation.SimulateMsgSecurityTypesResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSecurityTypesRequestReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSecurityTypesRequestReject, &weightMsgSecurityTypesRequestReject, nil,
		func(_ *rand.Rand) {
			weightMsgSecurityTypesRequestReject = defaultWeightMsgSecurityTypesRequestReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSecurityTypesRequestReject,
		fixsimulation.SimulateMsgSecurityTypesRequestReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
