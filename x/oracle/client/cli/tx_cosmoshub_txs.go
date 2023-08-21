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

func CmdCosmoshubTxs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cosmoshub-txs [address]",
		Short: "Broadcast message cosmoshub-txs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// GenerateRandomString function uniquely generates oracleID for every cosmoshub-txs initiated
			oracleID, _ := helpers.GenerateRandomString(constants.OracleID)

			argOracleId := oracleID

			argAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCosmoshubTxs(
				clientCtx.GetFromAddress().String(),
				argOracleId,
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
