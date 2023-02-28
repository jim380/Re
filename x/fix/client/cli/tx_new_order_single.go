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

func CmdNewOrderSingle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-order-single [session-name] [ClOrdID] [Symbol] [Side] [OrderQty] [OrdType] [Price] [TimeInForce] [Text]",
		Short: "Broadcast message new-order-single",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argSessionName := args[0]

			argClOrdID := args[1]

			argSymbol := args[2]

			argSide, err := cast.ToInt64E(args[3])
			if err != nil {
				return err
			}

			argOrderQty := args[4]

			argOrdType, err := cast.ToInt64E(args[5])
			if err != nil {
				return err
			}

			argPrice := args[6]

			argTimeInForce, err := cast.ToInt64E(args[7])
			if err != nil {
				return err
			}

			argText := args[8]

			msg := types.NewMsgNewOrderSingle(
				clientCtx.GetFromAddress().String(),
				argSessionName,
				argClOrdID,
				argSymbol,
				argSide,
				argOrderQty,
				argOrdType,
				argPrice,
				argTimeInForce,
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
