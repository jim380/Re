package cli

import (
    "strconv"
	
	 "encoding/json"
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/fix/types"
)

var _ = strconv.Itoa(0)

func CmdQuoteRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote-request [session-id] [quote-request]",
		Short: "Broadcast message quote-request",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argSessionID := args[0]
             argQuoteRequest := new(types.QuoteRequest)
					err = json.Unmarshal([]byte(args[1]), argQuoteRequest)
    				if err != nil {
                		return err
            		}
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgQuoteRequest(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argQuoteRequest,
				
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