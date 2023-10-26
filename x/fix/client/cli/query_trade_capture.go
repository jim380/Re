package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListTradeCapture() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-trade-capture [chainID]",
		Short: "list all trade-capture",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTradeCaptureRequest{
				ChainID:    argChainID,
				Pagination: pageReq,
			}

			res, err := queryClient.TradeCaptureAll(context.Background(), params)
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

func CmdShowTradeCapture() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-trade-capture [chainID] [tradeReportID]",
		Short: "shows a trade-capture",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argTradeReportID := args[1]

			params := &types.QueryGetTradeCaptureRequest{
				ChainID:       argChainID,
				TradeReportID: argTradeReportID,
			}

			res, err := queryClient.TradeCapture(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
