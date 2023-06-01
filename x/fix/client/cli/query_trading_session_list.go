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
		Use:   "list-trading-session-list",
		Short: "list all trading-session-list",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTradingSessionListRequest{
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
		Use:   "show-trading-session-list [tradSesReqID ]",
		Short: "shows a trading-session-list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTradSesReqID := args[0]

			params := &types.QueryGetTradingSessionListRequest{
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
