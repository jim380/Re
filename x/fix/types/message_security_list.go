package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityListResponse = "security_list_response"

var _ sdk.Msg = &MsgSecurityListResponse{}

func NewMsgSecurityListResponse(creator string, sessionID string, securityReqID string, securityResponseID string, securityRequestResult string, totNoRelatedSym string, lastFragment string, noRelatedSym string, noUnderlyings string, currency string, noLegs string, roundLot string, minTradeVol string, tradingSessionID string, tradingSessionSubID string, expirationCycle string, text string, encodedTextLen string, encodedText string) *MsgSecurityListResponse {
	return &MsgSecurityListResponse{
		Creator:               creator,
		SessionID:             sessionID,
		SecurityReqID:         securityReqID,
		SecurityResponseID:    securityResponseID,
		SecurityRequestResult: securityRequestResult,
		TotNoRelatedSym:       totNoRelatedSym,
		LastFragment:          lastFragment,
		NoRelatedSym:          noRelatedSym,
		NoUnderlyings:         noUnderlyings,
		Currency:              currency,
		NoLegs:                noLegs,
		RoundLot:              roundLot,
		MinTradeVol:           minTradeVol,
		TradingSessionID:      tradingSessionID,
		TradingSessionSubID:   tradingSessionSubID,
		ExpirationCycle:       expirationCycle,
		Text:                  text,
		EncodedTextLen:        encodedTextLen,
		EncodedText:           encodedText,
	}
}

func (msg *MsgSecurityListResponse) Route() string {
	return RouterKey
}

func (msg *MsgSecurityListResponse) Type() string {
	return TypeMsgSecurityListResponse
}

func (msg *MsgSecurityListResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityListResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityListResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

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
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgSecurityListRequestReject = "security_list_request_reject"

var _ sdk.Msg = &MsgSecurityListRequestReject{}

func NewMsgSecurityListRequestReject(creator string, sessionID string, securityReqID string, securityListRequestType string, securityRequestResult string, text string, encodedTextLen string, encodedText string) *MsgSecurityListRequestReject {
	return &MsgSecurityListRequestReject{
		Creator:                 creator,
		SessionID:               sessionID,
		SecurityReqID:           securityReqID,
		SecurityListRequestType: securityListRequestType,
		SecurityRequestResult:   securityRequestResult,
		Text:                    text,
		EncodedTextLen:          encodedTextLen,
		EncodedText:             encodedText,
	}
}

func (msg *MsgSecurityListRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgSecurityListRequestReject) Type() string {
	return TypeMsgSecurityListRequestReject
}

func (msg *MsgSecurityListRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityListRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityListRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
