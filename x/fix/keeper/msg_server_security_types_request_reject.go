package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityTypesRequestReject(goCtx context.Context, msg *types.MsgSecurityTypesRequestReject) (*types.MsgSecurityTypesRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityTypesRequestRejectResponse{}, nil
}
