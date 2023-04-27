package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
	micTypes "github.com/jim380/Re/x/mic/types"
)

func (suite *KeeperTestSuite) TestQuoteRequest() {
	type args struct {
		creator                  string
		session                  types.Sessions
		marketIdentificationCode micTypes.MarketIdentificationCode
		quoteRequest             types.QuoteRequest
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
			"Valid SessionID in QuoteRequest",
			args{
				creator: suite.address[0].String(),
				session: types.Sessions{
					SessionID: "xhyhjjjk",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					InitiatorAddress: string(suite.address[0].String()),
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					AcceptorAddress: string(suite.address[1].String()),
					Status:          "loggedIn",
					IsAccepted:      true,
				},
				marketIdentificationCode: micTypes.MarketIdentificationCode{
					MIC:                   "XNYS",
					Operating_MIC:         "",
					OPRT_SGMT:             "",
					MarketName:            "",
					LegalEntityName:       "",
					LegalEntityIdentifier: "",
					MarketCategory:        "",
					Acronym:               "",
					ISOCountryCode:        "",
					City:                  "",
					Website:               "",
					Status:                "",
					CreationDate:          "",
					LastUpdateDate:        "",
					LastValidationDate:    "",
					ExpiryDate:            "",
					Comments:              "",
					Creator:               suite.address[0].String(),
				},
				quoteRequest: types.QuoteRequest{
					QuoteReqID:       "aaaa",
					Symbol:           "AAPL",
					SecurityID:       "123456",
					SecurityIDSource: "CUSIP",
					Side:             "BUY",
					OrderQty:         "100",
					FutSettDate:      "20230410",
					SettlDate2:       "20230412",
					Account:          "ACCT123",
					BidPx:            "145.50",
					OfferPx:          "145.60",
					Currency:         "USD",
					ValidUntilTime:   time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
					ExpireTime:       time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
					QuoteType:        "IND",
					BidSize:          "50",
					OfferSize:        "75",
					Mic:              "XNYS",
					Text:             "Example quote request",
					Creator:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Check that both account addresses involved in a logon session can send a Quote Request",
			args{
				creator: suite.address[1].String(),
				session: types.Sessions{
					SessionID: "xhyhjjjk",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					InitiatorAddress: string(suite.address[1].String()),
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					AcceptorAddress: string(suite.address[0].String()),
					Status:          "loggedIn",
					IsAccepted:      true,
				},
				marketIdentificationCode: micTypes.MarketIdentificationCode{
					MIC:                   "XNYS",
					Operating_MIC:         "",
					OPRT_SGMT:             "",
					MarketName:            "",
					LegalEntityName:       "",
					LegalEntityIdentifier: "",
					MarketCategory:        "",
					Acronym:               "",
					ISOCountryCode:        "",
					City:                  "",
					Website:               "",
					Status:                "",
					CreationDate:          "",
					LastUpdateDate:        "",
					LastValidationDate:    "",
					ExpiryDate:            "",
					Comments:              "",
					Creator:               suite.address[1].String(),
				},
				quoteRequest: types.QuoteRequest{
					QuoteReqID:       "aaaa",
					Symbol:           "AAPL",
					SecurityID:       "123456",
					SecurityIDSource: "CUSIP",
					Side:             "BUY",
					OrderQty:         "100",
					FutSettDate:      "20230410",
					SettlDate2:       "20230412",
					Account:          "ACCT123",
					BidPx:            "145.50",
					OfferPx:          "145.60",
					Currency:         "USD",
					ValidUntilTime:   time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
					ExpireTime:       time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
					QuoteType:        "IND",
					BidSize:          "50",
					OfferSize:        "75",
					Mic:              "XNYS",
					Text:             "Example quote request",
					Creator:          suite.address[1].String(),
				},
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Account Address not involved in a logon session can not send a Quote Request",
			args{
				creator: suite.address[2].String(),
				session: types.Sessions{
					SessionID: "xhyhjjjk",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					InitiatorAddress: string(suite.address[0].String()),
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					AcceptorAddress: string(suite.address[1].String()),
					Status:          "loggedIn",
					IsAccepted:      true,
				},
				marketIdentificationCode: micTypes.MarketIdentificationCode{
					MIC:                   "XNYS",
					Operating_MIC:         "",
					OPRT_SGMT:             "",
					MarketName:            "",
					LegalEntityName:       "",
					LegalEntityIdentifier: "",
					MarketCategory:        "",
					Acronym:               "",
					ISOCountryCode:        "",
					City:                  "",
					Website:               "",
					Status:                "",
					CreationDate:          "",
					LastUpdateDate:        "",
					LastValidationDate:    "",
					ExpiryDate:            "",
					Comments:              "",
					Creator:               suite.address[0].String(),
				},
				quoteRequest: types.QuoteRequest{
					QuoteReqID:       "aaaa",
					Symbol:           "AAPL",
					SecurityID:       "123456",
					SecurityIDSource: "CUSIP",
					Side:             "BUY",
					OrderQty:         "100",
					FutSettDate:      "20230410",
					SettlDate2:       "20230412",
					Account:          "ACCT123",
					BidPx:            "145.50",
					OfferPx:          "145.60",
					Currency:         "USD",
					ValidUntilTime:   time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
					ExpireTime:       time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
					QuoteType:        "IND",
					BidSize:          "50",
					OfferSize:        "75",
					Mic:              "XNYS",
					Text:             "Example quote request",
					Creator:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Valid SessionID but logon is not yet established between both parties",
			args{
				creator: suite.address[0].String(),
				session: types.Sessions{
					SessionID: "xhyhjjjk",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					InitiatorAddress: string(suite.address[0].String()),
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					AcceptorAddress: string(suite.address[1].String()),
					Status:          "logon-request",
					IsAccepted:      false,
				},
				marketIdentificationCode: micTypes.MarketIdentificationCode{
					MIC:                   "XNYS",
					Operating_MIC:         "",
					OPRT_SGMT:             "",
					MarketName:            "",
					LegalEntityName:       "",
					LegalEntityIdentifier: "",
					MarketCategory:        "",
					Acronym:               "",
					ISOCountryCode:        "",
					City:                  "",
					Website:               "",
					Status:                "",
					CreationDate:          "",
					LastUpdateDate:        "",
					LastValidationDate:    "",
					ExpiryDate:            "",
					Comments:              "",
					Creator:               suite.address[0].String(),
				},
				quoteRequest: types.QuoteRequest{
					QuoteReqID:       "aaaa",
					Symbol:           "AAPL",
					SecurityID:       "123456",
					SecurityIDSource: "CUSIP",
					Side:             "BUY",
					OrderQty:         "100",
					FutSettDate:      "20230410",
					SettlDate2:       "20230412",
					Account:          "ACCT123",
					BidPx:            "145.50",
					OfferPx:          "145.60",
					Currency:         "USD",
					ValidUntilTime:   time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
					ExpireTime:       time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
					QuoteType:        "IND",
					BidSize:          "50",
					OfferSize:        "75",
					Mic:              "XNYS",
					Text:             "Example quote request",
					Creator:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"QuoteReqID is Empty",
			args{
				creator: suite.address[0].String(),
				session: types.Sessions{
					SessionID: "xhyhjjjk",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					InitiatorAddress: string(suite.address[0].String()),
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   54,
							MsgSeqNum:    1,
							MsgType:      "R",
							SenderCompID: "did:re:ntzwRLvyUyo9pWPofjuBzED9FgL2XHzHRWgnMzG1CCq",
							SendingTime:  "20230404-01:17:48.322",
							TargetCompID: "did:re:7BayfLBBR4JAnPsxiQDNGZsWA3YrcMqhrvcBsTzYp4HN",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					AcceptorAddress: string(suite.address[1].String()),
					Status:          "loggedIn",
					IsAccepted:      true,
				},
				marketIdentificationCode: micTypes.MarketIdentificationCode{
					MIC:                   "XNYS",
					Operating_MIC:         "",
					OPRT_SGMT:             "",
					MarketName:            "",
					LegalEntityName:       "",
					LegalEntityIdentifier: "",
					MarketCategory:        "",
					Acronym:               "",
					ISOCountryCode:        "",
					City:                  "",
					Website:               "",
					Status:                "",
					CreationDate:          "",
					LastUpdateDate:        "",
					LastValidationDate:    "",
					ExpiryDate:            "",
					Comments:              "",
					Creator:               suite.address[0].String(),
				},
				quoteRequest: types.QuoteRequest{
					QuoteReqID:       "",
					Symbol:           "AAPL",
					SecurityID:       "123456",
					SecurityIDSource: "CUSIP",
					Side:             "BUY",
					OrderQty:         "100",
					FutSettDate:      "20230410",
					SettlDate2:       "20230412",
					Account:          "ACCT123",
					BidPx:            "145.50",
					OfferPx:          "145.60",
					Currency:         "USD",
					ValidUntilTime:   time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
					ExpireTime:       time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
					QuoteType:        "IND",
					BidSize:          "50",
					OfferSize:        "75",
					Mic:              "XNYS",
					Text:             "Example quote request",
					Creator:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Empty Session",
			args{
				creator: suite.address[0].String(),
				session: types.Sessions{
					SessionID: " ",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  " ",
							BodyLength:   0,
							MsgSeqNum:    0,
							MsgType:      " ",
							SenderCompID: " ",
							SendingTime:  " ",
							TargetCompID: " ",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 0,
						},
					},
					InitiatorAddress: "",
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  " ",
							BodyLength:   0,
							MsgSeqNum:    1,
							MsgType:      " ",
							SenderCompID: " ",
							SendingTime:  " ",
							TargetCompID: " ",
						},
						HeartBtInt:    20,
						EncryptMethod: 0,
						Trailer: &types.Trailer{
							CheckSum: 0,
						},
					},
					AcceptorAddress: "",
					Status:          " ",
					IsAccepted:      false,
				},
				marketIdentificationCode: micTypes.MarketIdentificationCode{
					MIC:                   "XNYS",
					Operating_MIC:         "",
					OPRT_SGMT:             "",
					MarketName:            "",
					LegalEntityName:       "",
					LegalEntityIdentifier: "",
					MarketCategory:        "",
					Acronym:               "",
					ISOCountryCode:        "",
					City:                  "",
					Website:               "",
					Status:                "",
					CreationDate:          "",
					LastUpdateDate:        "",
					LastValidationDate:    "",
					ExpiryDate:            "",
					Comments:              "",
					Creator:               suite.address[0].String(),
				},
				quoteRequest: types.QuoteRequest{
					QuoteReqID:       "aaaa",
					Symbol:           "AAPL",
					SecurityID:       "123456",
					SecurityIDSource: "CUSIP",
					Side:             "BUY",
					OrderQty:         "100",
					FutSettDate:      "20230410",
					SettlDate2:       "20230412",
					Account:          "ACCT123",
					BidPx:            "145.50",
					OfferPx:          "145.60",
					Currency:         "USD",
					ValidUntilTime:   time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
					ExpireTime:       time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
					QuoteType:        "IND",
					BidSize:          "50",
					OfferSize:        "75",
					Mic:              "XNYS",
					Text:             "Example quote request",
					Creator:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// set session
			suite.fixKeeper.SetSessions(suite.ctx, tc.args.session.SessionID, tc.args.session)

			// get session
			session, _ := suite.fixKeeper.GetSessions(suite.ctx, tc.args.session.SessionID)

			// set mic from mic module
			suite.micKeeper.SetMarketIdentificationCode(suite.ctx, tc.args.marketIdentificationCode)

			// get mic from mic module
			mic, _ := suite.micKeeper.GetMarketIdentificationCode(suite.ctx, tc.args.marketIdentificationCode.MIC)

			// set QuoteRequest
			quoteRequest := types.NewQuoteRequest(tc.args.quoteRequest.QuoteReqID, tc.args.quoteRequest.Symbol, tc.args.quoteRequest.SecurityID, tc.args.quoteRequest.SecurityIDSource, tc.args.quoteRequest.Side, tc.args.quoteRequest.OrderQty, tc.args.quoteRequest.FutSettDate, tc.args.quoteRequest.SettlDate2, tc.args.quoteRequest.Account, tc.args.quoteRequest.BidPx, tc.args.quoteRequest.OfferPx, tc.args.quoteRequest.Currency, tc.args.quoteRequest.ValidUntilTime, tc.args.quoteRequest.ExpireTime, tc.args.quoteRequest.QuoteType, tc.args.quoteRequest.BidSize, tc.args.quoteRequest.OfferSize, mic.MIC, tc.args.quoteRequest.Text, string(tc.args.creator))

			// NewMsgQuoteRequest instance
			msg := types.NewMsgQuoteRequest(tc.args.creator, session.SessionID, quoteRequest)

			// call QuoteRequest method
			res, err := suite.msgServer.QuoteRequest(sdk.WrapSDKContext(suite.ctx), msg)

			// get Quote if QuoteRequest method created a Quote Request
			getQuote, found := suite.fixKeeper.GetQuote(suite.ctx, tc.args.quoteRequest.QuoteReqID)

			if tc.errArgs.shouldPass {
				suite.Require().True(found)
				suite.Require().NoError(err, tc.name)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(getQuote)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Nil(res)
				suite.Require().Empty(getQuote)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}
