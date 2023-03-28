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

// CmdQuoteRequest is the command line tool for creating Quote Request
func CmdQuoteRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote-request [session-id] [symbol] [securityID] [securityIDSource] [side] [orderQty] [futSettDate] [settlDate2] [account] [bidPx] [offerPx] [currency] [validUntilTime] [expireTime] [quoteType] [bidSize] [offerSize] [mic] text",
		Short: "Broadcast message quote-request",
		Args:  cobra.ExactArgs(19),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]
			argSymbol := args[1]
			argSecurityID := args[2]
			argSecurityIDSource := args[3]
			argSide := args[4]
			argOrderQty := args[5]
			argFutSettDate := args[6]
			argSettlDate2 := args[7]
			argAccount := args[8]
			argBidPx := args[9]
			argOfferPx := args[10]
			argCurrency := args[11]
			argValidUntilTime := args[12]
			argExpireTime := args[13]
			argQuoteType := args[14]
			argBidSize := args[15]
			argOfferSize := args[16]
			argMIC := args[17]
			argText := args[18]

			//GenerateRandomString function uniquely generates quoteReqID for every Quote Request
			quoteReqID, _ := types.GenerateRandomString(types.QuoteReqIDLength)

			quoteRequests := types.NewQuoteRequest(quoteReqID, argSymbol, argSecurityID, argSecurityIDSource, argSide, argOrderQty, argFutSettDate, argSettlDate2, argAccount, argBidPx, argOfferPx, argCurrency, argValidUntilTime, argExpireTime, argQuoteType, argBidSize, argOfferSize, argMIC, argText, clientCtx.GetFromAddress().String())

			msg := types.NewMsgQuoteRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				quoteRequests,
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

// CmdQuoteAcknowledgement is the command line tool for Quote Acknowledgement
func CmdQuoteAcknowledgement() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote-acknowledgement [session-id]",
		Short: "Broadcast message quote-acknowledgement",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			ack := types.QuoteAcknowledgement{}

			msg := types.NewMsgQuoteAcknowledgement(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				&ack,
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
