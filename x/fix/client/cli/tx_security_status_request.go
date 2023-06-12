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
		Use:   "security-status-request [session-id] [security-status-req-id] [instrument] [no-underlyings] [underlying-instrument] [no-legs] [instrument-leg] [currency] [subscription-request-type] [trading-session-id] [trading-session-sub-id]",
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
