package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityListRequestReject(goCtx context.Context, msg *types.MsgSecurityListRequestReject) (*types.MsgSecurityListRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityListRequestRejectResponse{}, nil
}
