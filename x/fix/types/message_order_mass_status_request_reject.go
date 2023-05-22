package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderMassStatusRequestReject = "order_mass_status_request_reject"

var _ sdk.Msg = &MsgOrderMassStatusRequestReject{}

func NewMsgOrderMassStatusRequestReject(creator string, sessionID string, refSeqID string, rejCode string, text string) *MsgOrderMassStatusRequestReject {
	return &MsgOrderMassStatusRequestReject{
		Creator:   creator,
		SessionID: sessionID,
		RefSeqID:  refSeqID,
		RejCode:   rejCode,
		Text:      text,
	}
}

func (msg *MsgOrderMassStatusRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgOrderMassStatusRequestReject) Type() string {
	return TypeMsgOrderMassStatusRequestReject
}

func (msg *MsgOrderMassStatusRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderMassStatusRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderMassStatusRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
