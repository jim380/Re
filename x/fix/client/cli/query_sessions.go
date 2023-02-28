package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

func CmdListSessions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-sessions",
		Short: "list all sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSessionsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SessionsAll(context.Background(), params)
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

func CmdShowSessions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-sessions [session-name]",
		Short: "shows a sessions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionName := args[0]

			params := &types.QueryGetSessionsRequest{
				SessionName: argSessionName,
			}

			res, err := queryClient.Sessions(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListSessionReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-session-reject",
		Short: "list all rejected sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSessionRejectRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SessionRejectAll(context.Background(), params)
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

func CmdShowSessionReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-session-reject [session-name]",
		Short: "shows a rejected session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionName := args[0]

			params := &types.QueryGetSessionRejectRequest{
				SessionName: argSessionName,
			}

			res, err := queryClient.SessionReject(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListSessionLogout() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-session-logout",
		Short: "list all session-logout",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSessionLogoutRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SessionLogoutAll(context.Background(), params)
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

func CmdShowSessionLogout() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-session-logout [id]",
		Short: "shows a session-logout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSessionName := args[0]

			params := &types.QueryGetSessionLogoutRequest{
				SessionName: argSessionName,
			}

			res, err := queryClient.SessionLogout(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
