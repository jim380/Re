package cli

import (
	"strconv"
	"time"

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

			sessionName, _ := types.GenerateRandomString(types.SessionNameLength)

			sendingTime := time.Now().UTC().Format("20060102-15:04:05.000")

			// sample of the msg
			// "8=FIX.4.4|9=59|35=A|34=1|49=SenderCompID|52=20230219-10:30:00.000|56=TargetCompID|98=0|108=30|141=Y|10=118|"
			// body length of logon message
			msgBody := sessionName + clientCtx.GetFromAddress().String() + strconv.FormatInt(int64(argEncryptMethod), 10) + strconv.FormatInt(int64(argHeartBtInt), 10)

			bodyLength := types.BodyLength(msgBody)

			msgSeqNum := int64(0) + 1

			header := types.NewHeader(bodyLength, argMsgType, argSenderCompID, argTargetCompID, msgSeqNum, sendingTime)

			// get the length of checksum excluding the checksum field
			checkSum := sessionName + clientCtx.GetFromAddress().String() + header.String() + strconv.FormatInt(int64(argEncryptMethod), 10) + strconv.FormatInt(int64(argHeartBtInt), 10)
			setCheckSum := types.CalculateChecksum(checkSum)

			trailer := types.NewTrailer(setCheckSum)

			argLogonIntiator := types.NewLogonInitiator(header, int64(argEncryptMethod), int64(argHeartBtInt), trailer)

			msg := types.NewMsgLogonInitiator(
				clientCtx.GetFromAddress().String(),
				sessionName,
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

func CmdLogonAcceptor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logon-acceptor [sessionID] [msg-type] [sender-compID] [target-compID] [encrypt-method] [heart-beat-int]",
		Short: "Broadcast message logon-acceptor",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			argMsgType := args[1]

			argSenderCompID := args[2]

			argTargetCompID := args[3]

			argEncryptMethod, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			argHeartBtInt, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			sendingTime := time.Now().UTC().Format("20060102-15:04:05.000")

			// sample of the msg
			// "8=FIX.4.4|9=59|35=A|34=1|49=SenderCompID|52=20230219-10:30:00.000|56=TargetCompID|98=0|108=30|141=Y|10=118|"
			// body length of logon message
			msgBody := argSessionID + clientCtx.GetFromAddress().String() + strconv.FormatInt(int64(argEncryptMethod), 10) + strconv.FormatInt(int64(argHeartBtInt), 10)

			bodyLength := types.BodyLength(msgBody)

			msgSeqNum := int64(0) + 1

			header := types.NewHeader(bodyLength, argMsgType, argSenderCompID, argTargetCompID, msgSeqNum, sendingTime)

			// get the length of checksum excluding the checksum field
			checkSum := argSessionID + clientCtx.GetFromAddress().String() + header.String() + strconv.FormatInt(int64(argEncryptMethod), 10) + strconv.FormatInt(int64(argHeartBtInt), 10)
			setCheckSum := types.CalculateChecksum(checkSum)

			trailer := types.NewTrailer(setCheckSum)

			argLogonAcceptor := types.NewLogonAcceptor(header, int64(argEncryptMethod), int64(argHeartBtInt), trailer)

			msg := types.NewMsgLogonAcceptor(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argLogonAcceptor,
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

func CmdLogonReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logon-reject [sessionID] [did] [text]",
		Short: "Broadcast message logon-reject",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argSessionID := args[0]

			argDID := args[1]
			header := types.Header{
				SenderCompID: argDID,
			}

			argText := args[2]

			msg := types.NewMsgLogonReject(
				clientCtx.GetFromAddress().String(),
				argSessionID,
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

func CmdTerminateLogon() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "terminate-logon [sessionID] [did]",
		Short: "Broadcast message terminate-logon",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionID := args[0]
			argDid := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTerminateLogon(
				clientCtx.GetFromAddress().String(),
				argSessionID,
				argDid,
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
