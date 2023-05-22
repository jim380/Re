package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderMassStatusRequest = "order_mass_status_request"

var _ sdk.Msg = &MsgOrderMassStatusRequest{}

func NewMsgOrderMassStatusRequest(creator string, sessionID string, massStatusReqID string, massStatusReqType string, clOrdID string, account string, symbol string, securityID string, tradingSessionID string) *MsgOrderMassStatusRequest {
	return &MsgOrderMassStatusRequest{
		Creator:           creator,
		SessionID:         sessionID,
		MassStatusReqID:   massStatusReqID,
		MassStatusReqType: massStatusReqType,
		ClOrdID:           clOrdID,
		Account:           account,
		Symbol:            symbol,
		SecurityID:        securityID,
		TradingSessionID:  tradingSessionID,
	}
}

func (msg *MsgOrderMassStatusRequest) Route() string {
	return RouterKey
}

func (msg *MsgOrderMassStatusRequest) Type() string {
	return TypeMsgOrderMassStatusRequest
}

func (msg *MsgOrderMassStatusRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderMassStatusRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderMassStatusRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
