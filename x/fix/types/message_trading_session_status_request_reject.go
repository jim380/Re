package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTradingSessionStatusRequestReject = "trading_session_status_request_reject"

var _ sdk.Msg = &MsgTradingSessionStatusRequestReject{}

func NewMsgTradingSessionStatusRequestReject(creator string, refSeqNum string, refMsgType string, sessionRejectReason int32, text string) *MsgTradingSessionStatusRequestReject {
	return &MsgTradingSessionStatusRequestReject{
		Creator:             creator,
		RefSeqNum:           refSeqNum,
		RefMsgType:          refMsgType,
		SessionRejectReason: sessionRejectReason,
		Text:                text,
	}
}

func (msg *MsgTradingSessionStatusRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgTradingSessionStatusRequestReject) Type() string {
	return TypeMsgTradingSessionStatusRequestReject
}

func (msg *MsgTradingSessionStatusRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradingSessionStatusRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradingSessionStatusRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
