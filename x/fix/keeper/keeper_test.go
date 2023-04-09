package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	reapp "github.com/jim380/Re/app"
	fixKeeper "github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
	micKeeper "github.com/jim380/Re/x/mic/keeper"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	denom            = "stake"
	initiatorAddress = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acceptorAddress  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	unusedAddress    = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

// shared setup
type KeeperTestSuite struct {
	suite.Suite

	address   []sdk.AccAddress
	app       *reapp.ReApp
	ctx       sdk.Context
	fixKeeper fixKeeper.Keeper
	micKeeper micKeeper.Keeper
	msgServer types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = reapp.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.fixKeeper = suite.app.FixKeeper
	suite.micKeeper = suite.app.MicKeeper
	suite.msgServer = fixKeeper.NewMsgServerImpl(suite.app.FixKeeper)

	suite.address = []sdk.AccAddress{initiatorAddress, acceptorAddress, unusedAddress}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
