package keeper

import (
	"context"

    "github.com/jim380/Re/x/fix/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) QuoteRequest(goCtx context.Context,  msg *types.MsgQuoteRequest) (*types.MsgQuoteRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgQuoteRequestResponse{}, nil
}
