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

func CmdSecurityListRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-list-request [session-id] [security-req-id] [security-list-request-type] [no-underlyings] [no-legs] [currency] [text] [encoded-text-len] [encoded-text] [trading-session-id] [trading-session-sub-id] [subscription-request-type]",
		Short: "Broadcast message security-list-request",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityListRequestType := args[2]
			argNoUnderlyings := args[3]
			argNoLegs := args[4]
			argCurrency := args[5]
			argText := args[6]
			argEncodedTextLen := args[7]
			argEncodedText := args[8]
			argTradingSessionID := args[9]
			argTradingSessionSubID := args[10]
			argSubscriptionRequestType := args[11]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityListRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityListRequestType,
				argNoUnderlyings,
				argNoLegs,
				argCurrency,
				argText,
				argEncodedTextLen,
				argEncodedText,
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
