package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityListRequest = "security_list_request"

var _ sdk.Msg = &MsgSecurityListRequest{}

func NewMsgSecurityListRequest(creator string, sessionID string, securityReqID string, securityListRequestType string, noUnderlyings string, noLegs string, currency string, text string, encodedTextLen string, encodedText string, tradingSessionID string, tradingSessionSubID string, subscriptionRequestType string) *MsgSecurityListRequest {
	return &MsgSecurityListRequest{
		Creator:                 creator,
		SessionID:               sessionID,
		SecurityReqID:           securityReqID,
		SecurityListRequestType: securityListRequestType,
		NoUnderlyings:           noUnderlyings,
		NoLegs:                  noLegs,
		Currency:                currency,
		Text:                    text,
		EncodedTextLen:          encodedTextLen,
		EncodedText:             encodedText,
		TradingSessionID:        tradingSessionID,
		TradingSessionSubID:     tradingSessionSubID,
		SubscriptionRequestType: subscriptionRequestType,
	}
}

func (msg *MsgSecurityListRequest) Route() string {
	return RouterKey
}

func (msg *MsgSecurityListRequest) Type() string {
	return TypeMsgSecurityListRequest
}

func (msg *MsgSecurityListRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityListRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityListRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
