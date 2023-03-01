package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdOrderCancelRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-cancel-request [session-name] [OrigClOrdID] [ClOrdID]",
		Short: "Broadcast message order-cancel-request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSession := args[0]

			argOrigClOrdID := args[1]

			argclOrdID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderCancelRequest(
				clientCtx.GetFromAddress().String(),
				argSession,
				argOrigClOrdID,
				argclOrdID,
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
