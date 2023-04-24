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
		Use:   "trade-capture-report [trade-report-id] [trade-report-trans-type] [trade-report-type] [trade-request-id] [trd-type] [trd-sub-type] [side] [order-qty] [last-qty] [last-px] [gross-trade-amt] [exec-id] [order-id] [trade-id] [orig-trade-id] [symbol] [security-id] [security-id-source] [trade-date] [transact-time] [settl-type] [settl-date]",
		Short: "Broadcast message trade-capture-report",
		Args:  cobra.ExactArgs(22),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradeReportID := args[0]
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

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradeCaptureReport(
				clientCtx.GetFromAddress().String(),
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
		Use:   "trade-capture-report-acknowledgement [trade-report-id] [trade-request-id] [trade-request-type] [trade-id] [secondary-trade-id] [trade-report-type] [trd-type] [trd-sub-type] [exec-type] [trade-report-ref-id] [secondary-trade-report-id] [trade-report-status] [trade-trans-type] [trade-report-reject-reason] [text]",
		Short: "Broadcast message trade-capture-report-acknowledgement",
		Args:  cobra.ExactArgs(15),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradeReportID := args[0]
			argTradeRequestID := args[1]
			argTradeRequestType := args[2]
			argTradeID := args[3]
			argSecondaryTradeID := args[4]
			argTradeReportType := args[5]
			argTrdType := args[6]
			argTrdSubType := args[7]
			argExecType := args[8]
			argTradeReportRefID := args[9]
			argSecondaryTradeReportID := args[10]
			argTradeReportStatus := args[11]
			argTradeTransType := args[12]
			argTradeReportRejectReason := args[13]
			argText := args[14]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradeCaptureReportAcknowledgement(
				clientCtx.GetFromAddress().String(),
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
		Use:   "trade-capture-report-rejection [trade-report-id] [trade-request-id] [trade-request-type] [trade-report-reject-reason] [trade-report-reject-ref-id] [text]",
		Short: "Broadcast message trade-capture-report-rejection",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradeReportID := args[0]
			argTradeRequestID := args[1]
			argTradeRequestType := args[2]
			argTradeReportRejectReason := args[3]
			argTradeReportRejectRefID := args[4]
			argText := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTradeCaptureReportRejection(
				clientCtx.GetFromAddress().String(),
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
