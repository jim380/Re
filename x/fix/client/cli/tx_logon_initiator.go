package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jim380/Re/x/fix/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdLogonInitiator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logon-initiator [msg-type] [sender-compID] [target-compID] [encrypt-method] [heart-beat-int]",
		Short: "Broadcast message logon-initiator",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argMsgType := args[0]

			argSenderCompID := args[1]

			argTargetCompID := args[2]

			argEncryptMethod, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argHeartBtInt, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			header := types.NewHeader(argMsgType, argSenderCompID, argTargetCompID)

			argLogonIntiator := types.NewLogonInitiator(header, int64(argEncryptMethod), int64(argHeartBtInt))

			msg := types.NewMsgLogonInitiator(
				clientCtx.GetFromAddress().String(),
				argLogonIntiator,
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
