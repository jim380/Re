package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

// SecurityListRequest creates Security List Request
func (k msgServer) SecurityListRequest(goCtx context.Context, msg *types.MsgSecurityListRequest) (*types.MsgSecurityListRequestResponse, error) {
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
	if session.Status != constants.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Security List Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that the mandatory Security list Request field is not empty
	if msg.SecurityListRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListEmptyField, "SecurityListRequestType: %s", msg.SecurityListRequestType)
	}

	// security list request
	securityListRequest := types.SecurityList{
		SessionID: msg.SessionID,
		SecurityListRequest: &types.SecurityListRequest{
			SecurityReqID:           msg.SecurityReqID,
			SecurityListRequestType: msg.SecurityListRequestType,
			NoUnderlyings:           msg.NoUnderlyings,
			NoLegs:                  msg.NoLegs,
			Currency:                msg.Currency,
			Text:                    msg.Text,
			EncodedTextLen:          msg.EncodedTextLen,
			EncodedText:             msg.EncodedText,
			TradingSessionID:        msg.TradingSessionID,
			TradingSessionSubID:     msg.TradingSessionSubID,
			SubscriptionRequestType: msg.SubscriptionRequestType,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, a Security List Request message can be sent by either the initiator or the acceptor of the FIX session.
	// Determine whether we are the initiator or acceptor
	var header *types.Header
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, x is the message type for Security List Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Security List Request
	securityListRequest.SecurityListRequest.Header = header
	securityListRequest.SecurityListRequest.Header.MsgType = "x"
	securityListRequest.SecurityListRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	var trailer *types.Trailer
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// checksum in the trailer can be recalculated using CalculateChecksum function
	securityListRequest.SecurityListRequest.Trailer = trailer

	// set Security List Request to store
	k.SetSecurityList(ctx, msg.SecurityReqID, securityListRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityListRequestResponse{}, err
}

// SecurityListResponse creates Security List Response
func (k msgServer) SecurityListResponse(goCtx context.Context, msg *types.MsgSecurityListResponse) (*types.MsgSecurityListResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Security List Response matches with the sessionID for Security List Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user to acknowledge the Security List Request is the recipient of the Security List Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security List Response Creator: %s", msg.Creator)
	}

	// get the existing Security List Request
	securityList, found := k.GetSecurityList(ctx, msg.SecurityReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListIsNotFound, "Security List: %s", &securityList)
	}

	// check that Security List Request creator address is not same address acknowledging the Security List Response with the same SecurityReqID
	if securityList.SecurityListRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Security List Response", msg.Creator))
	}

	// check that this Security List Request to be acknowledged has not been rejected
	if securityList.SecurityListRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListRequestIsRejected, "Seurity List: %s", securityList.SecurityListRequestReject)
	}

	// check that this Security List is not acknowledged already
	if securityList.SecurityListResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListRequestIsAcknowledged, "Security List: %s", securityList.SecurityListResponse)
	}

	// check that the mandatory Security list Response field is not empty
	if msg.SecurityResponseID == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListEmptyField, "SecurityResponseID: %s", msg.SecurityResponseID)
	}
	if msg.SecurityRequestResult == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListEmptyField, "SecurityRequestResult: %s", msg.SecurityRequestResult)
	}

	securityListResponse := types.SecurityList{
		SessionID:           msg.SessionID,
		SecurityListRequest: securityList.SecurityListRequest,
		SecurityListResponse: &types.SecurityListResponse{
			SecurityReqID:         msg.SecurityReqID,
			SecurityResponseID:    msg.SecurityResponseID,
			SecurityRequestResult: msg.SecurityRequestResult,
			TotNoRelatedSym:       msg.TotNoRelatedSym,
			LastFragment:          msg.LastFragment,
			NoRelatedSym:          msg.NoRelatedSym,
			NoUnderlyings:         msg.NoUnderlyings,
			Currency:              msg.Currency,
			NoLegs:                msg.NoLegs,
			RoundLot:              msg.RoundLot,
			MinTradeVol:           msg.MinTradeVol,
			TradingSessionID:      msg.TradingSessionID,
			TradingSessionSubID:   msg.TradingSessionSubID,
			ExpirationCycle:       msg.ExpirationCycle,
			Text:                  msg.Text,
			EncodedTextLen:        msg.EncodedTextLen,
			EncodedText:           msg.EncodedText,
		},
	}

	// set header from the existing security list request
	securityListResponse.SecurityListResponse.Header = securityList.SecurityListRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *securityListResponse.SecurityListResponse.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = securityList.SecurityListRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = securityList.SecurityListRequest.Header.MsgSeqNum

	// set the msgType to Security List Response
	newHeader.MsgType = "y"

	// switch senderCompID and targetCompID between Security List Request and Security List Response
	// set senderCompID of Security List Response to the targetCompID of Security List Request in the header
	newHeader.SenderCompID = securityList.SecurityListRequest.Header.TargetCompID

	// set targetCompID of Security List Response to the senderCompID of Security List Request in the header
	newHeader.TargetCompID = securityList.SecurityListRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	securityListResponse.SecurityListResponse.Header = newHeader

	// set Trailer from the existing Security List Request
	securityListResponse.SecurityListResponse.Trailer = securityList.SecurityListRequest.Trailer

	// set Security List Response to store
	k.SetSecurityList(ctx, msg.SecurityReqID, securityListResponse)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityListResponseResponse{}, err
}

