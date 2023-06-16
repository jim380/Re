package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityTypesRequest = "security_types_request"

var _ sdk.Msg = &MsgSecurityTypesRequest{}

func NewMsgSecurityTypesRequest(creator string, sessionID string, securityReqID string, text string, tradingSessionID string, tradingSessionSubID string, product string, securityType string, securitySubType string) *MsgSecurityTypesRequest {
	return &MsgSecurityTypesRequest{
		Creator:             creator,
		SessionID:           sessionID,
		SecurityReqID:       securityReqID,
		Text:                text,
		TradingSessionID:    tradingSessionID,
		TradingSessionSubID: tradingSessionSubID,
		Product:             product,
		SecurityType:        securityType,
		SecuritySubType:     securitySubType,
	}
}

func (msg *MsgSecurityTypesRequest) Route() string {
	return RouterKey
}

func (msg *MsgSecurityTypesRequest) Type() string {
	return TypeMsgSecurityTypesRequest
}

func (msg *MsgSecurityTypesRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityTypesRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityTypesRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
