package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/oracle/types"
)

func (k msgServer) CosmoshubTxs(goCtx context.Context, msg *types.MsgCosmoshubTxs) (*types.MsgCosmoshubTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCosmoshubTxsResponse{}, nil
}
