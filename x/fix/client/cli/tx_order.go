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

func CmdNewOrderSingle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-order-single [sessionID] [ClOrdID] [Symbol] [Side] [OrderQty] [OrdType] [Price] [TimeInForce] [Text]",
		Short: "Broadcast message new-order-single",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argSessionID := args[0]

			argClOrdID := args[1]

			argSymbol := args[2]

			argSide, err := cast.ToInt64E(args[3])
			if err != nil {
				return err
			}

			argOrderQty := args[4]

			argOrdType, err := cast.ToInt64E(args[5])
			if err != nil {
				return err
			}

			argPrice := args[6]

			argTimeInForce, err := cast.ToInt64E(args[7])
			if err != nil {
				return err
			}

			argText := args[8]

			msg := types.NewMsgNewOrderSingle(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argClOrdID,
				argSymbol,
				argSide,
				argOrderQty,
				argOrdType,
				argPrice,
				argTimeInForce,
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

func CmdOrderCancelRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-cancel-request [session-name] [OrigClOrdID] [ClOrdID]",
		Short: "Broadcast message order-cancel-request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argSession := args[0]

			argOrigClOrdID := args[1]

			argclOrdID := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderCancelRequest(
				clientCtx.GetFromAddress().String(),
				argSession,
				argOrigClOrdID,
				argclOrdID,
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

func CmdOrderExecutionReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-execution-report [sessionID] [ClOrdID] [OrderID] [ExecID] [OrdStatus] [ExecType] [Symbol] [Side] [OrderQty] [Price] [TimeInForce] [LastPx] [LastQty] [LeavesQty] [CumQty] [AvgPx] [Text]",
		Short: "Broadcast message order_execution-report",
		Args:  cobra.ExactArgs(17),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			argClOrdID := args[1]

			argOrderID := args[2]

			argExecID := args[3]

			argOrdStatus := args[4]

			argExecType := args[5]

			argSymbol := args[6]

			argSide, err := cast.ToInt64E(args[7])
			if err != nil {
				return err
			}

			argOrderQty, err := cast.ToInt64E(args[8])
			if err != nil {
				return err
			}

			argPrice, err := cast.ToInt64E(args[9])
			if err != nil {
				return err
			}

			argTimeInForce, err := cast.ToInt64E(args[10])
			if err != nil {
				return err
			}

			argLastPx, err := cast.ToInt64E(args[11])
			if err != nil {
				return err
			}

			argLastQty, err := cast.ToInt64E(args[12])
			if err != nil {
				return err
			}

			argLeavesQty, err := cast.ToInt64E(args[13])
			if err != nil {
				return err
			}

			argCumQty, err := cast.ToInt64E(args[14])
			if err != nil {
				return err
			}

			argAvgPx, err := cast.ToInt64E(args[15])
			if err != nil {
				return err
			}

			argText := args[16]

			msg := types.NewMsgOrderExecutionReport(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argClOrdID,
				argOrderID,
				argExecID,
				argOrdStatus,
				argExecType,
				argSymbol,
				argSide,
				argOrderQty,
				argPrice,
				argTimeInForce,
				argLastPx,
				argLastQty,
				argLeavesQty,
				argCumQty,
				argAvgPx,
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

func CmdOrderCancelReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-cancel-reject [sessionID] [OrderID] [OrigClOrdID] [ClOrdID] [CxlRejReason] [CxlRejResponseTo]",
		Short: "Broadcast message order-cancel-reject",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			argOrderID := args[1]

			argOrigClOrdID := args[2]

			argClOrdID := args[3]

			argCxlRejReason, err := cast.ToInt64E(args[4])
			if err != nil {
				return err
			}

			argCxlRejResponseTo, err := cast.ToInt64E(args[5])
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderCancelReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argOrderID,
				argOrigClOrdID,
				argClOrdID,
				argCxlRejReason,
				argCxlRejResponseTo,
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
