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

func CmdSecurityListResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-list-response [session-id] [security-req-id] [security-response-id] [security-request-result] [tot-no-related-sym] [last-fragment] [no-related-sym] [no-underlyings] [currency] [no-legs] [round-lot] [min-trade-vol] [trading-session-id] [trading-session-sub-id] [expiration-cycle] [text] [encoded-text-len] [encoded-text]",
		Short: "Broadcast message security-list-response",
		Args:  cobra.ExactArgs(18),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityResponseID := args[2]
			argSecurityRequestResult := args[3]
			argTotNoRelatedSym := args[4]
			argLastFragment := args[5]
			argNoRelatedSym := args[6]
			argNoUnderlyings := args[7]
			argCurrency := args[8]
			argNoLegs := args[9]
			argRoundLot := args[10]
			argMinTradeVol := args[11]
			argTradingSessionID := args[12]
			argTradingSessionSubID := args[13]
			argExpirationCycle := args[14]
			argText := args[15]
			argEncodedTextLen := args[16]
			argEncodedText := args[17]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityListResponse(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityResponseID,
				argSecurityRequestResult,
				argTotNoRelatedSym,
				argLastFragment,
				argNoRelatedSym,
				argNoUnderlyings,
				argCurrency,
				argNoLegs,
				argRoundLot,
				argMinTradeVol,
				argTradingSessionID,
				argTradingSessionSubID,
				argExpirationCycle,
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
