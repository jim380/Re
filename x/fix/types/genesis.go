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
		AccountList:       []Account{},
		SessionsList:      []Sessions{},
		SessionRejectList: []SessionReject{},
		SessionLogoutList: []SessionLogout{},
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
		if _, ok := sessionsIdMap[elem.SessionName]; ok {
			return fmt.Errorf("duplicated id for sessions")
		}
		sessionName, _ := strconv.ParseUint(elem.SessionName, 10, 64)
		if sessionName >= sessionsCount {
			return fmt.Errorf("sessions id should be lower or equal than the last id")
		}
		sessionsIdMap[elem.SessionName] = true
	}
	// Check for duplicated ID in sessionReject
	sessionRejectIdMap := make(map[string]bool)
	sessionRejectCount := gs.GetSessionRejectCount()
	for _, elem := range gs.SessionRejectList {
		if _, ok := sessionRejectIdMap[elem.SessionName]; ok {
			return fmt.Errorf("duplicated id for sessionReject")
		}
		sessionName, _ := strconv.ParseUint(elem.SessionName, 10, 64)
		if sessionName >= sessionRejectCount {
			return fmt.Errorf("sessionReject id should be lower or equal than the last id")
		}
		sessionRejectIdMap[elem.SessionName] = true
	}
	// Check for duplicated sessionName in sessionLogout
	sessionLogoutIdMap := make(map[string]bool)
	sessionLogoutCount := gs.GetSessionLogoutCount()
	for _, elem := range gs.SessionLogoutList {
		if _, ok := sessionLogoutIdMap[elem.SessionName]; ok {
			return fmt.Errorf("duplicated id for sessionLogout")
		}
		sessionName, _ := strconv.ParseUint(elem.SessionName, 10, 64)
		if sessionName >= sessionLogoutCount {
			return fmt.Errorf("sessionLogout id should be lower or equal than the last id")
		}
		sessionLogoutIdMap[elem.SessionName] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
