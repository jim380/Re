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
		Use:   "list-orders [chainID]",
		Short: "list all orders",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]

			params := &types.QueryAllOrdersRequest{
				ChainID:    argChainID,
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
		Use:   "show-order [chainID] [clOrdID]",
		Short: "shows an order",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argClOrdID := args[1]

			params := &types.QueryGetOrdersRequest{
				ChainID: argChainID,
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
		Use:   "show-orders-by-address [chainID] [address]",
		Short: "show orders by address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argAddress := args[1]

			params := &types.QueryGetOrdersByAddressRequest{
				ChainID:    argChainID,
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
		Use:   "list-orders-cancel-request [chianID]",
		Short: "list all ordersCancelRequest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]

			params := &types.QueryAllOrdersCancelRequestRequest{
				ChainID:    argChainID,
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
		Use:   "show-orders-cancel-request [chainID] [clOrdID]",
		Short: "shows an ordersCancelRequest",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argClOrdID := args[1]

			params := &types.QueryGetOrdersCancelRequestRequest{
				ChainID: argChainID,
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
		Use:   "list-orders-execution-report [chainID]",
		Short: "list all orders_execution-report",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]

			params := &types.QueryAllOrdersExecutionReportRequest{
				ChainID:    argChainID,
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
		Use:   "show-orders-execution-report [chainID] [clOrdID]",
		Short: "shows an orders_execution-report",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argClOrdID := args[1]

			params := &types.QueryGetOrdersExecutionReportRequest{
				ChainID: argChainID,
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
		Use:   "show-orders-execution-report-by-address [chainID] [address]",
		Short: "shows an orders execution report by address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argAddress := args[1]

			params := &types.QueryGetOrdersExecutionReportByAddressRequest{
				ChainID:    argChainID,
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
		Use:   "list-orders-cancel-reject [chainID]",
		Short: "list all orders-cancel-reject",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrdersCancelRejectRequest{
				ChainID:    argChainID,
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
		Use:   "show-orders-cancel-reject [chainID] [clOrdID]",
		Short: "shows a orders-cancel-reject",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argClOrdID := args[1]

			params := &types.QueryGetOrdersCancelRejectRequest{
				ChainID: argChainID,
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
