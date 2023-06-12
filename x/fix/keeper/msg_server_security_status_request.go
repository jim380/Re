package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityStatusRequest(goCtx context.Context, msg *types.MsgSecurityStatusRequest) (*types.MsgSecurityStatusRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityStatusRequestResponse{}, nil
}
