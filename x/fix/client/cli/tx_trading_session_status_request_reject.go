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

func CmdTradingSessionStatusRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-status-request-reject [ref-seq-num] [ref-msg-type] [session-reject-reason] [text]",
		Short: "Broadcast message trading-session-status-request-reject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRefSeqNum := args[0]
			argRefMsgType := args[1]
			argSessionRejectReason, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argText := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionStatusRequestReject(
				clientCtx.GetFromAddress().String(),
				argRefSeqNum,
				argRefMsgType,
				argSessionRejectReason,
				argText,
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
