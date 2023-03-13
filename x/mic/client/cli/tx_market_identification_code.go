package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/mic/types"
	"github.com/spf13/cobra"
)

func CmdRegisterMarketIdentificationCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-market-identification-code [mic] [name] [location] [asset-class] [currency] [regulatory-authority] [status]",
		Short: "Registers a new marketIdentificationCode",
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

			msg := types.NewMsgRegisterMarketIdentificationCode(clientCtx.GetFromAddress().String(), argMIC, argName, argLocation, argAssetClass, argCurrency, argRegulatoryAuthority, argStatus)
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
		Use:   "update-market-identification-code [mic] [name] [location] [asset-class] [currency] [regulatory-authority] [status]",
		Short: "Update a marketIdentificationCode",
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

			msg := types.NewMsgUpdateMarketIdentificationCode(clientCtx.GetFromAddress().String(), argMIC, argName, argLocation, argAssetClass, argCurrency, argRegulatoryAuthority, argStatus)
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
		Use:   "delete-market-identification-code [mic]",
		Short: "Delete a marketIdentificationCode by mic",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argMIC := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMarketIdentificationCode(clientCtx.GetFromAddress().String(), argMIC)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
