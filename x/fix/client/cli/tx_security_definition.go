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

func CmdSecurityDefinition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-definition [security-req-id] [security-response-id] [security-response-type] [symbol] [security-exchange] [issuer] [security-desc] [security-type] [currency] [contract-multiplier] [min-price-increment] [min-price-increment-amount] [unit-of-measure] [price-unit-of-measure] [settl-type] [settl-date] [maturity-month-year] [coupon-rate] [factor] [credit-rating] [security-exchange-id]",
		Short: "Broadcast message security-definition",
		Args:  cobra.ExactArgs(21),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
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
