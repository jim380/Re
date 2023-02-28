package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) NewOrderSingle(goCtx context.Context, msg *types.MsgNewOrderSingle) (*types.MsgNewOrderSingleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//TODO
	// get session, check if session exists and if the status is set to loggedIn
	// check for account address
	// include checks for all the field messages
	// add more check cases

	//check for if this session Name exists
	session, found := k.GetSessions(ctx, msg.SessionName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionName)
	}
   if session.Status != "loggedIn" {
	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s state is not logged in", msg.SessionName))
	 
   }
	//check if order exists
	order, found := k.GetOrders(ctx, msg.SessionName)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order exist already", &order))
	}


	newOrder := types.Orders{
		SessionName:  session.SessionName,
		Header:       session.LogonInitiator.Header,
		ClOrdID:      msg.ClOrdID,
		Symbol:       msg.Symbol,
		Side:         msg.Side,
		OrderQty:     msg.OrderQty,
		OrdType:      msg.OrdType,
		Price:        msg.Price,
		TimeInForce:  msg.TimeInForce,
		Text:         msg.Text,
		TransactTime: msg.TransactTime,
		Trailer:      session.LogonInitiator.Trailer,
		Status:       msg.Status,
		Creator:      msg.Creator,
	}
   
	//set the msgType
	newOrder.Header.MsgType = "D"

   k.SetOrders(ctx, msg.SessionName, newOrder)

	return &types.MsgNewOrderSingleResponse{}, nil
}
