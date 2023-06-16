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

func CmdSecurityTypesRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-types-request [session-id] [security-req-id] [text] [trading-session-id] [trading-session-sub-id] [product] [security-type] [security-sub-type]",
		Short: "Broadcast message security-types-request",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argText := args[2]
			argTradingSessionID := args[3]
			argTradingSessionSubID := args[4]
			argProduct := args[5]
			argSecurityType := args[6]
			argSecuritySubType := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityTypesRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argText,
				argTradingSessionID,
				argTradingSessionSubID,
				argProduct,
				argSecurityType,
				argSecuritySubType,
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
