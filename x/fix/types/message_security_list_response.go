package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
