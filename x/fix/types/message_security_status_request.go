package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityStatusRequest = "security_status_request"

var _ sdk.Msg = &MsgSecurityStatusRequest{}

func NewMsgSecurityStatusRequest(creator string, sessionID string, securityStatusReqID string, instrument string, noUnderlyings string, underlyingInstrument string, noLegs string, instrumentLeg string, currency string, subscriptionRequestType string, tradingSessionID string, tradingSessionSubID string) *MsgSecurityStatusRequest {
	return &MsgSecurityStatusRequest{
		Creator:                 creator,
		SessionID:               sessionID,
		SecurityStatusReqID:     securityStatusReqID,
		Instrument:              instrument,
		NoUnderlyings:           noUnderlyings,
		UnderlyingInstrument:    underlyingInstrument,
		NoLegs:                  noLegs,
		InstrumentLeg:           instrumentLeg,
		Currency:                currency,
		SubscriptionRequestType: subscriptionRequestType,
		TradingSessionID:        tradingSessionID,
		TradingSessionSubID:     tradingSessionSubID,
	}
}

func (msg *MsgSecurityStatusRequest) Route() string {
	return RouterKey
}

func (msg *MsgSecurityStatusRequest) Type() string {
	return TypeMsgSecurityStatusRequest
}

func (msg *MsgSecurityStatusRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityStatusRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityStatusRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
