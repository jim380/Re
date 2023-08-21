package types_test

import (
	"testing"

	"github.com/jim380/Re/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				MultiChainTxOracleList: []types.MultiChainTxOracle{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				MultiChainTxOracleCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated multiChainTxOracle",
			genState: &types.GenesisState{
				MultiChainTxOracleList: []types.MultiChainTxOracle{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid multiChainTxOracle count",
			genState: &types.GenesisState{
				MultiChainTxOracleList: []types.MultiChainTxOracle{
					{
						Id: 1,
					},
				},
				MultiChainTxOracleCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
