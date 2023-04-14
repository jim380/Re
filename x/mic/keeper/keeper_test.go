package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	reapp "github.com/jim380/Re/app"
	micKeeper "github.com/jim380/Re/x/mic/keeper"
	"github.com/jim380/Re/x/mic/types"
)

var (
	denom = "stake"
	acc1  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc3  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

// shared setup
type KeeperTestSuite struct {
	suite.Suite

	address   []sdk.AccAddress
	app       *reapp.ReApp
	ctx       sdk.Context
	micKeeper micKeeper.Keeper
	msgServer types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = reapp.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.micKeeper = suite.app.MicKeeper
	suite.msgServer = micKeeper.NewMsgServerImpl(suite.app.MicKeeper)

	suite.address = []sdk.AccAddress{acc1, acc2, acc3}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
