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
		Use:   "market-data-request [sessionID] [subscriptionRequestType] [marketDepth] [mdUpdateType] [noRelatedSym] [symbol]",
		Short: "Broadcast message market-data-request",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]

			// GenerateRandomString function uniquely generates MdReqID for every Market Data Request
			argMdReqID, _ := types.GenerateRandomString(types.MdReqID)

			argSubscriptionRequestType, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				panic(err)
			}

			argMarketDepth, err := strconv.ParseInt(args[2], 10, 32)
			if err != nil {
				panic(err)
			}

			argMdUpdateType, err := strconv.ParseInt(args[3], 10, 32)
			if err != nil {
				panic(err)
			}

			argNoRelatedSym, err := strconv.ParseInt(args[4], 10, 32)
			if err != nil {
				panic(err)
			}

			argSymbol := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarketDataRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
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
		Use:   "market-data-snapshot-full-refresh [sessionID] [mdReqID] [symbol] [noMdEntries] [mdUpdateAction] [mdEntryType] [mdEntryPx] [mdEntrySize]",
		Short: "Broadcast message market-data-snapshot-full-refresh",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			argMdReqID := args[1]

			argSymbol := args[2]

			argNoMDEntries, err := strconv.ParseInt(args[3], 10, 32)
			if err != nil {
				panic(err)
			}

			mdEntries := []*types.MDEntry{}

			argMdUpdateAction, err := strconv.ParseInt(args[4], 10, 32)
			if err != nil {
				panic(err)
			}

			argMdEntryType, err := strconv.ParseInt(args[5], 10, 32)
			if err != nil {
				panic(err)
			}

			argMdEntryPx := args[6]
			argMdEntrySize := args[7]

			mdEntries = append(mdEntries, &types.MDEntry{
				MdUpdateAction: argMdUpdateAction,
				MdEntryType:    argMdEntryType,
				MdEntryPx:      argMdEntryPx,
				MdEntrySize:    argMdEntrySize,
			})

			msg := types.NewMsgMarketDataSnapshotFullRefresh(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argMdReqID,
				argSymbol,
				argNoMDEntries,
				mdEntries,
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
		Use:   "market-data-incremental [mdReqID] [noMdEntries]",
		Short: "Broadcast message market-data-incremental",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMdReqID := args[0]

			argNoMDEntries, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				panic(err)
			}

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
		Use:   "market-data-request-reject [sessionID] [mdReqID] [mdReqRejReason] [text]",
		Short: "Broadcast message market-data-request-reject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			argMdReqID := args[1]

			argMdReqRejReason, err := strconv.ParseInt(args[2], 10, 32)
			if err != nil {
				panic(err)
			}

			argText := args[3]

			msg := types.NewMsgMarketDataRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
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
