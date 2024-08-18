package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderMassStatusRequest = "order_mass_status_request"

var _ sdk.Msg = &MsgOrderMassStatusRequest{}

func NewMsgOrderMassStatusRequest(creator string, sessionID string, massStatusReqID string, massStatusReqType string, clOrdID string, account string, symbol string, securityID string, tradingSessionID string) *MsgOrderMassStatusRequest {
	return &MsgOrderMassStatusRequest{
		Creator:           creator,
		SessionID:         sessionID,
		MassStatusReqID:   massStatusReqID,
		MassStatusReqType: massStatusReqType,
		ClOrdID:           clOrdID,
		Account:           account,
		Symbol:            symbol,
		SecurityID:        securityID,
		TradingSessionID:  tradingSessionID,
	}
}

func (msg *MsgOrderMassStatusRequest) Route() string {
	return RouterKey
}

func (msg *MsgOrderMassStatusRequest) Type() string {
	return TypeMsgOrderMassStatusRequest
}

func (msg *MsgOrderMassStatusRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderMassStatusRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderMassStatusRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgOrderMassStatusReport = "order_mass_status_report"

var _ sdk.Msg = &MsgOrderMassStatusReport{}

func NewMsgOrderMassStatusReport(creator string, sessionID string, clOrdID string, massStatusReqID string, account string, symbol string, securityID string, tradingSessionID string, ordStatus string, execType string, ordQty string, lastPx string, cumQty string, avgPx string, leavesQty string) *MsgOrderMassStatusReport {
	return &MsgOrderMassStatusReport{
		Creator:          creator,
		SessionID:        sessionID,
		ClOrdID:          clOrdID,
		MassStatusReqID:  massStatusReqID,
		Account:          account,
		Symbol:           symbol,
		SecurityID:       securityID,
		TradingSessionID: tradingSessionID,
		OrdStatus:        ordStatus,
		ExecType:         execType,
		OrdQty:           ordQty,
		LastPx:           lastPx,
		CumQty:           cumQty,
		AvgPx:            avgPx,
		LeavesQty:        leavesQty,
	}
}

func (msg *MsgOrderMassStatusReport) Route() string {
	return RouterKey
}

func (msg *MsgOrderMassStatusReport) Type() string {
	return TypeMsgOrderMassStatusReport
}

func (msg *MsgOrderMassStatusReport) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderMassStatusReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderMassStatusReport) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgOrderMassStatusRequestReject = "order_mass_status_request_reject"

var _ sdk.Msg = &MsgOrderMassStatusRequestReject{}

func NewMsgOrderMassStatusRequestReject(creator string, sessionID string, refSeqID string, rejCode string, text string) *MsgOrderMassStatusRequestReject {
	return &MsgOrderMassStatusRequestReject{
		Creator:   creator,
		SessionID: sessionID,
		RefSeqID:  refSeqID,
		RejCode:   rejCode,
		Text:      text,
	}
}

func (msg *MsgOrderMassStatusRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgOrderMassStatusRequestReject) Type() string {
	return TypeMsgOrderMassStatusRequestReject
}

func (msg *MsgOrderMassStatusRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderMassStatusRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderMassStatusRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
