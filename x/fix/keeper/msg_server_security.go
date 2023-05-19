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


func (k msgServer) SecurityDefinition(goCtx context.Context, msg *types.MsgSecurityDefinition) (*types.MsgSecurityDefinitionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityDefinitionResponse{}, nil
}

func (k msgServer) SecurityDefinitionRequestReject(goCtx context.Context, msg *types.MsgSecurityDefinitionRequestReject) (*types.MsgSecurityDefinitionRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityDefinitionRequestRejectResponse{}, nil
}
