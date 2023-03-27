package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) QuoteRequest(goCtx context.Context, msg *types.MsgQuoteRequest) (*types.MsgQuoteRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	//check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != "loggedIn" {
		return nil, sdkerrors.Wrapf(types.ErrQuoteSession, "Status of Session: %s", msg.SessionID)
	}

	//fetch existing Quote Request and use it to check for multiple Quote Requests
	//same quote with quoteReqID is not allowed
	quoteRequest, _ := k.GetQuote(ctx, msg.QuoteRequest.QuoteReqID)
	if quoteRequest.QuoteRequest.QuoteReqID == msg.QuoteRequest.QuoteReqID {
		return nil, sdkerrors.Wrapf(types.ErrQuoteReqIDIsTaken, "QuoteReqID: %s", msg.QuoteRequest.QuoteReqID)
	}

	//check that mandatory Quote Request fields are not empty
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

	//get market identification code from MIC module

	// call new instance of QuoteRequest
	quoteRequests := types.NewQuoteRequest(msg.QuoteRequest.Symbol, msg.QuoteRequest.SecurityID, msg.QuoteRequest.SecurityIDSource, msg.QuoteRequest.Side, msg.QuoteRequest.OrderQty, msg.QuoteRequest.FutSettDate, msg.QuoteRequest.SettlDate2, msg.QuoteRequest.Account, msg.QuoteRequest.BidPx, msg.QuoteRequest.OfferPx, msg.QuoteRequest.Currency, msg.QuoteRequest.ValidUntilTime, msg.QuoteRequest.ExpireTime, msg.QuoteRequest.QuoteType, msg.QuoteRequest.BidSize, msg.QuoteRequest.OfferSize, msg.QuoteRequest.Mic, msg.QuoteRequest.Text, msg.QuoteRequest.Creator)

	newQuoteRequest := types.Quote{
		SessionID:    msg.SessionID,
		QuoteRequest: quoteRequests,
	}

	//fetch Header from existing session
	//In the FIX Protocol, a Quote Request message can be sent by either the initiator or the acceptor of the FIX session.
	//Determine whether we are the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	//fetch Trailer from existing session
	var trailer *types.Trailer
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	//set the header
	newQuoteRequest.QuoteRequest.Header = header

	//set the Trailer
	newQuoteRequest.QuoteRequest.Trailer = trailer

	//set QuoteReqID using GenerateRandomString Function
	quoteReqID, _ := types.GenerateRandomString(types.QuoteReqIDLength)
	newQuoteRequest.QuoteRequest.QuoteReqID = quoteReqID

	//set Quote to store
	k.SetQuote(ctx, msg.QuoteRequest.QuoteReqID, newQuoteRequest)

	return &types.MsgQuoteRequestResponse{}, nil
}
