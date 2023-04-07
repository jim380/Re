package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// QuoteRequest creates Quote Request
func (k msgServer) QuoteRequest(goCtx context.Context, msg *types.MsgQuoteRequest) (*types.MsgQuoteRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != "loggedIn" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Quote Request
	if session.InitiatorAddress != msg.QuoteRequest.Creator && session.AcceptorAddress != msg.QuoteRequest.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.QuoteRequest.Creator)
	}

	// check that mandatory Quote Request fields are not empty
	if msg.QuoteRequest.QuoteReqID == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteReqIDIsEmpty, "QuoteReqID: %s", msg.QuoteRequest.QuoteReqID)
	}
	if msg.QuoteRequest.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteSymbolIsEmpty, "Symbol: %s", msg.QuoteRequest.Symbol)
	}
	if msg.QuoteRequest.Side == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteSideIsEmpty, "Side: %s", msg.QuoteRequest.Side)
	}
	if msg.QuoteRequest.BidPx == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteBidPxIsEmpty, "BidPx: %s", msg.QuoteRequest.BidPx)
	}
	if msg.QuoteRequest.Currency == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteCurrencyIsEmpty, "Currency: %s", msg.QuoteRequest.Currency)
	}
	if msg.QuoteRequest.QuoteType == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteTypeIsEmpty, "QuoteType: %s", msg.QuoteRequest.QuoteType)
	}

	// get market identification code from MIC module
	mic, found := k.micKeeper.GetMarketIdentificationCode(ctx, msg.QuoteRequest.Mic)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMICInQuoteRquestIsNotFound, "MIC: %s", msg.QuoteRequest.Mic)
	}

	// check that the address creating the Quote Request is same addresss used to register the MIC on the mic module
	if mic.Creator != msg.QuoteRequest.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "MIC Creator: %s", msg.QuoteRequest.Creator)
	}

	// call new instance of NewQuoteRequest
	quoteRequest := types.NewQuoteRequest(msg.QuoteRequest.QuoteReqID, msg.QuoteRequest.Symbol, msg.QuoteRequest.SecurityID, msg.QuoteRequest.SecurityIDSource, msg.QuoteRequest.Side, msg.QuoteRequest.OrderQty, msg.QuoteRequest.FutSettDate, msg.QuoteRequest.SettlDate2, msg.QuoteRequest.Account, msg.QuoteRequest.BidPx, msg.QuoteRequest.OfferPx, msg.QuoteRequest.Currency, msg.QuoteRequest.ValidUntilTime, msg.QuoteRequest.ExpireTime, msg.QuoteRequest.QuoteType, msg.QuoteRequest.BidSize, msg.QuoteRequest.OfferSize, msg.QuoteRequest.Mic, msg.QuoteRequest.Text, msg.QuoteRequest.Creator)

	newQuoteRequest := types.Quote{
		SessionID:    msg.SessionID,
		QuoteRequest: quoteRequest,
	}

	// fetch Header from existing session
	// In the FIX Protocol, a Quote Request message can be sent by either the initiator or the acceptor of the FIX session.
	//Determine whether we are the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, R is the message type for Quote Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Quote Request
	newQuoteRequest.QuoteRequest.Header = header
	newQuoteRequest.QuoteRequest.Header.MsgType = "R"
	newQuoteRequest.QuoteRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	var trailer *types.Trailer
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// checksum in the trailer can be recalculated using CalculateChecksum function
	newQuoteRequest.QuoteRequest.Trailer = trailer

	//set Quote Request to store
	k.SetQuote(ctx, msg.QuoteRequest.QuoteReqID, newQuoteRequest)

	// emit event
	err := ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgQuoteRequestResponse{}, err
}

