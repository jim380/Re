package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListQuote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-quote",
		Short: "list all quote",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllQuoteRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.QuoteAll(context.Background(), params)
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

func CmdShowQuote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-quote [quoteReqID]",
		Short: "shows a quote",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argQuoteReqID := args[0]

			params := &types.QueryGetQuoteRequest{
				QuoteReqID: argQuoteReqID,
			}

			res, err := queryClient.Quote(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowQuotesBySessionID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-session-quotes [sessionID]",
		Short: "shows a quote",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionID := args[0]

			params := &types.QuerySessionByIDQuoteRequest{
				SessionID: argSessionID,
			}

			res, err := queryClient.QuotesBySessionID(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
