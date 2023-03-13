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
		MarketIdentificationCodeList: []MarketIdentificationCode{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in marketIdentificationCode
	marketIdentificationCodeIdMap := make(map[string]bool)
	marketIdentificationCodeCount := gs.GetMarketIdentificationCodeCount()
	for _, elem := range gs.MarketIdentificationCodeList {
		if _, ok := marketIdentificationCodeIdMap[elem.MIC]; ok {
			return fmt.Errorf("duplicated id for marketIdentificationCode")
		}
		mic, _ := strconv.ParseUint(elem.MIC, 10, 64)
		if mic >= marketIdentificationCodeCount {
			return fmt.Errorf("marketIdentificationCode id should be lower or equal than the last id")
		}
		marketIdentificationCodeIdMap[elem.MIC] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
