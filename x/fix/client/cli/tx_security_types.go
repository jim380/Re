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

func CmdSecurityTypesRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-types-request [sessionID] [text] [tradingSessionID] [tradingSessionSubID] [product] [securityType] [securitySubType]",
		Short: "Broadcast message security-types-request",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argText := args[1]
			argTradingSessionID := args[2]
			argTradingSessionSubID := args[3]
			argProduct := args[4]
			argSecurityType := args[5]
			argSecuritySubType := args[6]

			// GenerateRandomString function uniquely generates SecurityReqID for every Security Types Request
			argSecurityReqID, _ := types.GenerateRandomString(types.SecurityReqID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityTypesRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argText,
				argTradingSessionID,
				argTradingSessionSubID,
				argProduct,
				argSecurityType,
				argSecuritySubType,
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

func CmdSecurityTypesResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-types-response [sessionID] [securityReqID] [securityResponseType] [totNoSecurityTypes] [lastFragment] [noSecurityTypes] [securityType] [securitySubType] [product] [cfiCode] [text] [tradingSessionID] [tradingSessionSubID] [subscriptionRequestType]",
		Short: "Broadcast message security-types-response",
		Args:  cobra.ExactArgs(14),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argSecurityResponseType := args[2]
			argTotNoSecurityTypes := args[3]
			argLastFragment := args[4]
			argNoSecurityTypes := args[5]
			argSecurityType := args[6]
			argSecuritySubType := args[7]
			argProduct := args[8]
			argCFICode := args[9]
			argText := args[10]
			argTradingSessionID := args[11]
			argTradingSessionSubID := args[12]
			argSubscriptionRequestType := args[13]

			// GenerateRandomString function uniquely generates SecurityResponseID for every Security Types Request
			argSecurityResponseID, _ := types.GenerateRandomString(types.SecurityResponseID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityTypesResponse(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argSecurityResponseID,
				argSecurityResponseType,
				argTotNoSecurityTypes,
				argLastFragment,
				argNoSecurityTypes,
				argSecurityType,
				argSecuritySubType,
				argProduct,
				argCFICode,
				argText,
				argTradingSessionID,
				argTradingSessionSubID,
				argSubscriptionRequestType,
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

func CmdSecurityTypesRequestReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-types-request-reject [sessionID] [securityReqID] [rejectReason] [text]",
		Short: "Broadcast message security-types-request-reject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argSecurityReqID := args[1]
			argRejectReason := args[2]
			argText := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSecurityTypesRequestReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argSecurityReqID,
				argRejectReason,
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
