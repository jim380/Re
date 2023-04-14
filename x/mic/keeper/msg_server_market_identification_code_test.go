package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/mic/types"
)

func (suite *KeeperTestSuite) TestRegisterMarketIdentificationCode() {

	type args struct {
		marketIdentificationCode types.MarketIdentificationCode
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
			name: "MIC is Found",
			args: args{
				marketIdentificationCode: types.MarketIdentificationCode{
					Creator:               string(suite.address[0]),
					MIC:                   "EXAMPLE",
					Operating_MIC:         "EXAMPLE",
					OPRT_SGMT:             "EXAMPLE",
					MarketName:            "Example Market",
					LegalEntityName:       "Example Legal Entity",
					LegalEntityIdentifier: "EXAMPLE",
					MarketCategory:        "EXAMPLE",
					Acronym:               "EXAMPLE",
					ISOCountryCode:        "EXAMPLE",
					City:                  "Example City",
					Website:               "www.example.com",
					Status:                "Active",
					CreationDate:          "2023-04-14",
					LastUpdateDate:        "2023-04-14",
					LastValidationDate:    "2023-04-14",
					ExpiryDate:            "2023-04-14",
					Comments:              "Example comments",
				},
			},
			errArgs: errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// NewMsgRegisterMarketIdentificationCode instance
			msg := types.NewMsgRegisterMarketIdentificationCode(tc.args.marketIdentificationCode.Creator, tc.args.marketIdentificationCode.MIC, tc.args.marketIdentificationCode.Operating_MIC, tc.args.marketIdentificationCode.OPRT_SGMT, tc.args.marketIdentificationCode.MarketName, tc.args.marketIdentificationCode.LegalEntityName, tc.args.marketIdentificationCode.LegalEntityIdentifier, tc.args.marketIdentificationCode.MarketCategory, tc.args.marketIdentificationCode.Acronym, tc.args.marketIdentificationCode.ISOCountryCode, tc.args.marketIdentificationCode.City, tc.args.marketIdentificationCode.Website, tc.args.marketIdentificationCode.Status, tc.args.marketIdentificationCode.CreationDate, tc.args.marketIdentificationCode.LastUpdateDate, tc.args.marketIdentificationCode.LastValidationDate, tc.args.marketIdentificationCode.ExpiryDate, tc.args.marketIdentificationCode.Comments)

			// set mic from mic module
			suite.micKeeper.SetMarketIdentificationCode(suite.ctx, tc.args.marketIdentificationCode)

			// call RegisterMarketIdentificationCode method
			res, err := suite.msgServer.RegisterMarketIdentificationCode(sdk.WrapSDKContext(suite.ctx), msg)

			// get mic from mic module
			mic, found := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.marketIdentificationCode.MIC)

			println("testing", res.String(), err.Error(), mic.String(), found)

		})
	}

}
