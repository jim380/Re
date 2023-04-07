package keeper_test

import (
	"strings"
	"testing"
	"time"

	sdksimapp "github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	reapp "github.com/jim380/Re/app"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	denom = "stake"
	acc1  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc3  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc4  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

// shared setup
type KeeperTestSuite struct {
	suite.Suite

	address   []sdk.AccAddress
	app       *reapp.ReApp
	ctx       sdk.Context
	keeper    keeper.Keeper
	msgServer types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {

	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = suite.app.FixKeeper
	//suite.msgServer = keeper.NewMsgServerImpl(*&suite.keeper)

	for _, acc := range []sdk.AccAddress{acc1, acc2, acc3, acc4} {
		err := sdksimapp.FundAccount(
			suite.app.BankKeeper,
			suite.ctx,
			acc,
			sdk.NewCoins(
				sdk.NewCoin(denom, sdk.NewInt(1000)),
			),
		)
		if err != nil {
			panic(err)
		}
	}

	suite.address = []sdk.AccAddress{acc1, acc2, acc3, acc4}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestQuoteRequest() {

	type quoteRequest struct {
		quoteReqID       string
		symbol           string
		securityID       string
		securityIDSource string
		side             string
		orderQty         string
		futSettDate      string
		settlDate2       string
		account          string
		bidPx            string
		offerPx          string
		currency         string
		validUntilTime   string
		expireTime       string
		quoteType        string
		bidSize          string
		offerSize        string
		mic              string
		text             string
		creator          sdk.AccAddress
	}

	type args struct {
		creator      sdk.AccAddress
		sessionID    string
		quoteRequest quoteRequest
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
		{"Valid SessionID in QuoteRequest",
			args{
				creator:   suite.address[0],
				sessionID: "qwjhsjdudhd",
				quoteRequest: quoteRequest{
					quoteReqID:       "aaaa",
					symbol:           "AAPL",
					securityID:       "123456",
					securityIDSource: "CUSIP",
					side:             "BUY",
					orderQty:         "100",
					futSettDate:      "20230410",
					settlDate2:       "20230412",
					account:          "ACCT123",
					bidPx:            "145.50",
					offerPx:          "145.60",
					currency:         "USD",
					validUntilTime:   time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
					expireTime:       time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
					quoteType:        "IND",
					bidSize:          "50",
					offerSize:        "75",
					mic:              "XNYS",
					text:             "Example quote request",
					creator:          suite.address[0],
				},
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

			// set QuoteRequest
			quoteRequest := types.NewQuoteRequest(tc.args.quoteRequest.quoteReqID, tc.args.quoteRequest.symbol, tc.args.quoteRequest.securityID, tc.args.quoteRequest.securityIDSource, tc.args.quoteRequest.side, tc.args.quoteRequest.orderQty, tc.args.quoteRequest.futSettDate, tc.args.quoteRequest.settlDate2, tc.args.quoteRequest.account, tc.args.quoteRequest.bidPx, tc.args.quoteRequest.offerPx, tc.args.quoteRequest.currency, tc.args.quoteRequest.validUntilTime, tc.args.quoteRequest.expireTime, tc.args.quoteRequest.quoteType, tc.args.quoteRequest.bidSize, tc.args.quoteRequest.offerSize, tc.args.quoteRequest.mic, tc.args.quoteRequest.text, string(tc.args.creator))

			quoteRequest.Header = &types.Header{}

			// set session
			//session := types.Sessions {
			//SessionID: tc.args.sessionID,
			// LogonInitiator: &types.LogonInitiator{},
			//InitiatorAddress: string(suite.address[0]),
			//LogonAcceptor: &types.LogonAcceptor{},
			//AcceptorAddress: string(suite.address[1]),
			//Status: "loggedIn",
			//IsAccepted: true,
			//}
			//suite.keeper.SetSessions(suite.ctx, tc.args.sessionID, session)
			// get session
			session, found := suite.keeper.GetSessions(suite.ctx, tc.args.sessionID)

			// msg
			msg := types.NewMsgQuoteRequest(tc.args.creator.String(), session.SessionID, quoteRequest)

			// call QuoteRequest method
			res, err := suite.msgServer.QuoteRequest(sdk.WrapSDKContext(suite.ctx), msg)

			if tc.errArgs.shouldPass {
				suite.Require().True(found)
				suite.Require().NoError(err, tc.name)
				suite.Require().NotEmpty(res)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Empty(res)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})

	}

}

