package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListOrders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-orders",
		Short: "list all orders",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrdersRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrdersAll(context.Background(), params)
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

func CmdShowOrders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders [sessionID]",
		Short: "shows a orders",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionID := args[0]

			params := &types.QueryGetOrdersRequest{
				SessionID: argSessionID,
			}

			res, err := queryClient.Orders(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListOrdersCancelRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-orders-cancel-request",
		Short: "list all ordersCancelRequest",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrdersCancelRequestRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrdersCancelRequestAll(context.Background(), params)
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

func CmdShowOrdersCancelRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders-cancel-request [sessionID]",
		Short: "shows an ordersCancelRequest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionID := args[0]

			params := &types.QueryGetOrdersCancelRequestRequest{
				SessionID: argSessionID,
			}

			res, err := queryClient.OrdersCancelRequest(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

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

func CmdListOrdersCancelReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-orders-cancel-reject",
		Short: "list all orders-cancel-reject",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrdersCancelRejectRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrdersCancelRejectAll(context.Background(), params)
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

func CmdShowOrdersCancelReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders-cancel-reject [sessionID]",
		Short: "shows a orders-cancel-reject",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionID := args[0]

			params := &types.QueryGetOrdersCancelRejectRequest{
				SessionID: argSessionID,
			}

			res, err := queryClient.OrdersCancelReject(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
