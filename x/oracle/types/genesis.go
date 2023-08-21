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
		MultiChainTxOracleList: []MultiChainTxOracle{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in multiChainTxOracle
	multiChainTxOracleIdMap := make(map[string]bool)
	multiChainTxOracleCount := gs.GetMultiChainTxOracleCount()
	for _, elem := range gs.MultiChainTxOracleList {
		if _, ok := multiChainTxOracleIdMap[elem.OracleId]; ok {
			return fmt.Errorf("duplicated id for multiChainTxOracle")
		}
		oracleID, _ := strconv.ParseUint(elem.OracleId, 10, 64)
		if oracleID >= multiChainTxOracleCount {
			return fmt.Errorf("multiChainTxOracle id should be lower or equal than the last id")
		}
		multiChainTxOracleIdMap[elem.OracleId] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
