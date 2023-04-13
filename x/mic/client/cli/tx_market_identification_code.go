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
		Use:   "register-market-identification-code [mic] [operating-mic] [OPRT-SGMT] [market-name] [legal-entity-name] [legal-entity-identifier] [market-category] [acronym] [ISO-country-code] [city] [website] [status] [creation-date] [last-update-date] [last-validation-date] [expiry-date] [comments]",
		Short: "Registers a new marketIdentificationCode",
		Args:  cobra.ExactArgs(17),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argMIC := args[0]
			argOperating_MIC := args[1]
			argOPRT_SGMT := args[2]
			argMarketName := args[3]
			argLegalEntityName := args[4]
			argLegalEntityIdentifier := args[5]
			argMarketCategory := args[6]
			argAcronym := args[7]
			argISOCountryCode := args[8]
			argCity := args[9]
			argWebsite := args[10]
			argStatus := args[11]
			argCreationDate := args[12]
			argLastUpdateDate := args[13]
			argLastValidationDate := args[14]
			argExpiryDate := args[15]
			argComments := args[16]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterMarketIdentificationCode(clientCtx.GetFromAddress().String(), argMIC, argOperating_MIC, argOPRT_SGMT, argMarketName, argLegalEntityName, argLegalEntityIdentifier, argMarketCategory, argAcronym, argISOCountryCode, argCity, argWebsite, argStatus, argCreationDate, argLastUpdateDate, argLastValidationDate, argExpiryDate, argComments)
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
		Use:   "update-market-identification-code [mic] [operating-mic] [OPRT-SGMT] [market-name] [legal-entity-name] [legal-entity-identifier] [market-category] [acronym] [ISO-country-code] [city] [website] [status] [creation-date] [last-update-date] [last-validation-date] [expiry-date] [comments]",
		Short: "Update a marketIdentificationCode",
		Args:  cobra.ExactArgs(17),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argMIC := args[0]
			argOperating_MIC := args[1]
			argOPRT_SGMT := args[2]
			argMarketName := args[3]
			argLegalEntityName := args[4]
			argLegalEntityIdentifier := args[5]
			argMarketCategory := args[6]
			argAcronym := args[7]
			argISOCountryCode := args[8]
			argCity := args[9]
			argWebsite := args[10]
			argStatus := args[11]
			argCreationDate := args[12]
			argLastUpdateDate := args[13]
			argLastValidationDate := args[14]
			argExpiryDate := args[15]
			argComments := args[16]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMarketIdentificationCode(clientCtx.GetFromAddress().String(), argMIC, argOperating_MIC, argOPRT_SGMT, argMarketName, argLegalEntityName, argLegalEntityIdentifier, argMarketCategory, argAcronym, argISOCountryCode, argCity, argWebsite, argStatus, argCreationDate, argLastUpdateDate, argLastValidationDate, argExpiryDate, argComments)
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
