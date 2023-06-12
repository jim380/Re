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

func CmdSecurityStatusRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-status-request-reject [session-id] [security-status-req-id] [security-reject-reason] [text]",
		Short: "Broadcast message security-status-request-reject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityStatusReqID := args[1]
			argSecurityRejectReason := args[2]
			argText := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityStatusRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityStatusReqID,
				argSecurityRejectReason,
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
