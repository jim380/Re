package keeper_test

import (
	"time"

	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestLogonInitiator() {
	type args struct {
		msgLogonInitiator types.MsgLogonInitiator
	}

	type errArgs struct {
		shouldPass bool
		contains   string
	}

	tests := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			"Valid Logon Request",
			args{
				msgLogonInitiator: *types.NewMsgLogonInitiator(suite.address[0].String(), "sessionID1", types.LogonInitiator{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					EncryptMethod: 1,
					HeartBtInt:    1,
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				}),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// test logic
		})
	}
}
