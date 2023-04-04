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
		//Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in marketIdentificationCode
	marketIdentificationCodeIndexMap := make(map[string]struct{})
	for _, elem := range gs.MarketIdentificationCodeList {
		index := string(GetMarketIdentificationCodeKey(elem.MIC))
		if _, ok := marketIdentificationCodeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for marketIdentificationCode")
		}
		marketIdentificationCodeIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