// QuoteAcknowledgement creates Quote Acknowledgement for every Quote Request
func (k msgServer) QuoteAcknowledgement(goCtx context.Context, msg *types.MsgQuoteAcknowledgement) (*types.MsgQuoteAcknowledgementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Quote Acknowledgement matches with the sessionID for Quote Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrWrongSessionIDInQuote, "SessionID: %s", msg.SessionID)
	}

	// check that the user to acknowledge the Quote Request is the recipeint of the Quote Request
	if session.InitiatorAddress != msg.QuoteAcknowledgement.Creator && session.AcceptorAddress != msg.QuoteAcknowledgement.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Quote Acknowledgement Creator: %s", msg.QuoteAcknowledgement.Creator)
	}

	// get the existing Quote Request
	quoteRequest, found := k.GetQuote(ctx, msg.QuoteAcknowledgement.QuoteReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrQuoteIsEmpty, "Quote: %s", &quoteRequest)
	}

	// check that Quote Request creator address is not same address acknowledging the Quote Request
	if quoteRequest.QuoteRequest.Creator == msg.QuoteAcknowledgement.Creator {
		return nil, sdkerrors.Wrapf(types.ErrQuoteAcknowledgementCreatorIsWrong, "Quote Acknowledgement: %s", msg.QuoteAcknowledgement.Creator)
	}

	// check that this Quote Request to be acknowledged has not been rejected
	// access QuoteRequestReject from QuoteRequest, Quote Acknoledgement should be rejected if Quote Request Reject is not nil
	if quoteRequest.QuoteRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrQuoteRequestIsRejected, "Quote: %s", quoteRequest.QuoteRequestReject)
	}

	// check that mandatory Quote Acknowledgement fields are not empty
	if msg.QuoteAcknowledgement.QuoteReqID == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteReqIDIsEmpty, "QuoteReqID: %s", msg.QuoteAcknowledgement.QuoteReqID)
	}
	if msg.QuoteAcknowledgement.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteSymbolIsEmpty, "Symbol: %s", msg.QuoteAcknowledgement.Symbol)
	}
	if msg.QuoteAcknowledgement.Side == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteSideIsEmpty, "Side: %s", msg.QuoteAcknowledgement.Side)
	}
	if msg.QuoteAcknowledgement.BidPx == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteBidPxIsEmpty, "BidPx: %s", msg.QuoteAcknowledgement.BidPx)
	}
	if msg.QuoteAcknowledgement.Currency == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteCurrencyIsEmpty, "Currency: %s", msg.QuoteAcknowledgement.Currency)
	}
	if msg.QuoteAcknowledgement.QuoteType == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteTypeIsEmpty, "QuoteType: %s", msg.QuoteAcknowledgement.QuoteType)
	}

	// call new instance of NewQuoteAcknowledgement
	quoteAcknowledgement := types.NewQuoteAcknowledgement(msg.QuoteAcknowledgement.QuoteReqID, msg.QuoteAcknowledgement.QuoteID, msg.QuoteAcknowledgement.QuoteStatus, msg.QuoteAcknowledgement.QuoteType, msg.QuoteAcknowledgement.SecurityID, msg.QuoteAcknowledgement.SecurityIDSource, msg.QuoteAcknowledgement.Symbol, msg.QuoteAcknowledgement.Side, msg.QuoteAcknowledgement.OrderQty, msg.QuoteAcknowledgement.LastQty, msg.QuoteAcknowledgement.LastPx, msg.QuoteAcknowledgement.BidPx, msg.QuoteAcknowledgement.OfferPx, msg.QuoteAcknowledgement.Currency, msg.QuoteAcknowledgement.SettlDate, msg.QuoteAcknowledgement.ValidUntilTime, msg.QuoteAcknowledgement.ExpireTime, msg.QuoteAcknowledgement.Text, msg.QuoteAcknowledgement.NoQuoteQualifiers, msg.QuoteAcknowledgement.QuoteQualifier, msg.QuoteAcknowledgement.NoLegs, msg.QuoteAcknowledgement.LegSymbol, msg.QuoteAcknowledgement.LegSecurityID, msg.QuoteAcknowledgement.LegSecurityIDSource, msg.QuoteAcknowledgement.LegRatioQty, msg.QuoteAcknowledgement.Creator)

	newQuoteAcknowledgement := types.Quote{
		SessionID:            msg.SessionID,
		QuoteRequest:         quoteRequest.QuoteRequest,
		QuoteAcknowledgement: quoteAcknowledgement,
	}

	// set header from the existing quote request
	newQuoteAcknowledgement.QuoteAcknowledgement.Header = quoteRequest.QuoteRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *newQuoteAcknowledgement.QuoteAcknowledgement.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = quoteRequest.QuoteRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = quoteRequest.QuoteRequest.Header.MsgSeqNum

	// set the msgType to quote acknowledgement
	newHeader.MsgType = "b"

	// switch senderCompID and targetCompID between Quote Request and Quote Acknoledgement
	// set senderCompID of Quote Acknowledgement to the targetCompID of Quote Request in the header
	newHeader.SenderCompID = quoteRequest.QuoteRequest.Header.TargetCompID

	// set targetCompID of Quote Acknowledgement to the senderCompID of Quote Request in the header
	newHeader.TargetCompID = quoteRequest.QuoteRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	newQuoteAcknowledgement.QuoteAcknowledgement.Header = newHeader

	// set Trailer from the existing Quote Request
	newQuoteAcknowledgement.QuoteAcknowledgement.Trailer = quoteRequest.QuoteRequest.Trailer

	// set Quote Acknowledgement to store
	k.SetQuote(ctx, msg.QuoteAcknowledgement.QuoteReqID, newQuoteAcknowledgement)

	// emit event
	err := ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgQuoteAcknowledgementResponse{}, err
}