// SecurityListRequestReject creates Security List Request Reject
func (k msgServer) SecurityListRequestReject(goCtx context.Context, msg *types.MsgSecurityListRequestReject) (*types.MsgSecurityListRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Security List Request Reject matches with the sessionID for Security List Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user to acknowledge the Security List Request is the recipient of the Security List Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security List Request Reject Creator: %s", msg.Creator)
	}

	// get the existing Security List Request
	securityList, found := k.GetSecurityList(ctx, msg.SecurityReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListIsNotFound, "Security List: %s", &securityList)
	}

	// check that Security List Request creator address is not same address acknowledging the Security List Request Reject with the same SecurityReqID
	if securityList.SecurityListRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Security List Request Reject", msg.Creator))
	}

	// check that this Security List Request to be acknowledged has not been rejected
	if securityList.SecurityListRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListRequestIsRejected, "Seurity List: %s", securityList.SecurityListRequestReject)
	}

	// check that this Security List is not acknowledged already
	if securityList.SecurityListResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListRequestIsAcknowledged, "Security List: %s", securityList.SecurityListResponse)
	}

	// check that the mandatory Security list Request Reject fields are not empty
	if msg.SecurityListRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListEmptyField, "SecurityListRequestType: %s", msg.SecurityListRequestType)
	}
	if msg.SecurityRequestResult == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListEmptyField, "SecurityRequestResult: %s", msg.SecurityRequestResult)
	}

	securityListRequestReject := types.SecurityList{
		SessionID:            msg.SessionID,
		SecurityListRequest:  securityList.SecurityListRequest,
		SecurityListResponse: securityList.SecurityListResponse,
		SecurityListRequestReject: &types.SecurityListRequestReject{
			SecurityReqID:           msg.SecurityReqID,
			SecurityListRequestType: msg.SecurityListRequestType,
			SecurityRequestResult:   msg.SecurityRequestResult,
			Text:                    msg.Text,
			EncodedTextLen:          msg.EncodedTextLen,
			EncodedText:             msg.EncodedText,
		},
	}

	// set header from the existing security list request
	securityListRequestReject.SecurityListRequestReject.Header = securityList.SecurityListRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *securityListRequestReject.SecurityListRequestReject.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = securityList.SecurityListRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = securityList.SecurityListRequest.Header.MsgSeqNum

	// set the msgType to Security List Request Reject
	newHeader.MsgType = "y"

	// switch senderCompID and targetCompID between Security List Request and Security List Request Reject
	// set senderCompID of Security List Request Reject to the targetCompID of Security List Request in the header
	newHeader.SenderCompID = securityList.SecurityListRequest.Header.TargetCompID

	// set targetCompID of Security List Request Reject to the senderCompID of Security List Request in the header
	newHeader.TargetCompID = securityList.SecurityListRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	securityListRequestReject.SecurityListRequestReject.Header = newHeader

	// set Trailer from the existing Security List Request
	securityListRequestReject.SecurityListRequestReject.Trailer = securityList.SecurityListRequest.Trailer

	// set Security List Request Reject to store
	k.SetSecurityList(ctx, msg.SecurityReqID, securityListRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityListRequestRejectResponse{}, err
}
