package keeper_test

import (
	"log"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestLogonInitiator() {
	type args struct {
		msgRegisterAccount1 types.MsgRegisterAccount
		msgLogonInitiator   types.MsgLogonInitiator
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
		/*
			{
				"Same Address can not used for SenderCompID & TargetCompID",
				args{
					msgLogonInitiator: *types.NewMsgLogonInitiator(suite.address[0].String(), "sessionID2", types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[0].String(),
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
					shouldPass: false,
					contains:   "",
				},
			},
			{
				"InitiatorAddress is different from SenderCompID",
				args{
					msgLogonInitiator: *types.NewMsgLogonInitiator(suite.address[0].String(), "sessionID3", types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[1].String(),
							TargetCompID: suite.address[2].String(),
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
					shouldPass: false,
					contains:   "",
				},
			},
		*/
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// Register both address
			msgRegisterAccount1 := types.MsgRegisterAccount{
				Creator:          tc.args.msgLogonInitiator.LogonInitiator.Header.SenderCompID,
				Address:          tc.args.msgLogonInitiator.LogonInitiator.Header.SenderCompID,
				CompanyName:      "",
				Website:          "",
				SocialMediaLinks: "",
			}

			msgRegisterAccount2 := types.MsgRegisterAccount{
				Creator:          tc.args.msgLogonInitiator.LogonInitiator.Header.TargetCompID,
				Address:          tc.args.msgLogonInitiator.LogonInitiator.Header.TargetCompID,
				CompanyName:      "cypherCore",
				Website:          "cypherCore.io",
				SocialMediaLinks: "@cypherCore",
			}

			// call RegisterAccount method
			suite.msgServer.RegisterAccount(sdk.WrapSDKContext(suite.ctx), &msgRegisterAccount1)
			suite.msgServer.RegisterAccount(sdk.WrapSDKContext(suite.ctx), &msgRegisterAccount2)

			getAcc, found := suite.fixKeeper.GetAccountRegistration(suite.ctx, tc.args.msgLogonInitiator.LogonInitiator.Header.SenderCompID)

			getAcc1, found1 := suite.fixKeeper.GetAccountRegistration(suite.ctx, tc.args.msgLogonInitiator.LogonInitiator.Header.TargetCompID)

			// call LogonInitiator method
			//res, err := suite.msgServer.LogonInitiator(sdk.WrapSDKContext(suite.ctx), &tc.args.msgLogonInitiator)

			//sessionID, found := suite.fixKeeper.GetSessions(suite.ctx, tc.args.msgLogonInitiator.SessionID)

			log.Println("first", getAcc, found)
			log.Println("second", getAcc1, found1)
			//log.Println("third", r, error)

		})
	}
}
