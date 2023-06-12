package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityStatusResponse = "security_status_response"

var _ sdk.Msg = &MsgSecurityStatusResponse{}

func NewMsgSecurityStatusResponse(creator string, sessionID string, securityStatusReqID string, noUnderlyings string, underlyingInstrument string, noLegs string, instrumentLeg string, currency string, tradingSessionID string, tradingSessionSubID string, unsolicitedIndicator string, securityTradingStatus string, financialStatus string, corporateAction string, haltReason string, inViewOfCommon string, dueToRelated string, buyVolume string, sellVolume string, highPx string, lowPx string, lastPx string, transactTime string, adjustment string, text string) *MsgSecurityStatusResponse {
	return &MsgSecurityStatusResponse{
		Creator:               creator,
		SessionID:             sessionID,
		SecurityStatusReqID:   securityStatusReqID,
		NoUnderlyings:         noUnderlyings,
		UnderlyingInstrument:  underlyingInstrument,
		NoLegs:                noLegs,
		InstrumentLeg:         instrumentLeg,
		Currency:              currency,
		TradingSessionID:      tradingSessionID,
		TradingSessionSubID:   tradingSessionSubID,
		UnsolicitedIndicator:  unsolicitedIndicator,
		SecurityTradingStatus: securityTradingStatus,
		FinancialStatus:       financialStatus,
		CorporateAction:       corporateAction,
		HaltReason:            haltReason,
		InViewOfCommon:        inViewOfCommon,
		DueToRelated:          dueToRelated,
		BuyVolume:             buyVolume,
		SellVolume:            sellVolume,
		HighPx:                highPx,
		LowPx:                 lowPx,
		LastPx:                lastPx,
		TransactTime:          transactTime,
		Adjustment:            adjustment,
		Text:                  text,
	}
}

func (msg *MsgSecurityStatusResponse) Route() string {
	return RouterKey
}

func (msg *MsgSecurityStatusResponse) Type() string {
	return TypeMsgSecurityStatusResponse
}

func (msg *MsgSecurityStatusResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityStatusResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityStatusResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
