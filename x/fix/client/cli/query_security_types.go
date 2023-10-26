package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListSecurityTypes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-security-types [chainID]",
		Short: "list all security-types",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSecurityTypesRequest{
				ChainID:    argChainID,
				Pagination: pageReq,
			}

			res, err := queryClient.SecurityTypesAll(context.Background(), params)
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

func CmdShowSecurityTypes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-security-types [chainID] [SecurityReqID]",
		Short: "shows a security-types",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argSecurityReqID := args[1]

			params := &types.QueryGetSecurityTypesRequest{
				ChainID:       argChainID,
				SecurityReqID: argSecurityReqID,
			}

			res, err := queryClient.SecurityTypes(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