// QuoteRequestReject creates Quote Request Reject for Quote Request
func (k msgServer) QuoteRequestReject(goCtx context.Context, msg *types.MsgQuoteRequestReject) (*types.MsgQuoteRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Quote Request Reject matches with the sessionID for Quote Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrWrongSessionIDInQuote, "SessionID: %s", msg.SessionID)
	}

	// check that the user to Reject the Quote Request is the recipeint of the Quote Request
	if session.InitiatorAddress != msg.QuoteRequestReject.Creator && session.AcceptorAddress != msg.QuoteRequestReject.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Quote Acknowledgement Creator: %s", msg.QuoteRequestReject.Creator)
	}

	// get the existing Quote Request
	quoteRequest, found := k.GetQuote(ctx, msg.QuoteRequestReject.QuoteReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrQuoteIsEmpty, "Quote: %s", &quoteRequest)
	}

	// check that Quote Request creator address is not same address Rejecting the Quote Request
	if quoteRequest.QuoteRequest.Creator == msg.QuoteRequestReject.Creator {
		return nil, sdkerrors.Wrapf(types.ErrQuoteRequestRejectCreatorIsWrong, "Quote Request Rejection: %s", msg.QuoteRequestReject.Creator)
	}

	// check that this Quote Request to be Rejected has not been acknowledged
	// access QuoteAcknowledgement from the QuoteRequest, QuoteRequestRejection should be rejected if QuoteAcknowledgement is not nil
	if quoteRequest.QuoteAcknowledgement != nil {
		return nil, sdkerrors.Wrapf(types.ErrQuoteRequestIsAcknowledged, "Quote: %s", quoteRequest.QuoteAcknowledgement)
	}

	// check that mandatory Quote Request Reject fields are not empty
	if msg.QuoteRequestReject.QuoteReqID == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteReqIDIsEmpty, "QuoteReqID: %s", msg.QuoteRequestReject.QuoteReqID)
	}
	if msg.QuoteRequestReject.QuoteRequestRejectReason == "" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteRequestRejectReasonIsEmpty, "QuoteRequestRejectReason: %s", msg.QuoteRequestReject.QuoteRequestRejectReason)
	}

	// call new instance of NewQuoteRequestReject
	quoteRequestReject := types.NewQuoteRequestReject(msg.QuoteRequestReject.QuoteReqID, msg.QuoteRequestReject.QuoteRequestRejectReason, msg.QuoteRequestReject.Text, msg.QuoteRequestReject.Creator)

	newQuoteRequestReject := types.Quote{
		SessionID:            msg.SessionID,
		QuoteRequest:         quoteRequest.QuoteRequest,
		QuoteAcknowledgement: quoteRequest.QuoteAcknowledgement,
		QuoteRequestReject:   quoteRequestReject,
	}

	// set header from the existing quote request
	newQuoteRequestReject.QuoteRequestReject.Header = quoteRequest.QuoteRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *newQuoteRequestReject.QuoteRequestReject.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = quoteRequest.QuoteRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = quoteRequest.QuoteRequest.Header.MsgSeqNum

	// set the msgType to Quote Request Reject
	newHeader.MsgType = "AI"

	// switch senderCompID and targetCompID between Quote Request and Quote Request Reject
	// set senderCompID of Quote Request Reject to the targetCompID of Quote Request in the header
	newHeader.SenderCompID = quoteRequest.QuoteRequest.Header.TargetCompID

	// set targetCompID of Quote Request Reject to the senderCompID of Quote Request in the header
	newHeader.TargetCompID = quoteRequest.QuoteRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	newQuoteRequestReject.QuoteRequestReject.Header = newHeader

	// set Trailer from the existing Quote Request
	newQuoteRequestReject.QuoteRequestReject.Trailer = quoteRequest.QuoteRequest.Trailer

	// set Quote Request Reject to store
	k.SetQuote(ctx, msg.QuoteRequestReject.QuoteReqID, newQuoteRequestReject)

	// emit event
	err := ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgQuoteRequestRejectResponse{}, err
}
