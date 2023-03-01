package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

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
