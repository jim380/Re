package cli

import (
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
		Use:   "create-account [index] [company-name] [address] [email-address] [phone-number] [website] [social-media-links] [government-issued-id] [created-at]",
		Short: "Create a new account",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argCompanyName := args[1]
			argAddress := args[2]
			argEmailAddress := args[3]
			argPhoneNumber, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}
			argWebsite := args[5]
			argSocialMediaLinks := strings.Split(args[6], listSeparator)

			//TO DO
			// government issued id to be uploaded as image and converted to byte
			argGovernmentIssuedId := args[7]
			governmentIssuedIdBytes := []byte(argGovernmentIssuedId)

			argCreatedAt, err := cast.ToInt32E(args[8])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAccount(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argCompanyName,
				argAddress,
				argEmailAddress,
				argPhoneNumber,
				argWebsite,
				argSocialMediaLinks,
				governmentIssuedIdBytes,
				argCreatedAt,
			)
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
		Use:   "update-account [index] [company-name] [address] [email-address] [phone-number] [website] [social-media-links] [government-issued-id] [created-at]",
		Short: "Update a account",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argCompanyName := args[1]
			argAddress := args[2]
			argEmailAddress := args[3]
			argPhoneNumber, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}
			argWebsite := args[5]
			argSocialMediaLinks := strings.Split(args[6], listSeparator)

			//TO DO
			// government issued id to be uploaded as image and converted to byte
			argGovernmentIssuedId := args[7]
			governmentIssuedIdBytes := []byte(argGovernmentIssuedId)

			argCreatedAt, err := cast.ToInt32E(args[8])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAccount(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argCompanyName,
				argAddress,
				argEmailAddress,
				argPhoneNumber,
				argWebsite,
				argSocialMediaLinks,
				governmentIssuedIdBytes,
				argCreatedAt,
			)
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
		Use:   "delete-account [index]",
		Short: "Delete a account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAccount(
				clientCtx.GetFromAddress().String(),
				indexIndex,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
