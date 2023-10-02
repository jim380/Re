package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/utils/helpers"
	"github.com/jim380/Re/x/oracle/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEthereumTxs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ethereum-txs [address]",
		Short: "Broadcast message ethereum-txs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argAddress := args[0]

			// GenerateRandomString function uniquely generates oracleID for every ethereum-txs initiated
			argOracleID, _ := helpers.GenerateRandomString(constants.OracleID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEthereumTxs(
				clientCtx.GetFromAddress().String(),
				argOracleID,
				argAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
