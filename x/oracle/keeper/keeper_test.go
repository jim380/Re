package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	reapp "github.com/jim380/Re/app"
	fixKeeper "github.com/jim380/Re/x/fix/keeper"
	oracleKeeper "github.com/jim380/Re/x/oracle/keeper"
	"github.com/jim380/Re/x/oracle/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	address1 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	address2 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	address3 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

// shared setup
type KeeperTestSuite struct {
	suite.Suite

	address      []sdk.AccAddress
	app          *reapp.ReApp
	ctx          sdk.Context
	fixKeeper    fixKeeper.Keeper
	oracleKeeper oracleKeeper.Keeper
	msgServer    types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = reapp.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.fixKeeper = suite.app.FixKeeper
	suite.oracleKeeper = suite.app.OracleKeeper
	suite.msgServer = oracleKeeper.NewMsgServerImpl(suite.app.OracleKeeper)

	suite.address = []sdk.AccAddress{address1, address2, address3}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
