package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) OrderCancelReject(goCtx context.Context, msg *types.MsgOrderCancelReject) (*types.MsgOrderCancelRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	session, _ := k.GetSessions(ctx, msg.SessionID)

	// TODO: Handling the message
	//Fetch session
	// check for existing origClOrdID
	// check that account address equals the creator
	//Handle all FIX protocol message

	orderCancelReject := types.OrdersCancelReject{
		SessionID:        msg.SessionID,
		Header:           session.LogonAcceptor.Header,
		OrderID:          msg.OrderID,
		OrigClOrdID:      msg.OrigClOrdID,
		ClOrdID:          msg.ClOrdID,
		CxlRejReason:     msg.CxlRejReason,
		CxlRejResponseTo: msg.CxlRejResponseTo,
		TransactTime:     msg.TransactTime,
		Trailer:          session.LogonAcceptor.Trailer,
		Creator:          msg.Creator,
	}

	//set msgType
	orderCancelReject.Header.MsgType = "9"
	orderCancelReject.TransactTime = time.Now().UTC().Format("20060102-15:04:05.000")

	k.SetOrdersCancelReject(ctx, msg.SessionID, orderCancelReject)

	return &types.MsgOrderCancelRejectResponse{}, nil
}
