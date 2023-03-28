package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgQuoteRequest         = "quote_request"
	TypeMsgQuoteAcknowledgement = "quote_acknowledgement"
	TypeMsgQuoteRequestReject   = "quote_request_reject"
)

var _ sdk.Msg = &MsgQuoteAcknowledgement{}

var _ sdk.Msg = &MsgQuoteRequest{}

var _ sdk.Msg = &MsgQuoteRequestReject{}

func NewMsgQuoteRequest(creator string, sessionID string, quoteRequest *QuoteRequest) *MsgQuoteRequest {
	return &MsgQuoteRequest{
		Creator:      creator,
		SessionID:    sessionID,
		QuoteRequest: quoteRequest,
	}
}

func (msg *MsgQuoteRequest) Route() string {
	return RouterKey
}

func (msg *MsgQuoteRequest) Type() string {
	return TypeMsgQuoteRequest
}

func (msg *MsgQuoteRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgQuoteRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgQuoteRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewQuoteRequest(quoteReqID string, symbol string, securityID string, securityIDSource string, side string, orderQty string, futSettDate string, settlDate2 string, account string, bidPx string, offerPx string, currency string, validUntilTime string, expireTime string, quoteType string, bidSize string, offerSize string, mic string, text string, creator string) *QuoteRequest {
	return &QuoteRequest{
		QuoteReqID:       quoteReqID,
		Symbol:           symbol,
		SecurityID:       securityID,
		SecurityIDSource: securityIDSource,
		Side:             side,
		OrderQty:         orderQty,
		FutSettDate:      futSettDate,
		SettlDate2:       settlDate2,
		Account:          account,
		BidPx:            bidPx,
		OfferPx:          offerPx,
		Currency:         currency,
		ValidUntilTime:   validUntilTime,
		ExpireTime:       expireTime,
		QuoteType:        quoteType,
		BidSize:          bidSize,
		OfferSize:        offerSize,
		Mic:              mic,
		Text:             text,
		Creator:          creator,
	}
}

func NewMsgQuoteAcknowledgement(creator string, sessionID string, quoteAcknowledgement *QuoteAcknowledgement) *MsgQuoteAcknowledgement {
	return &MsgQuoteAcknowledgement{
		Creator:              creator,
		SessionID:            sessionID,
		QuoteAcknowledgement: quoteAcknowledgement,
	}
}

func (msg *MsgQuoteAcknowledgement) Route() string {
	return RouterKey
}

func (msg *MsgQuoteAcknowledgement) Type() string {
	return TypeMsgQuoteAcknowledgement
}

func (msg *MsgQuoteAcknowledgement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgQuoteAcknowledgement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgQuoteAcknowledgement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgQuoteRequestReject(creator string, sessionID string, quoteRequestReject *QuoteRequestReject) *MsgQuoteRequestReject {
	return &MsgQuoteRequestReject{
		Creator:            creator,
		SessionID:          sessionID,
		QuoteRequestReject: quoteRequestReject,
	}
}

func (msg *MsgQuoteRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgQuoteRequestReject) Type() string {
	return TypeMsgQuoteRequestReject
}

func (msg *MsgQuoteRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgQuoteRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgQuoteRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
