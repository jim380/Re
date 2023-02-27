package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) TerminateLogon(goCtx context.Context, msg *types.MsgTerminateLogon) (*types.MsgTerminateLogonResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check if the session exits with the DID
	session, found := k.GetSessions(ctx, msg.SessionName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionName)
	}

	if session.InitiatorAddress != msg.InitiatorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.InitiatorAddress)
	}

	//check that session is not accepted yet
	if session.Status != "logon-request" {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsAccepted, "Account: %s", msg.SessionName)
	}

	//remove session from store
	k.RemoveSessions(ctx, msg.SessionName)

	return &types.MsgTerminateLogonResponse{}, nil
}
