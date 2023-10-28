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

func CmdSecurityStatusRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-status-request [sessionID] [securityStatusReqID] [instrument] [noUnderlyings] [underlyingInstrument] [noLegs] [instrumentLeg] [currency] [subscriptionRequestType] [tradingSessionID] [tradingSessionSubID]",
		Short: "Broadcast message security-status-request",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityStatusReqID := args[1]
			argInstrument := args[2]
			argNoUnderlyings := args[3]
			argUnderlyingInstrument := args[4]
			argNoLegs := args[5]
			argInstrumentLeg := args[6]
			argCurrency := args[7]
			argSubscriptionRequestType := args[8]
			argTradingSessionID := args[9]
			argTradingSessionSubID := args[10]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityStatusRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityStatusReqID,
				argInstrument,
				argNoUnderlyings,
				argUnderlyingInstrument,
				argNoLegs,
				argInstrumentLeg,
				argCurrency,
				argSubscriptionRequestType,
				argTradingSessionID,
				argTradingSessionSubID,
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

func CmdSecurityStatusResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-status-response [sessionID] [securityStatusReqID] [noUnderlyings] [underlyingInstrument] [noLegs] [instrumentLeg] [currency] [tradingSessionID] [tradingSessionSubID] [unsolicitedIndicator] [securityTradingStatus] [financialStatus] [corporateAction] [haltReason] [inViewOfCommon] [dueToRelated] [buyVolume] [sellVolume] [highPx] [lowPx] [lastPx] [transactTime] [adjustment] [text]",
		Short: "Broadcast message security-status-response",
		Args:  cobra.ExactArgs(24),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityStatusReqID := args[1]
			argNoUnderlyings := args[2]
			argUnderlyingInstrument := args[3]
			argNoLegs := args[4]
			argInstrumentLeg := args[5]
			argCurrency := args[6]
			argTradingSessionID := args[7]
			argTradingSessionSubID := args[8]
			argUnsolicitedIndicator := args[9]
			argSecurityTradingStatus := args[10]
			argFinancialStatus := args[11]
			argCorporateAction := args[12]
			argHaltReason := args[13]
			argInViewOfCommon := args[14]
			argDueToRelated := args[15]
			argBuyVolume := args[16]
			argSellVolume := args[17]
			argHighPx := args[18]
			argLowPx := args[19]
			argLastPx := args[20]
			argTransactTime := args[21]
			argAdjustment := args[22]
			argText := args[23]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityStatusResponse(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityStatusReqID,
				argNoUnderlyings,
				argUnderlyingInstrument,
				argNoLegs,
				argInstrumentLeg,
				argCurrency,
				argTradingSessionID,
				argTradingSessionSubID,
				argUnsolicitedIndicator,
				argSecurityTradingStatus,
				argFinancialStatus,
				argCorporateAction,
				argHaltReason,
				argInViewOfCommon,
				argDueToRelated,
				argBuyVolume,
				argSellVolume,
				argHighPx,
				argLowPx,
				argLastPx,
				argTransactTime,
				argAdjustment,
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

func CmdSecurityStatusRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-status-request-reject [sessionID] [securityStatusReqID] [securityRejectReason] [text]",
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
