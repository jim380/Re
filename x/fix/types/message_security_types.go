package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
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
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

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
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgSecurityTypesRequestReject = "security_types_request_reject"

var _ sdk.Msg = &MsgSecurityTypesRequestReject{}

func NewMsgSecurityTypesRequestReject(creator string, sessionID string, securityReqID string, rejectReason string, text string) *MsgSecurityTypesRequestReject {
	return &MsgSecurityTypesRequestReject{
		Creator:       creator,
		SessionID:     sessionID,
		SecurityReqID: securityReqID,
		RejectReason:  rejectReason,
		Text:          text,
	}
}

func (msg *MsgSecurityTypesRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgSecurityTypesRequestReject) Type() string {
	return TypeMsgSecurityTypesRequestReject
}

func (msg *MsgSecurityTypesRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityTypesRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityTypesRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
