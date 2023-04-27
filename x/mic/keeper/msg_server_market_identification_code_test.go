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
			name: "Valid Market Identification Code",
			args: args{
				marketIdentificationCode: types.MarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
		{
			name: "Market Identification Code is Empty",
			args: args{
				marketIdentificationCode: types.MarketIdentificationCode{
					Creator:               suite.address[0].String(),
					MIC:                   "",
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
				shouldPass: false,
				contains:   "",
			},
		},
		{
			name: "MIC creator is Empty, Invalid Address",
			args: args{
				marketIdentificationCode: types.MarketIdentificationCode{
					Creator:               "",
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
				shouldPass: false,
				contains:   "",
			},
		},
		{
			name: "When any Required field is empty",
			args: args{
				marketIdentificationCode: types.MarketIdentificationCode{
					Creator:               suite.address[0].String(),
					MIC:                   "EXAMPLE",
					Operating_MIC:         "",
					OPRT_SGMT:             "",
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
				shouldPass: false,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// Check if the MIC already exists
			_, found := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.marketIdentificationCode.MIC)
			if found {
				// Skip registration if the MIC already exists
				return
			}

			// NewMsgRegisterMarketIdentificationCode instance
			msg := types.NewMsgRegisterMarketIdentificationCode(tc.args.marketIdentificationCode.Creator, tc.args.marketIdentificationCode.MIC, tc.args.marketIdentificationCode.Operating_MIC, tc.args.marketIdentificationCode.OPRT_SGMT, tc.args.marketIdentificationCode.MarketName, tc.args.marketIdentificationCode.LegalEntityName, tc.args.marketIdentificationCode.LegalEntityIdentifier, tc.args.marketIdentificationCode.MarketCategory, tc.args.marketIdentificationCode.Acronym, tc.args.marketIdentificationCode.ISOCountryCode, tc.args.marketIdentificationCode.City, tc.args.marketIdentificationCode.Website, tc.args.marketIdentificationCode.Status, tc.args.marketIdentificationCode.CreationDate, tc.args.marketIdentificationCode.LastUpdateDate, tc.args.marketIdentificationCode.LastValidationDate, tc.args.marketIdentificationCode.ExpiryDate, tc.args.marketIdentificationCode.Comments)

			// call RegisterMarketIdentificationCode method
			_, err := suite.msgServer.RegisterMarketIdentificationCode(sdk.WrapSDKContext(suite.ctx), msg)

			// get mic
			mic, found := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.marketIdentificationCode.MIC)

			// Assert that no error occurred
			if tc.errArgs.shouldPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().True(found, tc.name)
				suite.Require().Equal(tc.args.marketIdentificationCode, mic, tc.name)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Contains(err.Error(), tc.errArgs.contains, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdateMarketIdentificationCode() {
	type args struct {
		registerMarketIdentificationCode types.MsgRegisterMarketIdentificationCode
		updateMarketIdentificationCode   types.MsgUpdateMarketIdentificationCode
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
			name: "Update Outdated Market Identification code",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				updateMarketIdentificationCode: types.MsgUpdateMarketIdentificationCode{
					Creator:               suite.address[0].String(),
					Old_MIC:               "EXAMPLE",
					New_MIC:               "EXAMPLE1",
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
		{
			name: "When the provided MIC to be updated does not exist",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				updateMarketIdentificationCode: types.MsgUpdateMarketIdentificationCode{
					Creator:               suite.address[0].String(),
					Old_MIC:               "EXAMPLEXXX",
					New_MIC:               "EXAMPLE1",
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
				shouldPass: false,
				contains:   "",
			},
		},
		{
			name: "When the creator updating MIC is empty, Invalid Address",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				updateMarketIdentificationCode: types.MsgUpdateMarketIdentificationCode{
					Creator:               "",
					Old_MIC:               "EXAMPLE",
					New_MIC:               "EXAMPLE1",
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
				shouldPass: false,
				contains:   "",
			},
		},
		{
			name: "When the account address updating MIC is not the original creator of the MIC",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				updateMarketIdentificationCode: types.MsgUpdateMarketIdentificationCode{
					Creator:               suite.address[1].String(),
					Old_MIC:               "EXAMPLE",
					New_MIC:               "EXAMPLE1",
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
				shouldPass: false,
				contains:   "",
			},
		},
		{
			name: "When any of the required field to be updated is empty",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				updateMarketIdentificationCode: types.MsgUpdateMarketIdentificationCode{
					Creator:               suite.address[0].String(),
					Old_MIC:               "EXAMPLE",
					New_MIC:               "EXAMPLE1",
					Operating_MIC:         "EXAMPLE",
					OPRT_SGMT:             "",
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
				shouldPass: false,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// Check if the MIC already exists
			_, found := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.registerMarketIdentificationCode.MIC)
			if found {
				// Skip registration if the MIC already exists
				return
			}

			// call RegisterMarketIdentificationCode method
			_, _ = suite.msgServer.RegisterMarketIdentificationCode(sdk.WrapSDKContext(suite.ctx), &tc.args.registerMarketIdentificationCode)

			// get old mic
			_, _ = suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.updateMarketIdentificationCode.Old_MIC)

			// now that mic is registered using RegisterMarketIdentificationCode, call UpdateMarketIdentificationCode method
			_, err := suite.msgServer.UpdateMarketIdentificationCode(sdk.WrapSDKContext(suite.ctx), &tc.args.updateMarketIdentificationCode)

			// get new mic after updating
			newMIC, found := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.updateMarketIdentificationCode.New_MIC)

			// Assert that no error occurred
			if tc.errArgs.shouldPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().True(found, tc.name)
				suite.Require().Equal(tc.args.updateMarketIdentificationCode.New_MIC, newMIC.MIC, tc.name)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Contains(err.Error(), tc.errArgs.contains, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteMarketIdentificationCode() {
	type args struct {
		registerMarketIdentificationCode types.MsgRegisterMarketIdentificationCode
		deleteMarketIdentificationCode   types.MsgDeleteMarketIdentificationCode
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
			name: "When the MIC to be deleted exists and the creator is the one deleting the MIC",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				deleteMarketIdentificationCode: types.MsgDeleteMarketIdentificationCode{
					Creator: suite.address[0].String(),
					MIC:     "EXAMPLE",
				},
			},
			errArgs: errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			name: "When the MIC to be deleted does not exist",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				deleteMarketIdentificationCode: types.MsgDeleteMarketIdentificationCode{
					Creator: suite.address[0].String(),
					MIC:     "EXAMPLEEEE",
				},
			},
			errArgs: errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			name: "When a different address deletes the MIC",
			args: args{
				registerMarketIdentificationCode: types.MsgRegisterMarketIdentificationCode{
					Creator:               suite.address[0].String(),
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
				deleteMarketIdentificationCode: types.MsgDeleteMarketIdentificationCode{
					Creator: suite.address[1].String(),
					MIC:     "EXAMPLE",
				},
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

			// Check if the MIC already exists
			_, found := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.registerMarketIdentificationCode.MIC)
			if found {
				// Skip registration if the MIC already exists
				return
			}

			// call RegisterMarketIdentificationCode method
			_, _ = suite.msgServer.RegisterMarketIdentificationCode(sdk.WrapSDKContext(suite.ctx), &tc.args.registerMarketIdentificationCode)

			// get mic
			_, _ = suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.registerMarketIdentificationCode.MIC)

			// now that mic is registered using RegisterMarketIdentificationCode, call DeleteMarketIdentificationCode method
			_, err := suite.msgServer.DeleteMarketIdentificationCode(sdk.WrapSDKContext(suite.ctx), &tc.args.deleteMarketIdentificationCode)

			// remove MIC
			suite.micKeeper.RemoveMarketIdentificationCode(suite.ctx, tc.args.deleteMarketIdentificationCode.MIC)

			// mic does not exist after deleting
			deletedMIC, found := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.deleteMarketIdentificationCode.MIC)

			// Assert that no error occurred
			if tc.errArgs.shouldPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().Empty(deletedMIC.String())
				suite.Require().False(found, tc.name)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Empty(deletedMIC.String())
				suite.Require().Contains(err.Error(), tc.errArgs.contains, tc.name)
			}
		})
	}
}
