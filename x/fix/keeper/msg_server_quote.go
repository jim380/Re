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
	//mic, found := k.micKeeper.GetMarketIdentificationCode(ctx, msg.QuoteRequest.Mic)
	//if !found {
	//	return nil, sdkerrors.Wrapf(types.ErrMICIsNotFound, "MIC: %s", msg.QuoteRequest.Mic)
	//}

	//check that the address creating the Quote Request is same addresss used to register the MIC on the mic module
	//if mic.MIC != msg.QuoteRequest.Creator {
	//	return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "MIC Creator: %s", msg.QuoteRequest.Creator)
	//}

	//set quote request data
	quoteRequests := types.QuoteRequest{
		Header:           msg.QuoteRequest.Header,
		QuoteReqID:       msg.QuoteRequest.QuoteReqID,
		Symbol:           msg.QuoteRequest.Symbol,
		SecurityID:       msg.QuoteRequest.SecurityID,
		SecurityIDSource: msg.QuoteRequest.SecurityIDSource,
		Side:             msg.QuoteRequest.Side,
		OrderQty:         msg.QuoteRequest.OrderQty,
		FutSettDate:      msg.QuoteRequest.FutSettDate,
		SettlDate2:       msg.QuoteRequest.SettlDate2,
		Account:          msg.QuoteRequest.Account,
		BidPx:            msg.QuoteRequest.BidPx,
		OfferPx:          msg.QuoteRequest.OfferPx,
		Currency:         msg.QuoteRequest.Currency,
		ValidUntilTime:   msg.QuoteRequest.ValidUntilTime,
		ExpireTime:       msg.QuoteRequest.ExpireTime,
		QuoteType:        msg.QuoteRequest.QuoteType,
		BidSize:          msg.QuoteRequest.BidSize,
		OfferSize:        msg.QuoteRequest.OfferSize,
		Mic:              msg.QuoteRequest.Mic,
		Text:             msg.QuoteRequest.Text,
		Trailer:          msg.QuoteRequest.Trailer,
		Creator:          msg.QuoteRequest.Creator,
	}

	//fetch Header from existing session
	//In the FIX Protocol, a Quote Request message can be sent by either the initiator or the acceptor of the FIX session.
	//Determine whether we are the initiator or acceptor
	//var header *types.Header
	//if session.InitiatorAddress == msg.Creator {
	//header = session.LogonInitiator.Header
	//} else {
	//	header = session.LogonAcceptor.Header
	//}

	//fetch Trailer from existing session
	//var trailer *types.Trailer
	//if session.InitiatorAddress == msg.Creator {
	//	trailer = session.LogonInitiator.Trailer
	//} else {
	//	trailer = session.LogonAcceptor.Trailer
	//}

	newQuoteRequest := types.Quote{
		SessionID:    msg.SessionID,
		QuoteRequest: &quoteRequests,
	}

	//set Quote to store
	k.SetQuote(ctx, msg.QuoteRequest.QuoteReqID, newQuoteRequest)

	return &types.MsgQuoteRequestResponse{}, nil
}
