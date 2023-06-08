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

func CmdSecurityListRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-list-request-reject [session-id] [security-req-id] [security-list-request-type] [security-request-result] [text] [encoded-text-len] [encoded-text]",
		Short: "Broadcast message security-list-request-reject",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityListRequestType := args[2]
			argSecurityRequestResult := args[3]
			argText := args[4]
			argEncodedTextLen := args[5]
			argEncodedText := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityListRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityListRequestType,
				argSecurityRequestResult,
				argText,
				argEncodedTextLen,
				argEncodedText,
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
