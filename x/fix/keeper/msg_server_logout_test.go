package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestLogoutInitiator() {
	type args struct {
		session            types.Sessions
		msgLogoutInitiator types.MsgLogoutInitiator
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
			"Valid Logout Request",
			args{
				session:            types.Sessions{},
				msgLogoutInitiator: *types.NewMsgLogoutInitiator(suite.address[0].String(), "sessionID", types.SessionLogoutInitiator{Text: "Request to Logout from the current session"}),
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

			// call LogoutInitiator method
			res, err := suite.msgServer.LogoutInitiator(sdk.WrapSDKContext(suite.ctx), &tc.args.msgLogoutInitiator)
			session, found := suite.fixKeeper.GetSessions(suite.ctx, tc.args.msgLogoutInitiator.SessionID)

		})
	}
}
