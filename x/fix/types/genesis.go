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
		AccountList:  []Account{},
		SessionsList: []Sessions{},
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
	sessionsIdMap := make(map[uint64]bool)
	sessionsCount := gs.GetSessionsCount()
	for _, elem := range gs.SessionsList {
		if _, ok := sessionsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for sessions")
		}
		if elem.Id >= sessionsCount {
			return fmt.Errorf("sessions id should be lower or equal than the last id")
		}
		sessionsIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
