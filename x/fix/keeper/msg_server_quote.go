package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// QuoteRequest creates Quote Request 
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
	mic, found := k.micKeeper.GetMarketIdentificationCode(ctx, msg.QuoteRequest.Mic)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMICIsNotFound, "MIC: %s", msg.QuoteRequest.Mic)
	}

	//check that the address creating the Quote Request is same addresss used to register the MIC on the mic module
	if mic.MIC != msg.QuoteRequest.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "MIC Creator: %s", msg.QuoteRequest.Creator)
	}

	//call new instance of QuoteRequest
	quoteRequests := types.NewQuoteRequest(msg.QuoteRequest.QuoteReqID, msg.QuoteRequest.Symbol, msg.QuoteRequest.SecurityID, msg.QuoteRequest.SecurityIDSource, msg.QuoteRequest.Side, msg.QuoteRequest.OrderQty, msg.QuoteRequest.FutSettDate, msg.QuoteRequest.SettlDate2, msg.QuoteRequest.Account, msg.QuoteRequest.BidPx, msg.QuoteRequest.OfferPx, msg.QuoteRequest.Currency, msg.QuoteRequest.ValidUntilTime, msg.QuoteRequest.ExpireTime, msg.QuoteRequest.QuoteType, msg.QuoteRequest.BidSize, msg.QuoteRequest.OfferSize, msg.QuoteRequest.Mic, msg.QuoteRequest.Text, msg.QuoteRequest.Creator)

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

	//set the header and make changes to the header
	//calculate and include all changes to the header
	//Message type, R is the message type for Quote Request
	//BodyLength should be calculated using the BodyLength function
	newQuoteRequest.QuoteRequest.Header = header
	newQuoteRequest.QuoteRequest.Header.MsgType = "R"

	//fetch Trailer from existing session
	var trailer *types.Trailer
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	//set the Trailer and make changes to the trailer
	//checksum in the trailer can be recalculated using CalculateChecksum function
	newQuoteRequest.QuoteRequest.Trailer = trailer

	//set Quote to store
	k.SetQuote(ctx, msg.QuoteRequest.QuoteReqID, newQuoteRequest)

	return &types.MsgQuoteRequestResponse{}, nil
}


// QuoteAcknowledgement creates Quote Acknowledgement
func (k msgServer) QuoteAcknowledgement(goCtx context.Context, msg *types.MsgQuoteAcknowledgement) (*types.MsgQuoteAcknowledgementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgQuoteAcknowledgementResponse{}, nil
}
