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
		Use:   "security-list-request [sessionID] [securityReqID] [securityListRequestType] [noUnderlyings] [noLegs] [currency] [text] [encodedTextLen] [encodedText] [tradingSessionID] [tradingSessionSubID] [subscriptionRequestType]",
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

func CmdSecurityListResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-list-response [sessionID] [securityReqID] [securityResponseID] [securityRequestResult] [totNoRelatedSym] [lastFragment] [noRelatedSym] [noUnderlyings] [currency] [noLegs] [roundLot] [minTradeVol] [tradingSessionID] [tradingSessionSubID] [expirationCycle] [text] [encodedTextLen] [encodedText]",
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

func CmdSecurityListRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-list-request-reject [sessionID] [securityReqID] [securityListRequestType] [securityRequestResult] [text] [encodedTextLen] [encodedText]",
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
