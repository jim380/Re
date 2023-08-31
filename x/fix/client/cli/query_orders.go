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
		Use:   "show-order [clOrdID]",
		Short: "shows an order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argClOrdID := args[0]

			params := &types.QueryGetOrdersRequest{
				ClOrdID: argClOrdID,
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

func CmdShowOrdersByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders-by-address [address]",
		Short: "show orders by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetOrdersByAddressRequest{
				Address:    argAddress,
				Pagination: pageReq,
			}

			res, err := queryClient.OrdersByAddress(context.Background(), params)
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
		Use:   "show-orders-cancel-request [clOrdID]",
		Short: "shows an ordersCancelRequest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argClOrdID := args[0]

			params := &types.QueryGetOrdersCancelRequestRequest{
				ClOrdID: argClOrdID,
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
		Use:   "show-orders-execution-report [clOrdID]",
		Short: "shows an orders_execution-report",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argClOrdID := args[0]

			params := &types.QueryGetOrdersExecutionReportRequest{
				ClOrdID: argClOrdID,
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

func CmdShowOrdersExecutionReportByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders-execution-report-by-address [address]",
		Short: "shows an orders execution report by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetOrdersExecutionReportByAddressRequest{
				Address:    argAddress,
				Pagination: pageReq,
			}

			res, err := queryClient.OrdersExecutionReportByAddress(context.Background(), params)
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
		Use:   "show-orders-cancel-reject [clOrdID]",
		Short: "shows a orders-cancel-reject",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argClOrdID := args[0]

			params := &types.QueryGetOrdersCancelRejectRequest{
				ClOrdID: argClOrdID,
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
