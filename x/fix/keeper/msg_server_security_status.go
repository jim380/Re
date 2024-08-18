package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/utils/helpers"
	"github.com/jim380/Re/x/fix/types"
)

// SecurityStatusRequest creates Security Status Request
func (k msgServer) SecurityStatusRequest(goCtx context.Context, msg *types.MsgSecurityStatusRequest) (*types.MsgSecurityStatusRequestResponse, error) {
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

	// check that the parties involved in a session are the ones using the sessionID and are able to create Security Status Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that the mandatory Security Status Request field is not empty
	if msg.SubscriptionRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusEmptyField, "SubscriptionRequestType: %s", msg.SubscriptionRequestType)
	}

	// Must be unique, or the ID of previous Security Status Request (e) to disable if SubscriptionRequestType (263) = Disable previous Snapshot + Updates Request (2)
	if msg.SecurityStatusReqID == "" {
		// Generate a random TradSesReqID using helpers.GenerateRandomString
		securityStatusReqID, _ := helpers.GenerateRandomString(constants.TradSesReqID)
		msg.SecurityStatusReqID = securityStatusReqID
	} else {
		// check that the provided SecurityStatusReqID exists from the Security Status
		securityStatus, found := k.GetSecurityStatus(ctx, msg.SecurityStatusReqID)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrSecurityStatusIsNotFound, ": %s", securityStatus.SecurityStatusRequest)
		}
		msg.SecurityStatusReqID = securityStatus.SecurityStatusRequest.SecurityStatusReqID
	}

	securityStatusRequest := types.SecurityStatus{
		SessionID: msg.SessionID,
		SecurityStatusRequest: &types.SecurityStatusRequest{
			SecurityStatusReqID:     msg.SecurityStatusReqID,
			Instrument:              msg.Instrument,
			NoUnderlyings:           msg.NoUnderlyings,
			UnderlyingInstrument:    msg.UnderlyingInstrument,
			NoLegs:                  msg.NoLegs,
			InstrumentLeg:           msg.InstrumentLeg,
			SubscriptionRequestType: msg.SubscriptionRequestType,
			TradingSessionID:        msg.TradingSessionID,
			TradingSessionSubID:     msg.TradingSessionSubID,
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
	// Message type, e is the message type for Security Status Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Security Status Request
	securityStatusRequest.SecurityStatusRequest.Header = header
	securityStatusRequest.SecurityStatusRequest.Header.MsgType = "x"
	securityStatusRequest.SecurityStatusRequest.Header.SendingTime = constants.SendingTime

	// fetch Trailer from existing session
	var trailer *types.Trailer
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// checksum in the trailer can be recalculated using CalculateChecksum function
	securityStatusRequest.SecurityStatusRequest.Trailer = trailer

	// set Security Status Request to store
	k.SetSecurityStatus(ctx, msg.SecurityStatusReqID, securityStatusRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityStatusRequestResponse{}, err
}

// SecurityStatusResponse creates Security Status Response
func (k msgServer) SecurityStatusResponse(goCtx context.Context, msg *types.MsgSecurityStatusResponse) (*types.MsgSecurityStatusResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Security Status Response matches with the sessionID for Security Status Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Security Status Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security Status Response Creator: %s", msg.Creator)
	}

	// get security status
	securityStatus, found := k.GetSecurityStatus(ctx, msg.SecurityStatusReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusIsNotFound, ": %s", securityStatus.SecurityStatusRequest)
	}

	// same account can not used for creating Security Status Request and Security Status Response with the same SecurityStatusReqID
	if securityStatus.SecurityStatusRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Security Status Response", msg.Creator))
	}

	// check that the Security Status Request is not rejected already
	if securityStatus.SecurityStatusRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusRequestIsRejected, "Security Status: %s", securityStatus.SecurityStatusRequestReject)
	}

	// check that the Security Status Request is not acknowledged already
	if securityStatus.SecurityStatusResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusRequestIsAcknowledged, "Security Status: %s", securityStatus.SecurityStatusResponse)
	}

	securityStatusResponse := types.SecurityStatus{
		SessionID:             msg.SessionID,
		SecurityStatusRequest: securityStatus.SecurityStatusRequest,
		SecurityStatusResponse: &types.SecurityStatusResponse{
			SecurityStatusReqID:   msg.SecurityStatusReqID,
			NoUnderlyings:         msg.NoUnderlyings,
			UnderlyingInstrument:  msg.UnderlyingInstrument,
			NoLegs:                msg.NoLegs,
			InstrumentLeg:         msg.InstrumentLeg,
			Currency:              msg.Currency,
			TradingSessionID:      msg.TradingSessionID,
			TradingSessionSubID:   msg.TradingSessionSubID,
			UnsolicitedIndicator:  msg.UnsolicitedIndicator,
			SecurityTradingStatus: msg.SecurityTradingStatus,
			FinancialStatus:       msg.FinancialStatus,
			CorporateAction:       msg.CorporateAction,
			HaltReason:            msg.HaltReason,
			InViewOfCommon:        msg.InViewOfCommon,
			DueToRelated:          msg.DueToRelated,
			BuyVolume:             msg.BuyVolume,
			SellVolume:            msg.SellVolume,
			HighPx:                msg.HighPx,
			LowPx:                 msg.LowPx,
			LastPx:                msg.LastPx,
			TransactTime:          msg.TransactTime,
			Adjustment:            msg.Adjustment,
			Text:                  msg.Text,
		},
	}

	// set header from the existing security status request
	securityStatusResponse.SecurityStatusResponse.Header = securityStatus.SecurityStatusRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *securityStatusResponse.SecurityStatusResponse.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = securityStatus.SecurityStatusRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = securityStatus.SecurityStatusRequest.Header.MsgSeqNum

	// set the msgType to Security Status Response
	newHeader.MsgType = "f"

	// switch senderCompID and targetCompID between Security Status Request and Security Status Response
	// set senderCompID of Security Status Response to the targetCompID of Security Status Request in the header
	newHeader.SenderCompID = securityStatus.SecurityStatusRequest.Header.TargetCompID

	// set targetCompID of Security Status Response to the senderCompID of Security Status Request in the header
	newHeader.TargetCompID = securityStatus.SecurityStatusRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	securityStatusResponse.SecurityStatusResponse.Header = newHeader

	// set Trailer from the existing Security Status Request
	securityStatusResponse.SecurityStatusResponse.Trailer = securityStatus.SecurityStatusRequest.Trailer

	// set Security Status Response to store
	k.SetSecurityStatus(ctx, msg.SecurityStatusReqID, securityStatusResponse)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityStatusResponseResponse{}, err
}

