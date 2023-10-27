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

func CmdTradingSessionListRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-list-request [sessionID] [tradSesReqID] [tradingSessionID] [tradingSessionSubID] [securityExchange] [tradSesMethod] [tradSesMode] [subscriptionRequestType]",
		Short: "Broadcast message trading-session-list-request",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argTradSesReqID := args[1]
			argTradingSessionID := args[2]
			argTradingSessionSubID := args[3]
			argSecurityExchange := args[4]
			argTradSesMethod := args[5]
			argTradSesMode := args[6]
			argSubscriptionRequestType := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionListRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argTradSesReqID,
				argTradingSessionID,
				argTradingSessionSubID,
				argSecurityExchange,
				argTradSesMethod,
				argTradSesMode,
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

func CmdTradingSessionListResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-list-response [sessionID] [tradSesReqID] [noTradingSessions] [tradingSessionID] [tradingSessionSubID] [securityExchange] [tradSesMethod] [tradSesMode] [unsolicitedIndicator] [tradSesStatus] [tradSesStatusRejReason] [tradSesStartTime] [tradSesOpenTime] [tradSesPreCloseTime] [tradSesCloseTime] [tradSesEndTime]",
		Short: "Broadcast message trading-session-list-response",
		Args:  cobra.ExactArgs(16),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argTradSesReqID := args[1]
			argNoTradingSessions := args[2]
			argTradingSessionID := args[3]
			argTradingSessionSubID := args[4]
			argSecurityExchange := args[5]
			argTradSesMethod := args[6]
			argTradSesMode := args[7]
			argUnsolicitedIndicator := args[8]
			argTradSesStatus := args[9]
			argTradSesStatusRejReason := args[10]
			argTradSesStartTime := args[11]
			argTradSesOpenTime := args[12]
			argTradSesPreCloseTime := args[13]
			argTradSesCloseTime := args[14]
			argTradSesEndTime := args[15]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionListResponse(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argTradSesReqID,
				argNoTradingSessions,
				argTradingSessionID,
				argTradingSessionSubID,
				argSecurityExchange,
				argTradSesMethod,
				argTradSesMode,
				argUnsolicitedIndicator,
				argTradSesStatus,
				argTradSesStatusRejReason,
				argTradSesStartTime,
				argTradSesOpenTime,
				argTradSesPreCloseTime,
				argTradSesCloseTime,
				argTradSesEndTime,
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

func CmdTradingSessionListRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-session-list-request-reject [sessionID] [tradSesReqID] [tradSesStatus] [tradSesStatusRejReason] [text]",
		Short: "Broadcast message trading-session-list-request-reject",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argTradSesReqID := args[1]
			argTradSesStatus := args[2]
			argTradSesStatusRejReason := args[3]
			argText := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradingSessionListRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argTradSesReqID,
				argTradSesStatus,
				argTradSesStatusRejReason,
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
