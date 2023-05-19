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
		Use:   "security-definition-request [security-req-id] [security-request-type] [symbol] [security-exchange] [issuer] [security-desc] [security-type] [currency]",
		Short: "Broadcast message security-definition-request",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
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
