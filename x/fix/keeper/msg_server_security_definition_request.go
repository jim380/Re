package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityDefinitionRequest(goCtx context.Context, msg *types.MsgSecurityDefinitionRequest) (*types.MsgSecurityDefinitionRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityDefinitionRequestResponse{}, nil
}
