package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListMarketData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-market-data [chainID]",
		Short: "list all market-data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]

			params := &types.QueryAllMarketDataRequest{
				ChainID:    argChainID,
				Pagination: pageReq,
			}

			res, err := queryClient.MarketDataAll(context.Background(), params)
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

func CmdShowMarketData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-market-data [chainID] [mdReqID]",
		Short: "shows a market-data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argmdReqID := args[1]

			params := &types.QueryGetMarketDataRequest{
				ChainID: argChainID,
				MdReqID: argmdReqID,
			}

			res, err := queryClient.MarketData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
