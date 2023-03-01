package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) OrderCancelRequest(goCtx context.Context, msg *types.MsgOrderCancelRequest) (*types.MsgOrderCancelRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	session, _ := k.GetSessions(ctx, msg.SessionID)

	// TODO: Handling the message
	//Fetch session
	// check for existing origClOrdID
	// check that account address equals the creator

	//set order cancel request data
	orderCancelRequest := types.OrdersCancelRequest{
		SessionID:   msg.SessionID,
		Header:      session.LogonInitiator.Header,
		OrigClOrdID: msg.OrigClOrdID,
		ClOrdID:     msg.ClOrdID,
		Trailer:     session.LogonInitiator.Trailer,
		Creator:     msg.Creator,
	}

	// msgType = F
	orderCancelRequest.Header.MsgType = "F"

	//set order cancel request to store
	k.SetOrdersCancelRequest(ctx, msg.SessionID, orderCancelRequest)

	return &types.MsgOrderCancelRequestResponse{}, nil
}
