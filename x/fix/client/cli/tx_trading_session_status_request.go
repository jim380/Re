package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTradingSessionStatusRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-status-request [trad-ses-req-id] [trading-session-id] [trading-session-sub-id] [market-id] [subscription-request] [security-id] [security-id-source] [symbol] [security-exchange] [market-segment-id] [trad-ses-req-type] [trad-ses-sub-req-type] [trad-ses-mode] [trading-session-date] [trading-session-time] [trading-session-sub-time] [expiration-date]",
		Short: "Broadcast message trading-session-status-request",
		Args:  cobra.ExactArgs(17),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradSesReqID := args[0]
			argTradingSessionID := args[1]
			argTradingSessionSubID := args[2]
			argMarketID := args[3]
			argSubscriptionRequest := args[4]
			argSecurityID := args[5]
			argSecurityIDSource := args[6]
			argSymbol := args[7]
			argSecurityExchange := args[8]
			argMarketSegmentID := args[9]
			argTradSesReqType, err := cast.ToInt32E(args[10])
			if err != nil {
				return err
			}
			argTradSesSubReqType, err := cast.ToInt32E(args[11])
			if err != nil {
				return err
			}
			argTradSesMode, err := cast.ToInt32E(args[12])
			if err != nil {
				return err
			}
			argTradingSessionDate := args[13]
			argTradingSessionTime := args[14]
			argTradingSessionSubTime := args[15]
			argExpirationDate := args[16]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionStatusRequest(
				clientCtx.GetFromAddress().String(),
				argTradSesReqID,
				argTradingSessionID,
				argTradingSessionSubID,
				argMarketID,
				argSubscriptionRequest,
				argSecurityID,
				argSecurityIDSource,
				argSymbol,
				argSecurityExchange,
				argMarketSegmentID,
				argTradSesReqType,
				argTradSesSubReqType,
				argTradSesMode,
				argTradingSessionDate,
				argTradingSessionTime,
				argTradingSessionSubTime,
				argExpirationDate,
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
