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
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// add sessionID

			argSecurityReqID := args[0]
			argSecurityRequestType := args[1]
			argSymbol := args[2]
			argSecurityExchange := args[3]
			argIssuer := args[4]
			argSecurityDesc := args[5]
			argSecurityType := args[6]
			argCurrency := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityDefinitionRequest(
				clientCtx.GetFromAddress().String(),
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
		Args:  cobra.ExactArgs(21),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// sessionID
			
			argSecurityReqID := args[0]
			argSecurityResponseID := args[1]
			argSecurityResponseType := args[2]
			argSymbol := args[3]
			argSecurityExchange := args[4]
			argIssuer := args[5]
			argSecurityDesc := args[6]
			argSecurityType := args[7]
			argCurrency := args[8]
			argContractMultiplier := args[9]
			argMinPriceIncrement := args[10]
			argMinPriceIncrementAmount := args[11]
			argUnitOfMeasure := args[12]
			argPriceUnitOfMeasure := args[13]
			argSettlType := args[14]
			argSettlDate := args[15]
			argMaturityMonthYear := args[16]
			argCouponRate := args[17]
			argFactor := args[18]
			argCreditRating := args[19]
			argSecurityExchangeID := args[20]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityDefinition(
				clientCtx.GetFromAddress().String(),
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
