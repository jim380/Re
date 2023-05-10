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

func CmdTradeCaptureReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trade-capture-report [session-id] [trade-report-id] [trade-report-trans-type] [trade-report-type] [trade-request-id] [trd-type] [trd-sub-type] [side] [order-qty] [last-qty] [last-px] [gross-trade-amt] [exec-id] [order-id] [trade-id] [orig-trade-id] [symbol] [security-id] [security-id-source] [trade-date] [transact-time] [settl-type] [settl-date]",
		Short: "Broadcast message trade-capture-report",
		Args:  cobra.ExactArgs(22),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSessionID := args[0]

			argTradeReportTransType := args[1]

			argTradeReportType := args[2]

			argTradeRequestID := args[3]

			argTrdType := args[4]

			argTrdSubType := args[5]

			argSide := args[6]

			argOrderQty := args[7]

			argLastQty := args[8]

			argLastPx := args[9]

			argGrossTradeAmt := args[10]

			argExecID := args[11]

			argOrderID := args[12]

			argTradeID := args[13]

			argOrigTradeID := args[14]

			argSymbol := args[15]

			argSecurityID := args[16]

			argSecurityIDSource := args[17]

			argTradeDate := args[18]

			argTransactTime := args[19]

			argSettlType := args[20]

			argSettlDate := args[21]

			// GenerateRandomString function uniquely generates tradeReportID for every Trade Capture Report
			argTradeReportID, _ := types.GenerateRandomString(types.TradeReportID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradeCaptureReport(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argTradeReportID,
				argTradeReportTransType,
				argTradeReportType,
				argTradeRequestID,
				argTrdType,
				argTrdSubType,
				argSide,
				argOrderQty,
				argLastQty,
				argLastPx,
				argGrossTradeAmt,
				argExecID,
				argOrderID,
				argTradeID,
				argOrigTradeID,
				argSymbol,
				argSecurityID,
				argSecurityIDSource,
				argTradeDate,
				argTransactTime,
				argSettlType,
				argSettlDate,
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

func CmdTradeCaptureReportAcknowledgement() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trade-capture-report-acknowledgement [session-id] [trade-report-id] [trade-request-id] [trade-request-type] [trade-id] [secondary-trade-id] [trade-report-type] [trd-type] [trd-sub-type] [exec-type] [trade-report-ref-id] [secondary-trade-report-id] [trade-report-status] [trade-trans-type] [trade-report-reject-reason] [text]",
		Short: "Broadcast message trade-capture-report-acknowledgement",
		Args:  cobra.ExactArgs(16),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSessionID := args[0]
			argTradeReportID := args[1]
			argTradeRequestID := args[2]
			argTradeRequestType := args[3]
			argTradeID := args[4]
			argSecondaryTradeID := args[5]
			argTradeReportType := args[6]
			argTrdType := args[7]
			argTrdSubType := args[8]
			argExecType := args[9]
			argTradeReportRefID := args[10]
			argSecondaryTradeReportID := args[11]
			argTradeReportStatus := args[12]
			argTradeTransType := args[13]

			argTradeReportRejectReason, err := strconv.ParseInt(args[14], 10, 32)
			if err != nil {
				panic(err)
			}

			argText := args[15]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradeCaptureReportAcknowledgement(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argTradeReportID,
				argTradeRequestID,
				argTradeRequestType,
				argTradeID,
				argSecondaryTradeID,
				argTradeReportType,
				argTrdType,
				argTrdSubType,
				argExecType,
				argTradeReportRefID,
				argSecondaryTradeReportID,
				argTradeReportStatus,
				argTradeTransType,
				argTradeReportRejectReason,
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

func CmdTradeCaptureReportRejection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trade-capture-report-rejection [session-id] [trade-report-id] [trade-request-id] [trade-request-type] [trade-report-reject-reason] [trade-report-reject-ref-id] [text]",
		Short: "Broadcast message trade-capture-report-rejection",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSessionID := args[0]
			argTradeReportID := args[1]
			argTradeRequestID := args[2]
			argTradeRequestType := args[3]

			argTradeReportRejectReason, err := strconv.ParseInt(args[4], 10, 32)
			if err != nil {
				panic(err)
			}

			argTradeReportRejectRefID := args[5]
			argText := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradeCaptureReportRejection(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argTradeReportID,
				argTradeRequestID,
				argTradeRequestType,
				argTradeReportRejectReason,
				argTradeReportRejectRefID,
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
