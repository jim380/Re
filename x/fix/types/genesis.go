package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccountList: []Account{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in account
	accountIdMap := make(map[uint64]bool)
	accountCount := gs.GetAccountCount()
	for _, elem := range gs.AccountList {
		if _, ok := accountIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for account")
		}
		if elem.Id >= accountCount {
			return fmt.Errorf("account id should be lower or equal than the last id")
		}
		accountIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
