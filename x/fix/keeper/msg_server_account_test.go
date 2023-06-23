package keeper_test

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestRegisterAccount() {
	type args struct {
		creator          string
		address          string
		companyName      string
		website          string
		socialMediaLinks string
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
			name: "Valid Address",
			args: args{
				creator:          suite.address[0].String(),
				address:          suite.address[0].String(),
				companyName:      "CypherCore",
				website:          "CypherCore.io",
				socialMediaLinks: "@CypherCore",
			},
			errArgs: errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			name: "InValid Address",
			args: args{
				creator:          suite.address[1].String(),
				address:          "",
				companyName:      "CypherCore",
				website:          "CypherCore.io",
				socialMediaLinks: "@CypherCore",
			},
			errArgs: errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			name: "When the Address is not the user's Account Address",
			args: args{
				creator:          suite.address[0].String(),
				address:          suite.address[1].String(),
				companyName:      "CypherCore",
				website:          "CypherCore.io",
				socialMediaLinks: "@CypherCore",
			},
			errArgs: errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			msg := types.MsgRegisterAccount{
				Creator:          tc.args.creator,
				Address:          tc.args.address,
				CompanyName:      tc.args.companyName,
				Website:          tc.args.website,
				SocialMediaLinks: tc.args.socialMediaLinks,
			}

			// call RegisterAccount method
			res, err := suite.msgServer.RegisterAccount(sdk.WrapSDKContext(suite.ctx), &msg)

			// GetAccountRegistration
			getAcc, found := suite.fixKeeper.GetAccountRegistration(suite.ctx, tc.args.address)

			if tc.errArgs.shouldPass {
				suite.Require().True(found)
				suite.Require().NoError(err, tc.name)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(getAcc)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Nil(res)
				suite.Require().Empty(getAcc)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}

		})
	}

}
