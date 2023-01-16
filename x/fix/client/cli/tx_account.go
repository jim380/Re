package cli

import (
	"strconv"

	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-account [company-name] [address] [email-address] [phone-number] [website] [social-media-links] [government-issued-id]",
		Short: "Create a new account",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCompanyName := args[0]
			argAddress := args[1]
			argEmailAddress := args[2]
			argPhoneNumber, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argWebsite := args[4]
			argSocialMediaLinks := strings.Split(args[5], listSeparator)
			argGovernmentIssuedId := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAccount(clientCtx.GetFromAddress().String(), argCompanyName, argAddress, argEmailAddress, argPhoneNumber, argWebsite, argSocialMediaLinks, argGovernmentIssuedId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-account [id] [company-name] [address] [email-address] [phone-number] [website] [social-media-links] [government-issued-id]",
		Short: "Update a account",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argCompanyName := args[1]

			argAddress := args[2]

			argEmailAddress := args[3]

			argPhoneNumber, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			argWebsite := args[5]

			argSocialMediaLinks := strings.Split(args[6], listSeparator)

			argGovernmentIssuedId := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAccount(clientCtx.GetFromAddress().String(), id, argCompanyName, argAddress, argEmailAddress, argPhoneNumber, argWebsite, argSocialMediaLinks, argGovernmentIssuedId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-account [id]",
		Short: "Delete a account by id",
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

			msg := types.NewMsgDeleteAccount(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
