package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// SecurityTypesRequest creates Security Types Request
func (k msgServer) SecurityTypesRequest(goCtx context.Context, msg *types.MsgSecurityTypesRequest) (*types.MsgSecurityTypesRequestResponse, error) {
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
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Security Types Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	securityTypesRequest := types.SecurityTypes{
		SessionID: msg.SessionID,
		SecurityTypesRequest: &types.SecurityTypesRequest{
			SecurityReqID:       msg.SecurityReqID,
			Text:                msg.Text,
			TradingSessionID:    msg.TradingSessionID,
			TradingSessionSubID: msg.TradingSessionSubID,
			Product:             msg.Product,
			SecurityType:        msg.SecurityType,
			SecuritySubType:     msg.SecuritySubType,
		},
	}

	// fetch Header from existing session
	// Determine whether it is the initiator or acceptor
	var header *types.Header
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, v is the message type for Security Type Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Security Types Request
	securityTypesRequest.SecurityTypesRequest.Header = header
	securityTypesRequest.SecurityTypesRequest.Header.MsgType = "v"
	securityTypesRequest.SecurityTypesRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	var trailer *types.Trailer
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// checksum in the trailer can be recalculated using CalculateChecksum function
	securityTypesRequest.SecurityTypesRequest.Trailer = trailer

	// set Security Types Request to store
	k.SetSecurityTypes(ctx, msg.SecurityReqID, securityTypesRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityTypesRequestResponse{}, err
}

// SecurityTypesResponse creates Security Types Response
func (k msgServer) SecurityTypesResponse(goCtx context.Context, msg *types.MsgSecurityTypesResponse) (*types.MsgSecurityTypesResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Security Types Response matches with the sessionID for Security Types Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Security Types Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security Types Response Creator: %s", msg.Creator)
	}

	// get security Types
	securityTypes, found := k.GetSecurityTypes(ctx, msg.SecurityReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesIsNotFound, ": %s", securityTypes.SecurityTypesRequest)
	}

	// same account can not used for creating Security Types Request and Security Types Response with the same SecurityReqID
	if securityTypes.SecurityTypesRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Security Types Response", msg.Creator))
	}

	// check that the Security Types Request is not rejected already
	if securityTypes.SecurityTypesRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesRequestIsRejected, "Security Types: %s", securityTypes.SecurityTypesRequestReject)
	}

	// check that the Security Types Request is not acknowledged already
	if securityTypes.SecurityTypesResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesRequestIsAcknowledged, "Security Types: %s", securityTypes.SecurityTypesResponse)
	}

	securityTypesResponse := types.SecurityTypes{
		SessionID:            msg.SessionID,
		SecurityTypesRequest: securityTypes.SecurityTypesRequest,
		SecurityTypesResponse: &types.SecurityTypesResponse{
			SecurityReqID:           msg.SecurityReqID,
			SecurityResponseID:      msg.SecurityResponseID,
			SecurityResponseType:    msg.SecurityResponseType,
			TotNoSecurityTypes:      msg.TotNoSecurityTypes,
			LastFragment:            msg.LastFragment,
			NoSecurityTypes:         msg.NoSecurityTypes,
			SecurityType:            msg.SecurityType,
			SecuritySubType:         msg.SecuritySubType,
			Product:                 msg.Product,
			CFICode:                 msg.CFICode,
			Text:                    msg.Text,
			TradingSessionID:        msg.TradingSessionID,
			TradingSessionSubID:     msg.TradingSessionSubID,
			SubscriptionRequestType: msg.SubscriptionRequestType,
		},
	}

	// set header from the existing security types request
	securityTypesResponse.SecurityTypesResponse.Header = securityTypes.SecurityTypesRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *securityTypesResponse.SecurityTypesResponse.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = securityTypes.SecurityTypesRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = securityTypes.SecurityTypesRequest.Header.MsgSeqNum

	// set the msgType to Security Types Response
	newHeader.MsgType = "w"

	// switch senderCompID and targetCompID between Security Types Request and Security Types Response
	// set senderCompID of Security Types Response to the targetCompID of Security Types Request in the header
	newHeader.SenderCompID = securityTypes.SecurityTypesRequest.Header.TargetCompID

	// set targetCompID of Security Types Response to the senderCompID of Security Types Request in the header
	newHeader.TargetCompID = securityTypes.SecurityTypesRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	securityTypesResponse.SecurityTypesResponse.Header = newHeader

	// set Trailer from the existing Security Types Request
	securityTypesResponse.SecurityTypesResponse.Trailer = securityTypes.SecurityTypesRequest.Trailer

	// set Security Types Request to store
	k.SetSecurityTypes(ctx, msg.SecurityReqID, securityTypesResponse)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityTypesResponseResponse{}, err
}

