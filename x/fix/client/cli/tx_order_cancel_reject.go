package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdOrderCancelReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-cancel-reject [sessionID] [OrderID] [OrigClOrdID] [ClOrdID] [CxlRejReason] [CxlRejResponseTo]",
		Short: "Broadcast message order-cancel-reject",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			argOrderID := args[1]

			argOrigClOrdID := args[2]

			argClOrdID := args[3]

			argCxlRejReason, err := cast.ToInt64E(args[4])
			if err != nil {
				return err
			}

			argCxlRejResponseTo, err := cast.ToInt64E(args[5])
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderCancelReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argOrderID,
				argOrigClOrdID,
				argClOrdID,
				argCxlRejReason,
				argCxlRejResponseTo,
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
