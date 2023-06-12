package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListSecurityStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-security-status",
		Short: "list all security-status",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSecurityStatusRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SecurityStatusAll(context.Background(), params)
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

func CmdShowSecurityStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-security-status [securityStatusReqID]",
		Short: "shows a security-status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSecurityStatusReqID := args[0]

			params := &types.QueryGetSecurityStatusRequest{
				SecurityStatusReqID: argSecurityStatusReqID,
			}

			res, err := queryClient.SecurityStatus(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
