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

func CmdMarketDataRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-data-request [md-req-id] [subscription-request-type] [market-depth] [md-update-type] [no-related-sym] [symbol]",
		Short: "Broadcast message market-data-request",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMdReqID := args[0]
			argSubscriptionRequestType := args[1]
			argMarketDepth := args[2]
			argMdUpdateType := args[3]
			argNoRelatedSym := args[4]
			argSymbol := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarketDataRequest(
				clientCtx.GetFromAddress().String(),
				argMdReqID,
				argSubscriptionRequestType,
				argMarketDepth,
				argMdUpdateType,
				argNoRelatedSym,
				argSymbol,
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


func CmdMarketDataSnapshotFullRefresh() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-data-snapshot-full-refresh [md-req-id] [symbol] [no-md-entries]",
		Short: "Broadcast message market-data-snapshot-full-refresh",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMdReqID := args[0]
			argSymbol := args[1]
			argNoMDEntries := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarketDataSnapshotFullRefresh(
				clientCtx.GetFromAddress().String(),
				argMdReqID,
				argSymbol,
				argNoMDEntries,
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


func CmdMarketDataIncremental() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-data-incremental [md-req-id] [no-md-entries]",
		Short: "Broadcast message market-data-incremental",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMdReqID := args[0]
			argNoMDEntries := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarketDataIncremental(
				clientCtx.GetFromAddress().String(),
				argMdReqID,
				argNoMDEntries,
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


func CmdMarketDataRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-data-request-reject [md-req-id] [md-req-rej-reason] [text]",
		Short: "Broadcast message market-data-request-reject",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMdReqID := args[0]
			argMdReqRejReason := args[1]
			argText := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarketDataRequestReject(
				clientCtx.GetFromAddress().String(),
				argMdReqID,
				argMdReqRejReason,
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
