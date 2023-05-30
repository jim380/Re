package types

import (
	"fmt"
	"log"
	"strconv"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccountList:               []Account{},
		SessionsList:              []Sessions{},
		SessionRejectList:         []SessionReject{},
		SessionLogoutList:         []SessionLogout{},
		OrdersList:                []Orders{},
		OrdersCancelRequestList:   []OrdersCancelRequest{},
		OrdersCancelRejectList:    []OrdersCancelReject{},
		OrdersExecutionReportList: []OrdersExecutionReport{},
		QuoteList:                 []Quote{},
		TradeCaptureList:          []TradeCapture{},
		MarketDataList:            []MarketData{},
		SecurityList:              []Security{},
		OrderMassStatusList:       []OrderMassStatus{},
		TradingSessionList:        []TradingSession{},
		TradingSessionListList:    []TradingSessionList{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in account
	accountDidMap := make(map[string]bool)
	accountCount := gs.GetAccountCount()
	for _, elem := range gs.AccountList {
		if _, ok := accountDidMap[elem.Did]; ok {
			return fmt.Errorf("duplicated id for account")
		}
		did, err := strconv.ParseUint(elem.Did, 10, 64)
		if err != nil {
			// Handle the error if the conversion fails
			log.Printf("Error occurred: %v", err)
			return err
		}
		if did >= accountCount {
			return fmt.Errorf("account id should be lower or equal than the last id")
		}
		accountDidMap[elem.Did] = true
	}
	// Check for duplicated ID in sessions
	sessionsIDMap := make(map[string]bool)
	sessionsCount := gs.GetSessionsCount()
	for _, elem := range gs.SessionsList {
		if _, ok := sessionsIDMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for sessions")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= sessionsCount {
			return fmt.Errorf("sessions id should be lower or equal than the last id")
		}
		sessionsIDMap[elem.SessionID] = true
	}
	// Check for duplicated ID in sessionReject
	sessionRejectIDMap := make(map[string]bool)
	sessionRejectCount := gs.GetSessionRejectCount()
	for _, elem := range gs.SessionRejectList {
		if _, ok := sessionRejectIDMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for sessionReject")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= sessionRejectCount {
			return fmt.Errorf("sessionReject id should be lower or equal than the last id")
		}
		sessionRejectIDMap[elem.SessionID] = true
	}
	// Check for duplicated sessionName in sessionLogout
	sessionLogoutIDMap := make(map[string]bool)
	sessionLogoutCount := gs.GetSessionLogoutCount()
	for _, elem := range gs.SessionLogoutList {
		if _, ok := sessionLogoutIDMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for sessionLogout")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= sessionLogoutCount {
			return fmt.Errorf("sessionLogout id should be lower or equal than the last id")
		}
		sessionLogoutIDMap[elem.SessionID] = true
	}
	// Check for duplicated ID in orders
	ordersIDMap := make(map[string]bool)
	ordersCount := gs.GetOrdersCount()
	for _, elem := range gs.OrdersList {
		if _, ok := ordersIDMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for orders")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= ordersCount {
			return fmt.Errorf("orders id should be lower or equal than the last id")
		}
		ordersIDMap[elem.SessionID] = true
	}
	// Check for duplicated ID in ordersCancelRequest
	ordersCancelRequestIDMap := make(map[string]bool)
	ordersCancelRequestCount := gs.GetOrdersCancelRequestCount()
	for _, elem := range gs.OrdersCancelRequestList {
		if _, ok := ordersCancelRequestIDMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for ordersCancelRequest")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= ordersCancelRequestCount {
			return fmt.Errorf("ordersCancelRequest id should be lower or equal than the last id")
		}
		ordersCancelRequestIDMap[elem.SessionID] = true
	}
	// Check for duplicated ID in ordersCancelReject
	ordersCancelRejectIDMap := make(map[string]bool)
	ordersCancelRejectCount := gs.GetOrdersCancelRejectCount()
	for _, elem := range gs.OrdersCancelRejectList {
		if _, ok := ordersCancelRejectIDMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for ordersCancelReject")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= ordersCancelRejectCount {
			return fmt.Errorf("ordersCancelReject id should be lower or equal than the last id")
		}
		ordersCancelRejectIDMap[elem.SessionID] = true
	}
	// Check for duplicated ID in ordersExecutionReport
	ordersExecutionReportIDMap := make(map[string]bool)
	ordersExecutionReportCount := gs.GetOrdersExecutionReportCount()
	for _, elem := range gs.OrdersExecutionReportList {
		if _, ok := ordersExecutionReportIDMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for ordersExecutionReport")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= ordersExecutionReportCount {
			return fmt.Errorf("ordersExecutionReport id should be lower or equal than the last id")
		}
		ordersExecutionReportIDMap[elem.SessionID] = true
	}
	// Check for duplicated ID in quote
	quoteIDMap := make(map[string]bool)
	quoteCount := gs.GetQuoteCount()
	for _, elem := range gs.QuoteList {
		if _, ok := quoteIDMap[elem.QuoteRequest.QuoteReqID]; ok {
			return fmt.Errorf("duplicated id for quote")
		}
		quoteReqID, _ := strconv.ParseUint(elem.QuoteRequest.QuoteReqID, 10, 64)
		if quoteReqID >= quoteCount {
			return fmt.Errorf("quote id should be lower or equal than the last id")
		}
		quoteIDMap[elem.QuoteRequest.QuoteReqID] = true
	}
	// Check for duplicated ID in marketData
	marketDataIDMap := make(map[string]bool)
	marketDataCount := gs.GetMarketDataCount()
	for _, elem := range gs.MarketDataList {
		if _, ok := marketDataIDMap[elem.MarketDataRequest.MdReqID]; ok {
			return fmt.Errorf("duplicated id for marketData")
		}
		mdReqID, _ := strconv.ParseUint(elem.MarketDataRequest.MdReqID, 10, 64)
		if mdReqID >= marketDataCount {
			return fmt.Errorf("marketData id should be lower or equal than the last id")
		}
		marketDataIDMap[elem.MarketDataRequest.MdReqID] = true
	}
	// Check for duplicated ID in tradeCapture
	tradeCaptureIDMap := make(map[string]bool)
	tradeCaptureCount := gs.GetTradeCaptureCount()
	for _, elem := range gs.TradeCaptureList {
		if _, ok := tradeCaptureIDMap[elem.TradeCaptureReport.TradeReportID]; ok {
			return fmt.Errorf("duplicated id for tradeCapture")
		}
		tradeReportID, _ := strconv.ParseUint(elem.TradeCaptureReport.TradeReportID, 10, 64)
		if tradeReportID >= tradeCaptureCount {
			return fmt.Errorf("tradeCapture id should be lower or equal than the last id")
		}
		tradeCaptureIDMap[elem.TradeCaptureReport.TradeReportID] = true
	}
	// Check for duplicated ID in security
	securityIDMap := make(map[string]bool)
	securityCount := gs.GetSecurityCount()
	for _, elem := range gs.SecurityList {
		if _, ok := securityIDMap[elem.SecurityDefinitionRequest.SecurityReqID]; ok {
			return fmt.Errorf("duplicated id for security")
		}
		securityReqID, _ := strconv.ParseUint(elem.SecurityDefinitionRequest.SecurityReqID, 10, 64)
		if securityReqID >= securityCount {
			return fmt.Errorf("security id should be lower or equal than the last id")
		}
		securityIDMap[elem.SecurityDefinitionRequest.SecurityReqID] = true
	}
	// Check for duplicated ID in orderMassStatus
	orderMassStatusIDMap := make(map[string]bool)
	orderMassStatusCount := gs.GetOrderMassStatusCount()
	for _, elem := range gs.OrderMassStatusList {
		if _, ok := orderMassStatusIDMap[elem.OrderMassStatusRequest.MassStatusReqID]; ok {
			return fmt.Errorf("duplicated id for orderMassStatus")
		}
		massStatusReqID, _ := strconv.ParseUint(elem.OrderMassStatusRequest.MassStatusReqID, 10, 64)
		if massStatusReqID >= orderMassStatusCount {
			return fmt.Errorf("orderMassStatus id should be lower or equal than the last id")
		}
		orderMassStatusIDMap[elem.OrderMassStatusRequest.MassStatusReqID] = true
	}
	// Check for duplicated ID in tradingSession
	tradingSessionIDMap := make(map[string]bool)
	tradingSessionCount := gs.GetTradingSessionCount()
	for _, elem := range gs.TradingSessionList {
		if _, ok := tradingSessionIDMap[elem.TradingSessionStatusRequest.TradSesReqID]; ok {
			return fmt.Errorf("duplicated id for tradingSession")
		}
		tradSesReqID, _ := strconv.ParseUint(elem.TradingSessionStatusRequest.TradSesReqID, 10, 64)
		if tradSesReqID >= tradingSessionCount {
			return fmt.Errorf("tradingSession id should be lower or equal than the last id")
		}
		tradingSessionIDMap[elem.TradingSessionStatusRequest.TradSesReqID] = true
	}
	// Check for duplicated ID in tradingSessionList
	tradingSessionListIdMap := make(map[uint64]bool)
	tradingSessionListCount := gs.GetTradingSessionListCount()
	for _, elem := range gs.TradingSessionListList {
		if _, ok := tradingSessionListIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for tradingSessionList")
		}
		if elem.Id >= tradingSessionListCount {
			return fmt.Errorf("tradingSessionList id should be lower or equal than the last id")
		}
		tradingSessionListIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
