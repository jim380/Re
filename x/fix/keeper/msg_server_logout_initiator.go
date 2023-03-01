package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// LogoutInitiator requests for logout from the session
func (k msgServer) LogoutInitiator(goCtx context.Context, msg *types.MsgLogoutInitiator) (*types.MsgLogoutInitiatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	//TODO
	//include all cheecks
	//check that the address initiating a session logout belongs to the session

	sessionLogoutInitiator := types.SessionLogout{
		InitiatorAddress:       msg.InitiatorAddress,
		SessionID:            msg.SessionID,
		SessionLogoutInitiator: msg.SessionLogoutInitiator,
	}

	sessionLogoutInitiator.InitiatorAddress = session.InitiatorAddress
	sessionLogoutInitiator.SessionLogoutInitiator.Header = session.LogonInitiator.Header
	sessionLogoutInitiator.SessionLogoutInitiator.Trailer = session.LogonInitiator.Trailer
	sessionLogoutInitiator.SessionLogoutInitiator.Header.MsgType = "5"
	sessionLogoutInitiator.SessionLogoutInitiator.Text = msg.SessionLogoutInitiator.Text

	//set session logout to store
	k.SetSessionLogout(ctx, msg.SessionID, sessionLogoutInitiator)

	return &types.MsgLogoutInitiatorResponse{}, nil
}
