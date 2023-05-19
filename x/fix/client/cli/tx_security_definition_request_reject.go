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

func CmdSecurityDefinitionRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-definition-request-reject [session-id] [security-req-id] [security-request-result] [security-request-error] [security-request-error-code] [text]",
		Short: "Broadcast message security-definition-request-reject",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityRequestResult := args[2]
			argSecurityRequestError := args[3]
			argSecurityRequestErrorCode := args[4]
			argText := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityDefinitionRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityRequestResult,
				argSecurityRequestError,
				argSecurityRequestErrorCode,
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
