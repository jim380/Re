package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgQuoteRequest = "quote_request"

var _ sdk.Msg = &MsgQuoteRequest{}

func NewMsgQuoteRequest(creator string, sessionID string, quoteRequest *QuoteRequest) *MsgQuoteRequest {
	return &MsgQuoteRequest{
		Creator: creator,
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

func NewQuoteRequest(header Header, quoteReqID string,  symbol string, securityID string, securityIDSource string, side string, orderQty string, futSettDate string, settlDate2 string, account string, bidPx string, offerPx string, currency string, validUntilTime string, expireTime string, quoteType string, bidSize string, offerSize string, mic  string, text string, trailer Trailer, creator string) *QuoteRequest {
	return &QuoteRequest{
      Header: &header,
	  QuoteReqID: quoteReqID,
	  Symbol: symbol,
	  SecurityID: securityID,
	  SecurityIDSource: securityIDSource,
	  Side: side,
	  OrderQty: orderQty,
	  FutSettDate: futSettDate,
	  SettlDate2: settlDate2,
	  Account: account,
	  BidPx: bidPx,
	  OfferPx: offerPx,
	  Currency: currency,
	  ValidUntilTime: validUntilTime,
	  ExpireTime: expireTime,
	  QuoteType: quoteType,
	  BidSize: bidSize,
	  OfferSize: offerSize,
	  Mic: mic,
	  Text: text,
	  Trailer: &trailer,
	  Creator:  creator,

	}
}