// SecurityStatusRequestReject creates Security Status Request Reject
func (k msgServer) SecurityStatusRequestReject(goCtx context.Context, msg *types.MsgSecurityStatusRequestReject) (*types.MsgSecurityStatusRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Security Status Request Reject matches with the sessionID for Security Status Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Security Status Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security Status Response Creator: %s", msg.Creator)
	}

	// get security status
	securityStatus, found := k.GetSecurityStatus(ctx, msg.SecurityStatusReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusIsNotFound, ": %s", securityStatus.SecurityStatusRequest)
	}

	// same account can not used for creating Security Status Request and Security Status Request Reject with the same SecurityStatusReqID
	if securityStatus.SecurityStatusRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Security Status Response Reject", msg.Creator))
	}

	// check that the Security Status Request is not rejected already
	if securityStatus.SecurityStatusRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusRequestIsRejected, "Security Status: %s", securityStatus.SecurityStatusRequestReject)
	}

	// check that the Security Status Request is not acknowledged already
	if securityStatus.SecurityStatusResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusRequestIsAcknowledged, "Security Status: %s", securityStatus.SecurityStatusResponse)
	}

	// check that the mandatory Security Status Request Reject field is not empty
	if msg.SecurityRejectReason == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusEmptyField, "SecurityRejectReason: %s", msg.SecurityRejectReason)
	}

	securityStatusRequestReject := types.SecurityStatus{
		SessionID:              msg.SessionID,
		SecurityStatusRequest:  securityStatus.SecurityStatusRequest,
		SecurityStatusResponse: securityStatus.SecurityStatusResponse,
		SecurityStatusRequestReject: &types.SecurityStatusRequestReject{
			SecurityStatusReqID:  msg.SecurityStatusReqID,
			SecurityRejectReason: msg.SecurityRejectReason,
			Text:                 msg.Text,
		},
	}

	// set header from the existing security status request
	securityStatusRequestReject.SecurityStatusRequestReject.Header = securityStatus.SecurityStatusRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *securityStatusRequestReject.SecurityStatusRequestReject.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = securityStatus.SecurityStatusRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = securityStatus.SecurityStatusRequest.Header.MsgSeqNum

	// set the msgType to Security Status Request Reject
	newHeader.MsgType = "j"

	// switch senderCompID and targetCompID between Security Status Request and Security Status Request Reject
	// set senderCompID of Security Status Request Reject to the targetCompID of Security Status Request in the header
	newHeader.SenderCompID = securityStatus.SecurityStatusRequest.Header.TargetCompID

	// set targetCompID of Security Status Request Reject to the senderCompID of Security Status Request in the header
	newHeader.TargetCompID = securityStatus.SecurityStatusRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	securityStatusRequestReject.SecurityStatusRequestReject.Header = newHeader

	// set Trailer from the existing Security Status Request
	securityStatusRequestReject.SecurityStatusRequestReject.Trailer = securityStatus.SecurityStatusRequest.Trailer

	// set Security Status Request Reject to store
	k.SetSecurityStatus(ctx, msg.SecurityStatusReqID, securityStatusRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityStatusRequestRejectResponse{}, err
}
