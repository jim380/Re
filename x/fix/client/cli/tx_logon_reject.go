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

func CmdLogonReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logon-reject [session-name] [did] [text]",
		Short: "Broadcast message logon-reject",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionName := args[0]

			argDID := args[1]
			header := types.Header{
				SenderCompID: argDID,
			}

			argText := args[2]

			msg := types.NewMsgLogonReject(
				clientCtx.GetFromAddress().String(),
				argSessionName,
				argText,
				header,
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
