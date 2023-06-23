package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/utils/helpers"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSecurityDefinitionRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-definition-request [sesssionID] [securityRequestType] [symbol] [securityExchange] [issuer] [securityDesc] [securityType] [currency]",
		Short: "Broadcast message security-definition-request",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityRequestType := args[1]
			argSymbol := args[2]
			argSecurityExchange := args[3]
			argIssuer := args[4]
			argSecurityDesc := args[5]
			argSecurityType := args[6]
			argCurrency := args[7]

			// GenerateRandomString function uniquely generates SecurityReqID for every Security Definition Request
			argSecurityReqID, _ := helpers.GenerateRandomString(constants.SecurityReqID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityDefinitionRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityRequestType,
				argSymbol,
				argSecurityExchange,
				argIssuer,
				argSecurityDesc,
				argSecurityType,
				argCurrency,
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

func CmdSecurityDefinition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-definition [sessionID] [securityReqID] [securityResponseID] [securityResponseType] [symbol] [securityExchange] [issuer] [securityDesc] [securityType] [currency] [contractMultiplier] [minPriceIncrement] [minPriceIncrementAmount] [unitOfMeasure] [priceUnitOfMeasure] [settlType] [settlDate] [maturityMonthYear] [couponRate] [factor] [creditRating] [securityExchangeID]",
		Short: "Broadcast message security-definition",
		Args:  cobra.ExactArgs(22),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityResponseID := args[2]
			argSecurityResponseType := args[3]
			argSymbol := args[4]
			argSecurityExchange := args[5]
			argIssuer := args[6]
			argSecurityDesc := args[7]
			argSecurityType := args[8]
			argCurrency := args[9]
			argContractMultiplier := args[10]
			argMinPriceIncrement := args[11]
			argMinPriceIncrementAmount := args[12]
			argUnitOfMeasure := args[13]
			argPriceUnitOfMeasure := args[14]
			argSettlType := args[15]
			argSettlDate := args[16]
			argMaturityMonthYear := args[17]
			argCouponRate := args[18]
			argFactor := args[19]
			argCreditRating := args[20]
			argSecurityExchangeID := args[21]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityDefinition(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityResponseID,
				argSecurityResponseType,
				argSymbol,
				argSecurityExchange,
				argIssuer,
				argSecurityDesc,
				argSecurityType,
				argCurrency,
				argContractMultiplier,
				argMinPriceIncrement,
				argMinPriceIncrementAmount,
				argUnitOfMeasure,
				argPriceUnitOfMeasure,
				argSettlType,
				argSettlDate,
				argMaturityMonthYear,
				argCouponRate,
				argFactor,
				argCreditRating,
				argSecurityExchangeID,
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

func CmdSecurityDefinitionRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-definition-request-reject [sessionID] [securityReqID] [securityRequestResult] [securityRequestError] [securityRequestErrorCode] [text]",
		Short: "Broadcast message security-definition-request-reject",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityRequestResult := args[2]
			argSecurityRequestError := args[3]
			argSecurityRequestErrorCode := args[4]
			argText := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityDefinitionRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityRequestResult,
				argSecurityRequestError,
				argSecurityRequestErrorCode,
				argText,
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
