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

func CmdSecurityTypesResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-types-response [session-id] [security-req-id] [security-response-id] [security-response-type] [tot-no-security-types] [last-fragment] [no-security-types] [security-type] [security-sub-type] [product] [c-fi-code] [text] [trading-session-id] [trading-session-sub-id] [subscription-request-type]",
		Short: "Broadcast message security-types-response",
		Args:  cobra.ExactArgs(15),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityResponseID := args[2]
			argSecurityResponseType := args[3]
			argTotNoSecurityTypes := args[4]
			argLastFragment := args[5]
			argNoSecurityTypes := args[6]
			argSecurityType := args[7]
			argSecuritySubType := args[8]
			argProduct := args[9]
			argCFICode := args[10]
			argText := args[11]
			argTradingSessionID := args[12]
			argTradingSessionSubID := args[13]
			argSubscriptionRequestType := args[14]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityTypesResponse(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityResponseID,
				argSecurityResponseType,
				argTotNoSecurityTypes,
				argLastFragment,
				argNoSecurityTypes,
				argSecurityType,
				argSecuritySubType,
				argProduct,
				argCFICode,
				argText,
				argTradingSessionID,
				argTradingSessionSubID,
				argSubscriptionRequestType,
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
