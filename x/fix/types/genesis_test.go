package types_test

import (
	"testing"

	"github.com/jim380/Re/x/fix/types"
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

				OrdersExecutionReportList: []types.OrdersExecutionReport{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				OrdersExecutionReportCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated ordersExecutionReport",
			genState: &types.GenesisState{
				OrdersExecutionReportList: []types.OrdersExecutionReport{
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
			desc: "invalid ordersExecutionReport count",
			genState: &types.GenesisState{
				OrdersExecutionReportList: []types.OrdersExecutionReport{
					{
						Id: 1,
					},
				},
				OrdersExecutionReportCount: 0,
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
