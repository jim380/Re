package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

// MarketDataRequest creates Market Data Request
func (k msgServer) MarketDataRequest(goCtx context.Context, msg *types.MsgMarketDataRequest) (*types.MsgMarketDataRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != constants.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Market Data Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
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
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Market Data Request message can be sent by either the initiator or the acceptor of the FIX session.
	// Determine whether we are the initiator or acceptor
	var header *types.Header
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, V is the message type for Market Data Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating market data request
	marketDataRequest.MarketDataRequest.Header = header
	marketDataRequest.MarketDataRequest.Header.MsgType = "V"
	marketDataRequest.MarketDataRequest.Header.SendingTime = constants.SendingTime

	// fetch Trailer from existing session
	// for now copy trailer from session, it should be re-calculated
	var trailer *types.Trailer
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
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

// MarketDataSnapshotFullRefresh creates Market Data Snapshot Full Refresh
func (k msgServer) MarketDataSnapshotFullRefresh(goCtx context.Context, msg *types.MsgMarketDataSnapshotFullRefresh) (*types.MsgMarketDataSnapshotFullRefreshResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that the sessionID provided is active between both parties
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the market data request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Market Data Snap Shot Full Refresh Creator: %s", msg.Creator)
	}

	// get the existing Market Data Request
	marketDataRequest, found := k.GetMarketData(ctx, msg.MdReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMdReqIDIsNotFound, "MdReqID: %s", msg.MdReqID)
	}

	// check that Market Data Request creator address is not same responding to the Market Data Request
	if marketDataRequest.MarketDataRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataSnapShotFullRefreshCreatorIsWrong, "Market Data Snapshot Full Refresh: %s", msg.Creator)
	}

	// check that this Market Data Request to be acknowledged has not been rejected
	// if  MarketDataRequestReject is not nil, then market data snap shot full refresh should be rejected
	if marketDataRequest.MarketDataRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataRequestIsRejected, "Market Data Request: %s", marketDataRequest.MarketDataRequestReject)
	}

	// check that this Market Data Request is not acknowledged already
	if marketDataRequest.MarketDataSnapShotFullRefresh != nil {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataRequestIsAcknowlodged, "Market Data Request: %s", marketDataRequest.MarketDataSnapShotFullRefresh)
	}

	// check that mandatory Market Data Snap Shot Full Refresh fields are not empty
	if msg.MdReqID == "" {
		return nil, sdkerrors.Wrapf(types.ErrMdReqIDIsEmpty, "MdReqID: %s", msg.MdReqID)
	}
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrSymbolIsEmpty, "Symbol: %s", msg.Symbol)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.NoMDEntries), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrMarketDepthIsEmpty, "NoMDEntries: %v", msg.NoMDEntries)
	}
	for _, md := range msg.MdEntries {
		if _, err := strconv.ParseInt(fmt.Sprint(md.MdUpdateAction), 10, 64); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrMdUpdateActionIsEmpty, "MdUpdateAction: %v", md.MdUpdateAction)
		}
		if _, err := strconv.ParseInt(fmt.Sprint(md.MdEntryType), 10, 64); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrMdEntryTypeIsEmpty, "MdEntryType: %v", md.MdEntryType)
		}
		if md.MdEntryPx == "" {
			return nil, sdkerrors.Wrapf(types.ErrMdEntryPxIsEmpty, "MdEntryPx: %s", md.MdEntryPx)
		}
		if md.MdEntrySize == "" {
			return nil, sdkerrors.Wrapf(types.ErrMdEntrySizeIsEmpty, "MdEntrySize: %s", md.MdEntrySize)
		}
	}

	// market data snap shot full refresh
	marketDataSnapShotFullRefresh := types.MarketData{
		SessionID:         msg.SessionID,
		MarketDataRequest: marketDataRequest.MarketDataRequest,
		MarketDataSnapShotFullRefresh: &types.MarketDataSnapShotFullRefresh{
			MdReqID:     msg.MdReqID,
			Symbol:      msg.Symbol,
			NoMDEntries: msg.NoMDEntries,
			MdEntries:   msg.MdEntries,
		},
	}

	// set header from the existing Market Data Request
	marketDataSnapShotFullRefresh.MarketDataSnapShotFullRefresh.Header = marketDataRequest.MarketDataRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *marketDataSnapShotFullRefresh.MarketDataSnapShotFullRefresh.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = marketDataRequest.MarketDataRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = marketDataRequest.MarketDataRequest.Header.MsgSeqNum

	// set the msgType to Market Data Snapshot/Full Refresh
	newHeader.MsgType = "W"

	// switch senderCompID and targetCompID between Market Data Request and MArket Data Snapshot Full Refresh
	// set senderCompID of Market Data Snapshot Full Refresh to the targetCompID of Market Data Request in the header
	newHeader.SenderCompID = marketDataRequest.MarketDataRequest.Header.TargetCompID

	// set targetCompID of Market Data Snapshot Full Refresh to the senderCompID of Market Data Request in the header
	newHeader.TargetCompID = marketDataRequest.MarketDataRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	marketDataSnapShotFullRefresh.MarketDataSnapShotFullRefresh.Header = newHeader

	// set Trailer from the existing Market Data Request
	marketDataSnapShotFullRefresh.MarketDataSnapShotFullRefresh.Trailer = marketDataRequest.MarketDataRequest.Trailer

	// set new market data snapshot full refresh to store
	k.SetMarketData(ctx, msg.MdReqID, marketDataSnapShotFullRefresh)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgMarketDataSnapshotFullRefreshResponse{}, err
}

