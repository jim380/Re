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

	//fetch existing Quote Request and use it to check for multiple Quote Requests
	//same quote with quoteReqID is not allowed
	quoteRequest, _ := k.GetQuote(ctx, msg.SessionID)
	//if quoteRequest.QuoteRequest.QuoteReqID == msg.QuoteRequest.

	//fetch Header from existing session
	//In the FIX Protocol, a Quote Request message can be sent by either the initiator or the acceptor of the FIX session.
	// Determine whether we are the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	//get market identification code from MIC module

	newQuoteRequest := types.Quote{
		SessionID:    msg.SessionID,
		QuoteRequest: msg.QuoteRequest,
	}

	//set Quote to store
	k.SetQuote(ctx, msg.SessionID, newQuoteRequest)

	return &types.MsgQuoteRequestResponse{}, nil
}
