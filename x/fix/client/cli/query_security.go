package cli

import (
    "context"
    "strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
    "github.com/jim380/Re/x/fix/types"
)

func CmdListSecurity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-security",
		Short: "list all security",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllSecurityRequest{
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
		Use:   "show-security [id]",
		Short: "shows a security",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

            id, err := strconv.ParseUint(args[0], 10, 64)
            if err != nil {
                return err
            }

            params := &types.QueryGetSecurityRequest{
                Id: id,
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
