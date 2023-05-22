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

func CmdOrderMassStatusRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-mass-status-request [session-id] [mass-status-req-id] [mass-status-req-type] [cl-ord-id] [account] [symbol] [security-id] [trading-session-id]",
		Short: "Broadcast message order-mass-status-request",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argMassStatusReqID := args[1]
			argMassStatusReqType := args[2]
			argClOrdID := args[3]
			argAccount := args[4]
			argSymbol := args[5]
			argSecurityID := args[6]
			argTradingSessionID := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderMassStatusRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argMassStatusReqID,
				argMassStatusReqType,
				argClOrdID,
				argAccount,
				argSymbol,
				argSecurityID,
				argTradingSessionID,
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
