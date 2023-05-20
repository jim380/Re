package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityDefinitionRequest(goCtx context.Context, msg *types.MsgSecurityDefinitionRequest) (*types.MsgSecurityDefinitionRequestResponse, error) {
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
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Trade Capture Report
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that mandatory Security Definition Request fields are not empty
	if msg.SecurityReqID == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityReqIDIsEmpty, "SecurityReqID: %s", msg.SecurityReqID)
	}
	if msg.SecurityRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityRequestTypeIsEmpty, "SecurityRequestType: %s", msg.SecurityRequestType)
	}
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecuritySymbolIsEmpty, "Symbol: %s", msg.Symbol)
	}
	if msg.SecurityExchange == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityExchangeIsEmpty, "SecurityExchange: %s", msg.SecurityExchange)
	}

	// security definition request
	securityDefinitionRequest := types.Security{
		SessionID: msg.SessionID,
		SecurityDefinitionRequest: &types.SecurityDefinitionRequest{
			SecurityReqID:       msg.SecurityReqID,
			SecurityRequestType: msg.SecurityRequestType,
			Symbol:              msg.Symbol,
			SecurityExchange:    msg.SecurityExchange,
			Issuer:              msg.Issuer,
			SecurityDesc:        msg.SecurityDesc,
			SecurityType:        msg.SecurityType,
			Currency:            msg.Currency,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Security Definition Request message can be sent by either the initiator or the acceptor of the FIX session.
	// Determine whether it is the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, d is the message type for Security Definition Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Security Definition Request
	securityDefinitionRequest.SecurityDefinitionRequest.Header = header
	securityDefinitionRequest.SecurityDefinitionRequest.Header.MsgType = "d"
	securityDefinitionRequest.SecurityDefinitionRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	// for now copy trailer from session, it should be re-calculated
	var trailer *types.Trailer
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	securityDefinitionRequest.SecurityDefinitionRequest.Trailer = trailer

	// set new security definition request to store
	k.SetSecurity(ctx, msg.SecurityReqID, securityDefinitionRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityDefinitionRequestResponse{}, err
}

func (k msgServer) SecurityDefinition(goCtx context.Context, msg *types.MsgSecurityDefinition) (*types.MsgSecurityDefinitionResponse, error) {
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

	// check that the sessionID provided by the creator of Security Definition matches with the sessionID for Security Definition Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrWrongSessionIDInSecurity, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Security Definition Request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security Definition Creator: %s", msg.Creator)
	}

	// get Security Definition Request
	security, found := k.GetSecurity(ctx, msg.SecurityReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityIsNotFound, ": %s", security.SecurityDefinition)
	}

	// check that the Security Definition Request is not rejected already
	if security.SecurityDefinitionRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityDefinitionRequestIsRejected, "Security: %s", security.SecurityDefinitionRequestReject)
	}

	// check that the Security Definition Request is not acknowledged already
	if security.SecurityDefinition != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityDefinitionRequestIsAcknowledged, "Security: %s", security.SecurityDefinition)
	}

	// check that the mandatory field matches the value from Security Definition Request
	if security.SecurityDefinitionRequest.SecurityReqID != msg.SecurityReqID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityMismatchField, "SecurityReqID: %s", msg.SecurityReqID)
	}

	// check that the mandatory Security Definition fields are not empty
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecuritySymbolIsEmpty, "Symbol: %s", msg.Symbol)
	}
	if msg.SecurityExchange == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityExchangeIsEmpty, "SecurityExchange: %s", msg.SecurityExchange)
	}
	if msg.SecurityType == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityTypeIsEmpty, "SecurityType: %s", msg.SecurityType)
	}
	if msg.Currency == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityCurrencyIsEmpty, "Currency: %s", msg.Currency)
	}
	if msg.MinPriceIncrement == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityMinPriceIncrementIsEmpty, "MinPriceIncrement: %s", msg.MinPriceIncrement)
	}

	// security definition
	securityDefinition := types.Security{
		SessionID:                 msg.SessionID,
		SecurityDefinitionRequest: security.SecurityDefinitionRequest,
		SecurityDefinition: &types.SecurityDefinition{
			SecurityReqID:           msg.SecurityReqID,
			SecurityResponseID:      msg.SecurityResponseID,
			SecurityResponseType:    msg.SecurityResponseType,
			Symbol:                  msg.Symbol,
			SecurityExchange:        msg.SecurityExchange,
			Issuer:                  msg.Issuer,
			SecurityDesc:            msg.SecurityDesc,
			SecurityType:            msg.SecurityType,
			Currency:                msg.Currency,
			ContractMultiplier:      msg.ContractMultiplier,
			MinPriceIncrement:       msg.MinPriceIncrement,
			MinPriceIncrementAmount: msg.MinPriceIncrementAmount,
			UnitOfMeasure:           msg.UnitOfMeasure,
			PriceUnitOfMeasure:      msg.PriceUnitOfMeasure,
			SettlType:               msg.SettlType,
			SettlDate:               msg.SettlDate,
			MaturityMonthYear:       msg.MaturityMonthYear,
			CouponRate:              msg.CouponRate,
			Factor:                  msg.Factor,
			CreditRating:            msg.CreditRating,
			SecurityExchangeID:      msg.SecurityExchangeID,
		},
	}

	// set header from Security Definition Request
	securityDefinition.SecurityDefinition.Header = security.SecurityDefinitionRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *securityDefinition.SecurityDefinition.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = security.SecurityDefinitionRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = security.SecurityDefinitionRequest.Header.MsgSeqNum

	// set the msgType to Security Definition
	newHeader.MsgType = "d"

	// switch senderCompID and targetCompID between Security Definition Request and Security Definition
	// set senderCompID of Security Definition to the targetCompID of Security Definition Request in the header
	newHeader.SenderCompID = security.SecurityDefinitionRequest.Header.TargetCompID

	// set targetCompID of Security Definition to the senderCompID of Security Definition Request in the header
	newHeader.TargetCompID = security.SecurityDefinitionRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	securityDefinition.SecurityDefinition.Header = newHeader

	// set Trailer from the existing Security Definition Request
	// TODO
	// checksum should be recalcualted
	securityDefinition.SecurityDefinition.Trailer = security.SecurityDefinitionRequest.Trailer

	// set new security definition to store
	k.SetSecurity(ctx, msg.SecurityReqID, securityDefinition)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityDefinitionResponse{}, err
}

