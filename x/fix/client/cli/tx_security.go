package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSecurityDefinitionRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-definition-request [sesssion-id] [security-req-id] [security-request-type] [symbol] [security-exchange] [issuer] [security-desc] [security-type] [currency]",
		Short: "Broadcast message security-definition-request",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityRequestType := args[2]
			argSymbol := args[3]
			argSecurityExchange := args[4]
			argIssuer := args[5]
			argSecurityDesc := args[6]
			argSecurityType := args[7]
			argCurrency := args[8]

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
		Use:   "security-definition [session-id] [security-req-id] [security-response-id] [security-response-type] [symbol] [security-exchange] [issuer] [security-desc] [security-type] [currency] [contract-multiplier] [min-price-increment] [min-price-increment-amount] [unit-of-measure] [price-unit-of-measure] [settl-type] [settl-date] [maturity-month-year] [coupon-rate] [factor] [credit-rating] [security-exchange-id]",
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
		Use:   "security-definition-request-reject [session-id] [security-req-id] [security-request-result] [security-request-error] [security-request-error-code] [text]",
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
