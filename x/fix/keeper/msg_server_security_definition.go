package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityDefinition(goCtx context.Context, msg *types.MsgSecurityDefinition) (*types.MsgSecurityDefinitionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityDefinitionResponse{}, nil
}
