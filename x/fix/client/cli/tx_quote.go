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

// CmdQuoteRequest is the command line tool for creating Quote Request
func CmdQuoteRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote-request [sessionID] [symbol] [securityID] [securityIDSource] [side] [orderQty] [futSettDate] [settlDate2] [account] [bidPx] [offerPx] [currency] [validUntilTime] [expireTime] [quoteType] [bidSize] [offerSize] [mic] [text]",
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

			// GenerateRandomString function uniquely generates quoteReqID for every Quote Request
			quoteReqID, _ := helpers.GenerateRandomString(constants.QuoteReqIDLength)

			quoteRequests := types.NewQuoteRequest(quoteReqID, argSymbol, argSecurityID, argSecurityIDSource, argSide, argOrderQty, argFutSettDate, argSettlDate2, argAccount, argBidPx, argOfferPx, argCurrency, argValidUntilTime, argExpireTime, argQuoteType, argBidSize, argOfferSize, argMIC, argText)

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

// CmdQuoteAcknowledgement is the command line tool for creating Quote Acknowledgement
func CmdQuoteAcknowledgement() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote-acknowledgement [sessionID] [quoteReqID] [quoteStatus] [quoteType] [securityID] [securityIDSource] [symbol] [side] [orderQty] [lastQty] [lastPx] [bidPx] [offerPx] [currency] [settlDate] [validUntilTime] [expireTime] [text] [noQuoteQualifiers] [quoteQualifier] [noLegs] [legSymbol] [legSecurityID] [legSecurityIDSource] [legRatioQty]",
		Short: "Broadcast message quote-acknowledgement",
		Args:  cobra.ExactArgs(25),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]
			argQuoteReqID := args[1]
			argQuoteStatus := args[2]
			argQuoteType := args[3]
			argSecurityID := args[4]
			argSecurityIDSource := args[5]
			argSymbol := args[6]
			argSide := args[7]
			argOrderQty := args[8]
			argLastQty := args[9]
			argLastPx := args[10]
			argBidPx := args[11]
			argOfferPx := args[12]
			argCurrency := args[13]
			argSettlDate := args[14]
			argValidUntilTime := args[15]
			argExpireTime := args[16]
			argText := args[17]
			argNoQuoteQualifiers := args[18]
			argQuoteQualifier := args[19]
			argNoLegs := args[20]
			argLegSymbol := args[21]
			argLegSecurityID := args[22]
			argLegSecurityIDSource := args[23]
			argLegRatioQty := args[24]

			// auto generate QuoteID using GenerateRandomString function
			quoteID, _ := helpers.GenerateRandomString(constants.QuoteReqIDLength)

			quoteAcknowledgement := types.NewQuoteAcknowledgement(argQuoteReqID, quoteID, argQuoteStatus, argQuoteType, argSecurityID, argSecurityIDSource, argSymbol, argSide, argOrderQty, argLastQty, argLastPx, argBidPx, argOfferPx, argCurrency, argSettlDate, argValidUntilTime, argExpireTime, argText, argNoQuoteQualifiers, argQuoteQualifier, argNoLegs, argLegSymbol, argLegSecurityID, argLegSecurityIDSource, argLegRatioQty)

			msg := types.NewMsgQuoteAcknowledgement(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				quoteAcknowledgement,
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

// CmdQuoteRequestReject is the command line for creating Quote Request Reject
func CmdQuoteRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote-request-reject [sessionID] [quoteReqID] [quoteRequestRejectReason] [text]",
		Short: "Broadcast message quote-request-reject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]
			argQuoteReqID := args[1]
			argQuoteRequestRejectReason := args[2]
			argText := args[3]

			quoteRequestReject := types.NewQuoteRequestReject(argQuoteReqID, argQuoteRequestRejectReason, argText)

			msg := types.NewMsgQuoteRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				quoteRequestReject,
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
