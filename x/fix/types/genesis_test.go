package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/jim380/Re/x/fix/types"
)

func TestGenesisState_Validate(t *testing.T) {
    for _, tc := range []struct {
    		desc          string
    		genState      *types.GenesisState
    		valid bool
    } {
        {
            desc:     "default is valid",
            genState: types.DefaultGenesis(),
            valid:    true,
        },
        {
            desc:     "valid genesis state",
            genState: &types.GenesisState{
            	
                SecurityList: []types.Security{
	{
		Id: 0,
	},
	{
		Id: 1,
	},
},
SecurityCount: 2,
// this line is used by starport scaffolding # types/genesis/validField
            },
            valid:    true,
        },
        {
	desc:     "duplicated security",
	genState: &types.GenesisState{
		SecurityList: []types.Security{
			{
				Id: 0,
			},
			{
				Id: 0,
			},
		},
	},
	valid:    false,
},
{
	desc:     "invalid security count",
	genState: &types.GenesisState{
		SecurityList: []types.Security{
			{
				Id: 1,
			},
		},
		SecurityCount: 0,
	},
	valid:    false,
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