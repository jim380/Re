package fix

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the account
	for _, elem := range genState.AccountList {
		k.SetAccount(ctx, elem.Did, elem)
	}

	// Set all the sessions
	for _, elem := range genState.SessionsList {
		k.SetSessions(ctx, elem.SessionID, elem)
	}

	// Set sessions count
	k.SetSessionsCount(ctx, genState.SessionsCount)
	// Set all the sessionReject
	for _, elem := range genState.SessionRejectList {
		k.SetSessionReject(ctx, elem.SessionID, elem)
	}

	// Set all the sessionLogout
	for _, elem := range genState.SessionLogoutList {
		k.SetSessionLogout(ctx, elem.SessionID, elem)
	}

	// Set all the orders
	for _, elem := range genState.OrdersList {
		k.SetOrders(ctx, elem.SessionID, elem)
	}

	// Set all the ordersCancelRequest
	for _, elem := range genState.OrdersCancelRequestList {
		k.SetOrdersCancelRequest(ctx, elem.SessionID, elem)
	}

	// Set all the ordersCancelReject
	for _, elem := range genState.OrdersCancelRejectList {
		k.SetOrdersCancelReject(ctx, elem.SessionID, elem)
	}

	// Set all the ordersExecutionReport
	for _, elem := range genState.OrdersExecutionReportList {
		k.SetOrdersExecutionReport(ctx, elem.SessionID, elem)
	}

	// Set all the quote
	for _, elem := range genState.QuoteList {
		k.SetQuote(ctx, elem.QuoteRequest.QuoteReqID, elem)
	}

	// Set all the tradeCapture
	for _, elem := range genState.TradeCaptureList {
		k.SetTradeCapture(ctx, elem.TradeCaptureReport.TradeReportID, elem)
	}

	// Set tradeCapture count
	k.SetTradeCaptureCount(ctx, genState.TradeCaptureCount)
	// Set all the marketData
	for _, elem := range genState.MarketDataList {
		k.SetMarketData(ctx, elem.MarketDataRequest.MdReqID, elem)
	}

	// Set marketData count
	k.SetMarketDataCount(ctx, genState.MarketDataCount)
	// Set all the security
	for _, elem := range genState.SecurityList {
		k.SetSecurity(ctx, elem.SecurityDefinitionRequest.SecurityReqID, elem)
	}

	// Set security count
	k.SetSecurityCount(ctx, genState.SecurityCount)
	// Set all the orderMassStatus
	for _, elem := range genState.OrderMassStatusList {
		k.SetOrderMassStatus(ctx, elem.OrderMassStatusRequest.MassStatusReqID, elem)
	}

	// Set orderMassStatus count
	k.SetOrderMassStatusCount(ctx, genState.OrderMassStatusCount)
	// Set all the tradingSession
	for _, elem := range genState.TradingSessionList {
		k.SetTradingSession(ctx, elem.TradingSessionStatusRequest.TradSesReqID, elem)
	}

	// Set tradingSession count
	k.SetTradingSessionCount(ctx, genState.TradingSessionCount)
	// Set all the tradingSessionList
	for _, elem := range genState.TradingSessionListList {
		k.SetTradingSessionList(ctx, elem.TradingSessionListRequest.TradSesReqID, elem)
	}

	// Set tradingSessionList count
	k.SetTradingSessionListCount(ctx, genState.TradingSessionListCount)
	// Set all the securityList
	for _, elem := range genState.SecurityListList {
		k.SetSecurityList(ctx, elem)
	}

	// Set securityList count
	k.SetSecurityListCount(ctx, genState.SecurityListCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AccountList = k.GetAllAccount(ctx)
	genesis.SessionsList = k.GetAllSessions(ctx)
	genesis.SessionsCount = k.GetSessionsCount(ctx)

	genesis.QuoteList = k.GetAllQuote(ctx)
	genesis.TradeCaptureList = k.GetAllTradeCapture(ctx)
	genesis.TradeCaptureCount = k.GetTradeCaptureCount(ctx)
	genesis.MarketDataList = k.GetAllMarketData(ctx)
	genesis.MarketDataCount = k.GetMarketDataCount(ctx)
	genesis.SecurityList = k.GetAllSecurity(ctx)
	genesis.SecurityCount = k.GetSecurityCount(ctx)
	genesis.OrderMassStatusList = k.GetAllOrderMassStatus(ctx)
	genesis.OrderMassStatusCount = k.GetOrderMassStatusCount(ctx)
	genesis.TradingSessionList = k.GetAllTradingSession(ctx)
	genesis.TradingSessionCount = k.GetTradingSessionCount(ctx)
	genesis.TradingSessionListList = k.GetAllTradingSessionList(ctx)
	genesis.TradingSessionListCount = k.GetTradingSessionListCount(ctx)
	genesis.SecurityListList = k.GetAllSecurityList(ctx)
	genesis.SecurityListCount = k.GetSecurityListCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
