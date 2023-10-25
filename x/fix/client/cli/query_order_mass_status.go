package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListOrderMassStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-order-mass-status [chainID]",
		Short: "list all order-mass-status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]

			params := &types.QueryAllOrderMassStatusRequest{
				ChainID:    argChainID,
				Pagination: pageReq,
			}

			res, err := queryClient.OrderMassStatusAll(context.Background(), params)
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

func CmdShowOrderMassStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-order-mass-status [chainID] [massStatusReqID]",
		Short: "shows a order-mass-status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argMassStatusReqID := args[1]

			params := &types.QueryGetOrderMassStatusRequest{
				ChainID:         argChainID,
				MassStatusReqID: argMassStatusReqID,
			}

			res, err := queryClient.OrderMassStatus(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
