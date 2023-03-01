package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListOrdersExecutionReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-orders-execution-report",
		Short: "list all orders_execution-report",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrdersExecutionReportRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrdersExecutionReportAll(context.Background(), params)
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

func CmdShowOrdersExecutionReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders-execution-report [sessionID]",
		Short: "shows an orders_execution-report",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionID := args[0]

			params := &types.QueryGetOrdersExecutionReportRequest{
				SessionID: argSessionID,
			}

			res, err := queryClient.OrdersExecutionReport(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
