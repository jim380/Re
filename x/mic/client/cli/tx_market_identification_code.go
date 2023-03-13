package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/mic/types"
	"github.com/spf13/cobra"
)

func CmdCreateMarketIdentificationCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-market-identification-code [mic] [name] [location] [asset-class] [currency] [regulatory-authority] [status]",
		Short: "Create a new marketIdentificationCode",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMIC := args[0]
			argName := args[1]
			argLocation := args[2]
			argAssetClass := args[3]
			argCurrency := args[4]
			argRegulatoryAuthority := args[5]
			argStatus := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMarketIdentificationCode(clientCtx.GetFromAddress().String(), argMIC, argName, argLocation, argAssetClass, argCurrency, argRegulatoryAuthority, argStatus)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateMarketIdentificationCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-market-identification-code [id] [mic] [name] [location] [asset-class] [currency] [regulatory-authority] [status]",
		Short: "Update a marketIdentificationCode",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argMIC := args[1]

			argName := args[2]

			argLocation := args[3]

			argAssetClass := args[4]

			argCurrency := args[5]

			argRegulatoryAuthority := args[6]

			argStatus := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMarketIdentificationCode(clientCtx.GetFromAddress().String(), id, argMIC, argName, argLocation, argAssetClass, argCurrency, argRegulatoryAuthority, argStatus)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteMarketIdentificationCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-market-identification-code [id]",
		Short: "Delete a marketIdentificationCode by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMarketIdentificationCode(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