// MarketDataIncremental creates Market Data Incremental
func (k msgServer) MarketDataIncremental(goCtx context.Context, msg *types.MsgMarketDataIncremental) (*types.MsgMarketDataIncrementalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataIncrementalResponse{}, nil
}

// MarketDataRequestReject creates Market Data Request Reject
func (k msgServer) MarketDataRequestReject(goCtx context.Context, msg *types.MsgMarketDataRequestReject) (*types.MsgMarketDataRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that the sessionID provided is active between botth parties
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the market data request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Market Data Request Creator: %s", msg.Creator)
	}

	// get the existing Market Data Request
	marketDataRequest, found := k.GetMarketData(ctx, msg.MdReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMdReqIDIsNotFound, "MdReqID: %s", msg.MdReqID)
	}

	// check that Market Data Request creator address is not same responding to the Market Data Request Reject
	if marketDataRequest.MarketDataRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataRequestRejectCreatorIsWrong, "Market Data Request Reject: %s", msg.Creator)
	}

	// check that this Market Data Request to be rejected has not been acknowledged
	// if market data snap shot full refresh is not nil, then it means market data request has been acknowledged, hence market data request reject cannot be executed
	if marketDataRequest.MarketDataSnapShotFullRefresh != nil {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataRequestIsAcknowlodged, "Market Data Request Reject: %s", marketDataRequest.MarketDataSnapShotFullRefresh)
	}

	// check that this Market Data Request is not rejected already
	if marketDataRequest.MarketDataRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataRequestIsRejected, "Market Data Request: %s", marketDataRequest.MarketDataRequestReject)
	}

	// check that mandatory Market Data Request Reject field is not empty
	if _, err := strconv.ParseInt(fmt.Sprint(msg.MdReqRejReason), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrMdReqRejReasonIsEmpty, "MdReqRejReason: %v", msg.MdReqRejReason)
	}

	// market data request reject
	marketDataRequestReject := types.MarketData{
		SessionID:                     msg.SessionID,
		MarketDataRequest:             marketDataRequest.MarketDataRequest,
		MarketDataSnapShotFullRefresh: marketDataRequest.MarketDataSnapShotFullRefresh,
		MarketDataIncremental:         marketDataRequest.MarketDataIncremental,
		MarketDataRequestReject: &types.MarketDataRequestReject{
			MdReqID:        msg.MdReqID,
			MdReqRejReason: msg.MdReqRejReason,
			Text:           msg.Text,
		},
	}

	// set header from the existing Market Data Request
	marketDataRequestReject.MarketDataRequestReject.Header = marketDataRequest.MarketDataRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *marketDataRequestReject.MarketDataRequestReject.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = marketDataRequest.MarketDataRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = marketDataRequest.MarketDataRequest.Header.MsgSeqNum

	// set the msgType to Market Data Request Reject
	newHeader.MsgType = "Y"

	// switch senderCompID and targetCompID between Market Data Request and Market Data Request Reject
	// set senderCompID of Market Data Request Reject to the targetCompID of Market Data Request in the header
	newHeader.SenderCompID = marketDataRequest.MarketDataRequest.Header.TargetCompID

	// set targetCompID of Market Data Request Reject to the senderCompID of Market Data Request in the header
	newHeader.TargetCompID = marketDataRequest.MarketDataRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	marketDataRequestReject.MarketDataRequestReject.Header = newHeader

	// set Trailer from the existing Market Data Request
	// then calcualate the checksum
	marketDataRequestReject.MarketDataRequestReject.Trailer = marketDataRequest.MarketDataRequest.Trailer

	// set new market data request reject to store
	k.SetMarketData(ctx, msg.MdReqID, marketDataRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgMarketDataRequestRejectResponse{}, err
}
