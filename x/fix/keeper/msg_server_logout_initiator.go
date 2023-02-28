package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) LogoutInitiator(goCtx context.Context, msg *types.MsgLogoutInitiator) (*types.MsgLogoutInitiatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionName)
	}

	//TODO
	//include all cheecks
	//check that the address initiating a session logout belongs to the session

	sessionLogout := types.SessionLogout{
		InitiatorAddress:       msg.InitiatorAddress,
		SessionName:            msg.SessionName,
		SessionLogoutInitiator: msg.SessionLogoutInitiator,
	}

	sessionLogout.InitiatorAddress = session.InitiatorAddress
	sessionLogout.SessionLogoutInitiator.Header = session.LogonInitiator.Header
	sessionLogout.SessionLogoutInitiator.Trailer.CheckSum = session.LogonInitiator.Trailer.CheckSum
	sessionLogout.SessionLogoutInitiator.Header.MsgType = "5"
	sessionLogout.SessionLogoutInitiator.Text = msg.SessionLogoutInitiator.Text

	//set session logout to store
	k.SetSessionLogout(ctx, msg.SessionName, sessionLogout)

	return &types.MsgLogoutInitiatorResponse{}, nil
}
