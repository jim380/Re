package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/oracle/types"
)

func (k msgServer) EthereumTxs(goCtx context.Context, msg *types.MsgEthereumTxs) (*types.MsgEthereumTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEthereumTxsResponse{}, nil
}
