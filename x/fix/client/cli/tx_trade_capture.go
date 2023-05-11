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
		Use:   "trade-capture-report [session-id] [trade-report-id] [trade-report-trans-type] [trade-report-type] [trd-type] [trd-sub-type] [side] [order-qty] [last-qty] [last-px] [gross-trade-amt] [exec-id] [order-id] [trade-id] [orig-trade-id] [symbol] [security-id] [security-id-source] [trade-date] [transact-time] [settl-type] [settl-date]",
		Short: "Broadcast message trade-capture-report",
		Args:  cobra.ExactArgs(21),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSessionID := args[0]

			argTradeReportTransType := args[1]

			argTradeReportType := args[2]

			argTrdType := args[3]

			argTrdSubType := args[4]

			argSide := args[5]

			argOrderQty := args[6]

			argLastQty := args[7]

			argLastPx := args[8]

			argGrossTradeAmt := args[9]

			argExecID := args[10]

			argOrderID := args[11]

			argTradeID := args[12]

			argOrigTradeID := args[13]

			argSymbol := args[14]

			argSecurityID := args[15]

			argSecurityIDSource := args[16]

			argTradeDate := args[17]

			argTransactTime := args[18]

			argSettlType := args[19]

			argSettlDate := args[20]

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
		Use:   "trade-capture-report-acknowledgement [session-id] [trade-report-id] [trade-id] [secondary-trade-id] [trade-report-type] [trd-type] [trd-sub-type] [exec-type] [trade-report-ref-id] [secondary-trade-report-id] [trade-report-status] [trade-trans-type] [trade-report-reject-reason] [text]",
		Short: "Broadcast message trade-capture-report-acknowledgement",
		Args:  cobra.ExactArgs(14),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSessionID := args[0]
			argTradeReportID := args[1]
			argTradeID := args[2]
			argSecondaryTradeID := args[3]
			argTradeReportType := args[4]
			argTrdType := args[5]
			argTrdSubType := args[6]
			argExecType := args[7]
			argTradeReportRefID := args[8]
			argSecondaryTradeReportID := args[9]
			argTradeReportStatus := args[10]
			argTradeTransType := args[11]

			argTradeReportRejectReason, err := strconv.ParseInt(args[12], 10, 32)
			if err != nil {
				panic(err)
			}

			argText := args[13]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradeCaptureReportAcknowledgement(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argTradeReportID,
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
		Use:   "trade-capture-report-rejection [session-id] [trade-report-id] [trade-report-reject-reason] [trade-report-reject-ref-id] [text]",
		Short: "Broadcast message trade-capture-report-rejection",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSessionID := args[0]

			argTradeReportID := args[1]

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
