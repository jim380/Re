package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/mic/types"
	"github.com/spf13/cobra"
)

func CmdListMarketIdentificationCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-market-identification-code",
		Short: "list all marketIdentificationCode",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMarketIdentificationCodeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MarketIdentificationCodeAll(context.Background(), params)
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

func CmdShowMarketIdentificationCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-market-identification-code [MIC]",
		Short: "shows a marketIdentificationCode",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argMIC := args[0]

			params := &types.QueryGetMarketIdentificationCodeRequest{
				MIC: argMIC,
			}

			res, err := queryClient.MarketIdentificationCode(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
