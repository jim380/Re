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
			// TODO
			// add session id
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

func CmdTradingSessionStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-status [trad-ses-req-id] [trading-session-id] [trad-ses-status] [trad-ses-status-rej-reason] [trad-ses-start-time] [trad-ses-open-time] [trad-ses-pre-close-time] [trad-ses-close-time] [trad-ses-end-time] [total-volume-traded] [trad-ses-high-px] [trad-ses-low-px] [security-id] [security-id-source] [symbol] [security-exchange] [market-segment-id] [market-id]",
		Short: "Broadcast message trading-session-status",
		Args:  cobra.ExactArgs(18),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// TODO
			// add session id
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

func CmdTradingSessionStatusRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-status-request-reject [ref-seq-num] [ref-msg-type] [session-reject-reason] [text]",
		Short: "Broadcast message trading-session-status-request-reject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// TODO
			// add session id
			argRefSeqNum := args[0]
			argRefMsgType := args[1]
			argSessionRejectReason, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argText := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionStatusRequestReject(
				clientCtx.GetFromAddress().String(),
				argRefSeqNum,
				argRefMsgType,
				argSessionRejectReason,
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
