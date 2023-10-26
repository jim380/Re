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
		Use:   "list-quote [chainID]",
		Short: "list all quote",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllQuoteRequest{
				ChainID:    argChainID,
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
		Use:   "show-quote [chainID] [quoteReqID]",
		Short: "shows a quote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argQuoteReqID := args[1]

			params := &types.QueryGetQuoteRequest{
				ChainID:    argChainID,
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
		Use:   "show-session-quotes [chainID] [sessionID]",
		Short: "shows a quote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argSessionID := args[1]

			params := &types.QuerySessionByIDQuoteRequest{
				ChainID:   argChainID,
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
