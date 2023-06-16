package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityTypesResponse = "security_types_response"

var _ sdk.Msg = &MsgSecurityTypesResponse{}

func NewMsgSecurityTypesResponse(creator string, sessionID string, securityReqID string, securityResponseID string, securityResponseType string, totNoSecurityTypes string, lastFragment string, noSecurityTypes string, securityType string, securitySubType string, product string, cFICode string, text string, tradingSessionID string, tradingSessionSubID string, subscriptionRequestType string) *MsgSecurityTypesResponse {
	return &MsgSecurityTypesResponse{
		Creator:                 creator,
		SessionID:               sessionID,
		SecurityReqID:           securityReqID,
		SecurityResponseID:      securityResponseID,
		SecurityResponseType:    securityResponseType,
		TotNoSecurityTypes:      totNoSecurityTypes,
		LastFragment:            lastFragment,
		NoSecurityTypes:         noSecurityTypes,
		SecurityType:            securityType,
		SecuritySubType:         securitySubType,
		Product:                 product,
		CFICode:                 cFICode,
		Text:                    text,
		TradingSessionID:        tradingSessionID,
		TradingSessionSubID:     tradingSessionSubID,
		SubscriptionRequestType: subscriptionRequestType,
	}
}

func (msg *MsgSecurityTypesResponse) Route() string {
	return RouterKey
}

func (msg *MsgSecurityTypesResponse) Type() string {
	return TypeMsgSecurityTypesResponse
}

func (msg *MsgSecurityTypesResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityTypesResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityTypesResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
