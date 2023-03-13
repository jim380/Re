package mic

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jim380/Re/testutil/sample"
	micsimulation "github.com/jim380/Re/x/mic/simulation"
	"github.com/jim380/Re/x/mic/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = micsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterMarketIdentificationCode = "op_weight_msg_market_identification_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterMarketIdentificationCode int = 100

	opWeightMsgUpdateMarketIdentificationCode = "op_weight_msg_market_identification_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMarketIdentificationCode int = 100

	opWeightMsgDeleteMarketIdentificationCode = "op_weight_msg_market_identification_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMarketIdentificationCode int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	micGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MarketIdentificationCodeList: []types.MarketIdentificationCode{
			{
				MIC:      "",
				Creator: sample.AccAddress(),
			},
			{
				MIC:      "",
				Creator: sample.AccAddress(),
			},
		},
		MarketIdentificationCodeCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&micGenesis)
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

	var weightMsgRegisterMarketIdentificationCode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterMarketIdentificationCode, &weightMsgRegisterMarketIdentificationCode, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterMarketIdentificationCode = defaultWeightMsgRegisterMarketIdentificationCode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterMarketIdentificationCode,
		micsimulation.SimulateMsgRegisterMarketIdentificationCode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMarketIdentificationCode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMarketIdentificationCode, &weightMsgUpdateMarketIdentificationCode, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMarketIdentificationCode = defaultWeightMsgUpdateMarketIdentificationCode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMarketIdentificationCode,
		micsimulation.SimulateMsgUpdateMarketIdentificationCode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMarketIdentificationCode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMarketIdentificationCode, &weightMsgDeleteMarketIdentificationCode, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMarketIdentificationCode = defaultWeightMsgDeleteMarketIdentificationCode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMarketIdentificationCode,
		micsimulation.SimulateMsgDeleteMarketIdentificationCode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
