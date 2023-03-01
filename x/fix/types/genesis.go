package types

import (
	"fmt"
	"strconv"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccountList:             []Account{},
		SessionsList:            []Sessions{},
		SessionRejectList:       []SessionReject{},
		SessionLogoutList:       []SessionLogout{},
		OrdersList:              []Orders{},
		OrdersCancelRequestList: []OrdersCancelRequest{},
		OrdersCancelRejectList:  []OrdersCancelReject{},
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
		}
		if did >= accountCount {
			return fmt.Errorf("account id should be lower or equal than the last id")
		}
		accountDidMap[elem.Did] = true
	}
	// Check for duplicated ID in sessions
	sessionsIdMap := make(map[string]bool)
	sessionsCount := gs.GetSessionsCount()
	for _, elem := range gs.SessionsList {
		if _, ok := sessionsIdMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for sessions")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= sessionsCount {
			return fmt.Errorf("sessions id should be lower or equal than the last id")
		}
		sessionsIdMap[elem.SessionID] = true
	}
	// Check for duplicated ID in sessionReject
	sessionRejectIdMap := make(map[string]bool)
	sessionRejectCount := gs.GetSessionRejectCount()
	for _, elem := range gs.SessionRejectList {
		if _, ok := sessionRejectIdMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for sessionReject")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= sessionRejectCount {
			return fmt.Errorf("sessionReject id should be lower or equal than the last id")
		}
		sessionRejectIdMap[elem.SessionID] = true
	}
	// Check for duplicated sessionName in sessionLogout
	sessionLogoutIdMap := make(map[string]bool)
	sessionLogoutCount := gs.GetSessionLogoutCount()
	for _, elem := range gs.SessionLogoutList {
		if _, ok := sessionLogoutIdMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for sessionLogout")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= sessionLogoutCount {
			return fmt.Errorf("sessionLogout id should be lower or equal than the last id")
		}
		sessionLogoutIdMap[elem.SessionID] = true
	}
	// Check for duplicated ID in orders
	ordersIdMap := make(map[string]bool)
	ordersCount := gs.GetOrdersCount()
	for _, elem := range gs.OrdersList {
		if _, ok := ordersIdMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for orders")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= ordersCount {
			return fmt.Errorf("orders id should be lower or equal than the last id")
		}
		ordersIdMap[elem.SessionID] = true
	}
	// Check for duplicated ID in ordersCancelRequest
	ordersCancelRequestIdMap := make(map[string]bool)
	ordersCancelRequestCount := gs.GetOrdersCancelRequestCount()
	for _, elem := range gs.OrdersCancelRequestList {
		if _, ok := ordersCancelRequestIdMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for ordersCancelRequest")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= ordersCancelRequestCount {
			return fmt.Errorf("ordersCancelRequest id should be lower or equal than the last id")
		}
		ordersCancelRequestIdMap[elem.SessionID] = true
	}
	// Check for duplicated ID in ordersCancelReject
	ordersCancelRejectIdMap := make(map[string]bool)
	ordersCancelRejectCount := gs.GetOrdersCancelRejectCount()
	for _, elem := range gs.OrdersCancelRejectList {
		if _, ok := ordersCancelRejectIdMap[elem.SessionID]; ok {
			return fmt.Errorf("duplicated id for ordersCancelReject")
		}
		sessionID, _ := strconv.ParseUint(elem.SessionID, 10, 64)
		if sessionID >= ordersCancelRejectCount {
			return fmt.Errorf("ordersCancelReject id should be lower or equal than the last id")
		}
		ordersCancelRejectIdMap[elem.SessionID] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
