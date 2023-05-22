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

func CmdOrderMassStatusReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-mass-status-report [session-id] [cl-ord-id] [mass-status-req-id] [account] [symbol] [security-id] [trading-session-id] [ord-status] [exec-type] [ord-qty] [last-px] [cum-qty] [avg-px] [leaves-qty]",
		Short: "Broadcast message order-mass-status-report",
		Args:  cobra.ExactArgs(14),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argClOrdID := args[1]
			argMassStatusReqID := args[2]
			argAccount := args[3]
			argSymbol := args[4]
			argSecurityID := args[5]
			argTradingSessionID := args[6]
			argOrdStatus := args[7]
			argExecType := args[8]
			argOrdQty := args[9]
			argLastPx := args[10]
			argCumQty := args[11]
			argAvgPx := args[12]
			argLeavesQty := args[13]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderMassStatusReport(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argClOrdID,
				argMassStatusReqID,
				argAccount,
				argSymbol,
				argSecurityID,
				argTradingSessionID,
				argOrdStatus,
				argExecType,
				argOrdQty,
				argLastPx,
				argCumQty,
				argAvgPx,
				argLeavesQty,
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
