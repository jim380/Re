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

func CmdLogoutInitiator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout-initiator [session-name] [text]",
		Short: "Broadcast message logout-initiator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argText := args[1]

			sessionLogoutInitiator := types.SessionLogoutInitiator{
				Text: argText,
			}

			msg := types.NewMsgLogoutInitiator(
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
