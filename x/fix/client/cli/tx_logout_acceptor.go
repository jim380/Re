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

func CmdLogoutAcceptor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout-acceptor [session-name] [session-name",
		Short: "Broadcast message logout-acceptor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argText := args[0]

			sessionLogoutInitiator := types.SessionLogoutAcceptor{
				Text: argText,
			}

			msg := types.NewMsgLogoutAcceptor(
				clientCtx.GetFromAddress().String(),
				argSessionName,
				sessionLogoutInitiator,
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