// SecurityTypesRequestReject creates Security Types Request Reject
func (k msgServer) SecurityTypesRequestReject(goCtx context.Context, msg *types.MsgSecurityTypesRequestReject) (*types.MsgSecurityTypesRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Security Types Request Reject matches with the sessionID for Security Types Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Security Types Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security Types Request Reject Creator: %s", msg.Creator)
	}

	// get security Types
	securityTypes, found := k.GetSecurityTypes(ctx, msg.SecurityReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesIsNotFound, ": %s", securityTypes.SecurityTypesRequest)
	}

	// same account can not used for creating Security Types Request and Security Types Request Reject with the same SecurityReqID
	if securityTypes.SecurityTypesRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Security Types Request Reject", msg.Creator))
	}

	// check that the Security Types Request is not rejected already
	if securityTypes.SecurityTypesRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesRequestIsRejected, "Security Types: %s", securityTypes.SecurityTypesRequestReject)
	}

	// check that the Security Types Request is not acknowledged already
	if securityTypes.SecurityTypesResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesRequestIsAcknowledged, "Security Types: %s", securityTypes.SecurityTypesResponse)
	}

	// check that the mandatory Security Types Request Reject field is not empty
	if msg.RejectReason == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypesEmptyField, "RejectReason: %s", msg.RejectReason)
	}

	securityTypesRequestReject := types.SecurityTypes{
		SessionID:             msg.SessionID,
		SecurityTypesRequest:  securityTypes.SecurityTypesRequest,
		SecurityTypesResponse: securityTypes.SecurityTypesResponse,
		SecurityTypesRequestReject: &types.SecurityTypesRequestReject{
			SecurityReqID: msg.SecurityReqID,
			RejectReason:  msg.RejectReason,
			Text:          msg.Text,
		},
	}

	// set header from the existing security types request
	securityTypesRequestReject.SecurityTypesRequestReject.Header = securityTypes.SecurityTypesRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *securityTypesRequestReject.SecurityTypesRequestReject.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = securityTypes.SecurityTypesRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = securityTypes.SecurityTypesRequest.Header.MsgSeqNum

	// set the msgType to Security Types Request Reject
	newHeader.MsgType = "y"

	// switch senderCompID and targetCompID between Security Types Request and Security Types Request Reject
	// set senderCompID of Security Types Request Reject to the targetCompID of Security Types Request in the header
	newHeader.SenderCompID = securityTypes.SecurityTypesRequest.Header.TargetCompID

	// set targetCompID of Security Types Request Reject to the senderCompID of Security Types Request in the header
	newHeader.TargetCompID = securityTypes.SecurityTypesRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	securityTypesRequestReject.SecurityTypesRequestReject.Header = newHeader

	// set Trailer from the existing Security Types Request
	securityTypesRequestReject.SecurityTypesRequestReject.Trailer = securityTypes.SecurityTypesRequest.Trailer

	// set Security Types Request to store
	k.SetSecurityTypes(ctx, msg.SecurityReqID, securityTypesRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityTypesRequestRejectResponse{}, err
}
