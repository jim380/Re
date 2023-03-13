package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MarketIdentificationCodeList: []MarketIdentificationCode{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in marketIdentificationCode
	marketIdentificationCodeIdMap := make(map[uint64]bool)
	marketIdentificationCodeCount := gs.GetMarketIdentificationCodeCount()
	for _, elem := range gs.MarketIdentificationCodeList {
		if _, ok := marketIdentificationCodeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for marketIdentificationCode")
		}
		if elem.Id >= marketIdentificationCodeCount {
			return fmt.Errorf("marketIdentificationCode id should be lower or equal than the last id")
		}
		marketIdentificationCodeIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
