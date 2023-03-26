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

func CmdQuoteRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote-request [session-id] [quote-request]",
		Short: "Broadcast message quote-request",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argQuoteRequest := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
            
			//TODO
			t := types.QuoteRequest{QuoteReqID: argQuoteRequest}

			msg := types.NewMsgQuoteRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				&t,
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
