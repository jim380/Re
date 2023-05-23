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

func CmdOrderMassStatusRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-mass-status-request [session-id] [mass-status-req-type] [clord-id] [account] [symbol] [security-id] [trading-session-id]",
		Short: "Broadcast message order-mass-status-request",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argMassStatusReqType := args[1]
			argClOrdID := args[2]
			argAccount := args[3]
			argSymbol := args[4]
			argSecurityID := args[5]
			argTradingSessionID := args[6]

			// GenerateRandomString function uniquely generates MassStatusReqID for every Order Mass Status Request
			argMassStatusReqID, _ := types.GenerateRandomString(types.MassStatusReqID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderMassStatusRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argMassStatusReqID,
				argMassStatusReqType,
				argClOrdID,
				argAccount,
				argSymbol,
				argSecurityID,
				argTradingSessionID,
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

func CmdOrderMassStatusReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-mass-status-report [session-id] [clord-id] [mass-status-req-id] [account] [symbol] [security-id] [trading-session-id] [ord-status] [exec-type] [ord-qty] [last-px] [cum-qty] [avg-px] [leaves-qty]",
		Short: "Broadcast message order-mass-status-report",
		Args:  cobra.ExactArgs(14),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			
			argSessionID := args[0]
			argClOrdID := args[1]
			argMassStatusReqID := args[2]
			argAccount := args[3]
			argSymbol := args[4]
			argSecurityID := args[5]
			argTradingSessionID := args[6]
			argOrdStatus := args[7]
			argExecType := args[8]
			argOrdQty := args[9]
			argLastPx := args[10]
			argCumQty := args[11]
			argAvgPx := args[12]
			argLeavesQty := args[13]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderMassStatusReport(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argClOrdID,
				argMassStatusReqID,
				argAccount,
				argSymbol,
				argSecurityID,
				argTradingSessionID,
				argOrdStatus,
				argExecType,
				argOrdQty,
				argLastPx,
				argCumQty,
				argAvgPx,
				argLeavesQty,
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

func CmdOrderMassStatusRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-mass-status-request-reject [session-id] [ref-seq-id] [rej-code] [text]",
		Short: "Broadcast message order-mass-status-request-reject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argRefSeqID := args[1]
			argRejCode := args[2]
			argText := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderMassStatusRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argRefSeqID,
				argRejCode,
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
