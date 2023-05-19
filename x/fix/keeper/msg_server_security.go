package keeper

import (
	"context"

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

	// set new security definition request to store
	k.SetSecurity(ctx, msg.SecurityReqID, securityDefinitionRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityDefinitionRequestResponse{}, err
}

func (k msgServer) SecurityDefinition(goCtx context.Context, msg *types.MsgSecurityDefinition) (*types.MsgSecurityDefinitionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityDefinitionResponse{}, nil
}

func (k msgServer) SecurityDefinitionRequestReject(goCtx context.Context, msg *types.MsgSecurityDefinitionRequestReject) (*types.MsgSecurityDefinitionRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityDefinitionRequestRejectResponse{}, nil
}
