package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/oracle/types"
	"github.com/spf13/cobra"
)

func CmdListMultiChainTxOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-multi-chain-tx-oracle",
		Short: "list all multi-chain-tx-oracle",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMultiChainTxOracleRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MultiChainTxOracleAll(context.Background(), params)
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

func CmdShowMultiChainTxOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-multi-chain-tx-oracle [oracle-id]",
		Short: "shows a multi-chain-tx-oracle",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOracleId := args[0]

			params := &types.QueryGetMultiChainTxOracleRequest{
				OracleId: argOracleId,
			}

			res, err := queryClient.MultiChainTxOracle(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
