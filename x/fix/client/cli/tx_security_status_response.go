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

func CmdSecurityStatusResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-status-response [session-id] [security-status-req-id] [no-underlyings] [underlying-instrument] [no-legs] [instrument-leg] [currency] [trading-session-id] [trading-session-sub-id] [unsolicited-indicator] [security-trading-status] [financial-status] [corporate-action] [halt-reason] [in-view-of-common] [due-to-related] [buy-volume] [sell-volume] [high-px] [low-px] [last-px] [transact-time] [adjustment] [text]",
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
