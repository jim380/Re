package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListSecurity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-security-definition [chainID]",
		Short: "list all security definition",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSecurityRequest{
				ChainID:    argChainID,
				Pagination: pageReq,
			}

			res, err := queryClient.SecurityAll(context.Background(), params)
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

func CmdShowSecurity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-security-definition [chainID] [securityReqID]",
		Short: "shows a security definition",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argSecurityReqID := args[1]

			params := &types.QueryGetSecurityRequest{
				ChainID:       argChainID,
				SecurityReqID: argSecurityReqID,
			}

			res, err := queryClient.Security(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
