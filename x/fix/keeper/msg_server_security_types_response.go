package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityTypesResponse(goCtx context.Context, msg *types.MsgSecurityTypesResponse) (*types.MsgSecurityTypesResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityTypesResponseResponse{}, nil
}
