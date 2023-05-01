package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) MarketDataRequest(goCtx context.Context, msg *types.MsgMarketDataRequest) (*types.MsgMarketDataRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != types.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrQuoteSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Quote Request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that mandatory Market Data Request fields are not empty
	if _, err := strconv.ParseInt(fmt.Sprint(msg.SubscriptionRequestType), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrSubscriptionRequestTypeIsEmpty, "SubscriptionRequestType: %v", msg.SubscriptionRequestType)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.MarketDepth), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrMarketDepthIsEmpty, "MarketDepth: %v", msg.MarketDepth)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.MdUpdateType), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrMdUpdateTypeIsEmpty, "MdUpdateType: %v", msg.MdUpdateType)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.NoRelatedSym), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrNoRelatedSymIsEmpty, "NoRelatedSym: %v", msg.NoRelatedSym)
	}
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrSymbolIsEmpty, "Symbol: %s", msg.Symbol)
	}

	// market data request
	marketDataRequest := types.MarketData{
		SessionID: msg.SessionID,
		MarketDataRequest: &types.MarketDataRequest{
			MdReqID:                 msg.MdReqID,
			SubscriptionRequestType: msg.SubscriptionRequestType,
			MarketDepth:             msg.MarketDepth,
			MdUpdateType:            msg.MdUpdateType,
			NoRelatedSym:            msg.NoRelatedSym,
			Symbol:                  msg.Symbol,
			Creator:                 msg.Creator,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, a Quote Request message can be sent by either the initiator or the acceptor of the FIX session.
	// Determine whether we are the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, V is the message type for Quote Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating market data request
	marketDataRequest.MarketDataRequest.Header = header
	marketDataRequest.MarketDataRequest.Header.MsgType = "V"
	marketDataRequest.MarketDataRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	// for now copy trailer from session, it should be re-calculated
	var trailer *types.Trailer
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// checksum in the trailer can be recalculated using CalculateChecksum function
	marketDataRequest.MarketDataRequest.Trailer = trailer

	// set new market data request to store
	k.SetMarketData(ctx, msg.MdReqID, marketDataRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgMarketDataRequestResponse{}, err
}

func (k msgServer) MarketDataSnapshotFullRefresh(goCtx context.Context, msg *types.MsgMarketDataSnapshotFullRefresh) (*types.MsgMarketDataSnapshotFullRefreshResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataSnapshotFullRefreshResponse{}, nil
}

func (k msgServer) MarketDataIncremental(goCtx context.Context, msg *types.MsgMarketDataIncremental) (*types.MsgMarketDataIncrementalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataIncrementalResponse{}, nil
}

func (k msgServer) MarketDataRequestReject(goCtx context.Context, msg *types.MsgMarketDataRequestReject) (*types.MsgMarketDataRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataRequestRejectResponse{}, nil
}