func (k msgServer) SecurityDefinitionRequestReject(goCtx context.Context, msg *types.MsgSecurityDefinitionRequestReject) (*types.MsgSecurityDefinitionRequestRejectResponse, error) {
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

	// check that the sessionID provided by the creator of Security Definition Request Reject matches with the sessionID for Security Definition Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrWrongSessionIDInSecurity, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Security Definition Request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Security Definition Request Reject: %s", msg.Creator)
	}

	// get Security Definition Request
	security, found := k.GetSecurity(ctx, msg.SecurityReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSecurityIsNotFound, ": %s", security.SecurityDefinition)
	}

	// check that the Security Definition Request is not rejected already
	if security.SecurityDefinitionRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityDefinitionRequestIsRejected, "Security: %s", security.SecurityDefinitionRequestReject)
	}

	// check that the Security Definition Request is not acknowledged already
	if security.SecurityDefinition != nil {
		return nil, sdkerrors.Wrapf(types.ErrSecurityDefinitionRequestIsAcknowledged, "Security: %s", security.SecurityDefinition)
	}

	// check that the mandatory field matches the value from Security Definition Request
	if security.SecurityDefinitionRequest.SecurityReqID != msg.SecurityReqID {
		return nil, sdkerrors.Wrapf(types.ErrSecurityMismatchField, "SecurityReqID: %s", msg.SecurityReqID)
	}

	// check that the mandatory Security Definition fields are not empty
	if msg.SecurityRequestResult == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityRequestResultIsEmpty, "SecurityRequestResult: %s", msg.SecurityRequestResult)
	}
	if msg.SecurityRequestError == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecuritySymbolIsEmpty, "SecurityRequestError: %s", msg.SecurityRequestError)
	}

	// security definition request reject
	securityDefinitionRequestReject := types.Security{
		SessionID:                 msg.SessionID,
		SecurityDefinitionRequest: security.SecurityDefinitionRequest,
		SecurityDefinition:        security.SecurityDefinition,
		SecurityDefinitionRequestReject: &types.SecurityDefinitionRequestReject{
			SecurityReqID:            msg.SecurityReqID,
			SecurityRequestResult:    msg.SecurityRequestResult,
			SecurityRequestError:     msg.SecurityRequestError,
			SecurityRequestErrorCode: msg.SecurityRequestErrorCode,
			Text:                     msg.Text,
		},
	}

	// set Security Definition Request Reject to store
	k.SetSecurity(ctx, msg.SecurityReqID, securityDefinitionRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityDefinitionRequestRejectResponse{}, err
}
