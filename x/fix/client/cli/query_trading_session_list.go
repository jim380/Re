package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListTradingSessionList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-trading-session-list [chainID]",
		Short: "list all trading-session-list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTradingSessionListRequest{
				ChainID:    argChainID,
				Pagination: pageReq,
			}

			res, err := queryClient.TradingSessionListAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowTradingSessionList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-trading-session-list [chainID] [tradSesReqID]",
		Short: "shows a trading-session-list",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argTradSesReqID := args[1]

			params := &types.QueryGetTradingSessionListRequest{
				ChainID:      argChainID,
				TradSesReqID: argTradSesReqID,
			}

			res, err := queryClient.TradingSessionList(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
