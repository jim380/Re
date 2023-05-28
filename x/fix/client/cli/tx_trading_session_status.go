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

func CmdTradingSessionStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-status [trad-ses-req-id] [trading-session-id] [trad-ses-status] [trad-ses-status-rej-reason] [trad-ses-start-time] [trad-ses-open-time] [trad-ses-pre-close-time] [trad-ses-close-time] [trad-ses-end-time] [total-volume-traded] [trad-ses-high-px] [trad-ses-low-px] [security-id] [security-id-source] [symbol] [security-exchange] [market-segment-id] [market-id]",
		Short: "Broadcast message trading-session-status",
		Args:  cobra.ExactArgs(18),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradSesReqID := args[0]
			argTradingSessionID := args[1]
			argTradSesStatus, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argTradSesStatusRejReason, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argTradSesStartTime := args[4]
			argTradSesOpenTime := args[5]
			argTradSesPreCloseTime := args[6]
			argTradSesCloseTime := args[7]
			argTradSesEndTime := args[8]
			argTotalVolumeTraded, err := cast.ToInt32E(args[9])
			if err != nil {
				return err
			}
			argTradSesHighPx := args[10]
			argTradSesLowPx := args[11]
			argSecurityID := args[12]
			argSecurityIDSource := args[13]
			argSymbol := args[14]
			argSecurityExchange := args[15]
			argMarketSegmentID := args[16]
			argMarketID := args[17]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionStatus(
				clientCtx.GetFromAddress().String(),
				argTradSesReqID,
				argTradingSessionID,
				argTradSesStatus,
				argTradSesStatusRejReason,
				argTradSesStartTime,
				argTradSesOpenTime,
				argTradSesPreCloseTime,
				argTradSesCloseTime,
				argTradSesEndTime,
				argTotalVolumeTraded,
				argTradSesHighPx,
				argTradSesLowPx,
				argSecurityID,
				argSecurityIDSource,
				argSymbol,
				argSecurityExchange,
				argMarketSegmentID,
				argMarketID,
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
