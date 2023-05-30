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
		Use:   "trading-session-status-request [sessionID] [tradingSessionID] [tradingSessionSubID] [marketID] [subscriptionRequest] [securityID] [securityIDSource] [symbol] [securityExchange] [marketSegmentID] [tradSesReqType] [tradSesSubReqType] [tradSesMode] [tradingSessionDate] [tradingSessionTime] [tradingSessionSubTime] [expirationDate]",
		Short: "Broadcast message trading-session-status-request",
		Args:  cobra.ExactArgs(17),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSesssionID := args[0]
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

			// GenerateRandomString function uniquely generates TradSesReqID for every Trading Session Status Request
			argTradSesReqID, _ := types.GenerateRandomString(types.TradSesReqID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionStatusRequest(
				clientCtx.GetFromAddress().String(),
				argSesssionID,
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
		Use:   "trading-session-status [sessionID] [tradSesReqID] [tradingSessionID] [tradSesStatus] [tradSesStatusRejReason] [tradSesStartTime] [tradSesOpenTime] [tradSesPreCloseTime] [tradSesCloseTime] [tradSesEndTime] [totalVolumeTraded] [tradSesHighPx] [tradSesLowPx] [securityID] [securityIDSource] [symbol] [securityExchange] [marketSegmentID] [marketID]",
		Short: "Broadcast message trading-session-status",
		Args:  cobra.ExactArgs(19),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argTradSesReqID := args[1]
			argTradingSessionID := args[2]
			argTradSesStatus, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argTradSesStatusRejReason, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}
			argTradSesStartTime := args[5]
			argTradSesOpenTime := args[6]
			argTradSesPreCloseTime := args[7]
			argTradSesCloseTime := args[8]
			argTradSesEndTime := args[9]
			argTotalVolumeTraded, err := cast.ToInt32E(args[10])
			if err != nil {
				return err
			}
			argTradSesHighPx := args[11]
			argTradSesLowPx := args[12]
			argSecurityID := args[13]
			argSecurityIDSource := args[14]
			argSymbol := args[15]
			argSecurityExchange := args[16]
			argMarketSegmentID := args[17]
			argMarketID := args[18]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionStatus(
				clientCtx.GetFromAddress().String(),
				argSessionID,
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
		Use:   "trading-session-status-request-reject [sessionID] [refSeqNum] [refMsgType] [sessionRejectReason] [text]",
		Short: "Broadcast message trading-session-status-request-reject",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argRefSeqNum := args[1]
			argRefMsgType := args[2]
			argSessionRejectReason, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argText := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionStatusRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
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
