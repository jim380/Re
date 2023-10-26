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
		Use:   "list-sessions [chainID]",
		Short: "list all sessions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]

			params := &types.QueryAllSessionsRequest{
				ChainID:    argChainID,
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
		Use:   "show-sessions [chainID] [sessionID]",
		Short: "shows a sessions",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argSessionID := args[1]

			params := &types.QueryGetSessionsRequest{
				ChainID:   argChainID,
				SessionID: argSessionID,
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
		Use:   "list-session-reject [chainID]",
		Short: "list all rejected sessions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSessionRejectRequest{
				ChainID:    argChainID,
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
		Use:   "show-session-reject [chainID] [sessionID]",
		Short: "shows a rejected session",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argSessionID := args[1]

			params := &types.QueryGetSessionRejectRequest{
				ChainID:   argChainID,
				SessionID: argSessionID,
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
		Use:   "list-session-logout [chainID]",
		Short: "list all session-logout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argChainID := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSessionLogoutRequest{
				ChainID:    argChainID,
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
		Use:   "show-session-logout [chainID] [sessionID]",
		Short: "shows a session-logout",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainID := args[0]
			argSessionID := args[1]

			params := &types.QueryGetSessionLogoutRequest{
				ChainID:   argChainID,
				SessionID: argSessionID,
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
