package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListTradingSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-trading-session",
		Short: "list all trading-session",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTradingSessionRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TradingSessionAll(context.Background(), params)
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

func CmdShowTradingSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-trading-session [tradSesReqID]",
		Short: "shows a trading-session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTradSesReqID := args[0]

			params := &types.QueryGetTradingSessionRequest{
				TradSesReqID: argTradSesReqID,
			}

			res, err := queryClient.TradingSession(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
