package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMarketDataRequest = "market_data_request"

var _ sdk.Msg = &MsgMarketDataRequest{}

func NewMsgMarketDataRequest(creator string, sessionID string, mdReqID string, subscriptionRequestType int64, marketDepth int64, mdUpdateType int64, noRelatedSym int64, symbol string) *MsgMarketDataRequest {
	return &MsgMarketDataRequest{
		Creator:                 creator,
		SessionID:               sessionID,
		MdReqID:                 mdReqID,
		SubscriptionRequestType: subscriptionRequestType,
		MarketDepth:             marketDepth,
		MdUpdateType:            mdUpdateType,
		NoRelatedSym:            noRelatedSym,
		Symbol:                  symbol,
	}
}

func (msg *MsgMarketDataRequest) Route() string {
	return RouterKey
}

func (msg *MsgMarketDataRequest) Type() string {
	return TypeMsgMarketDataRequest
}

func (msg *MsgMarketDataRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarketDataRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarketDataRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgMarketDataSnapshotFullRefresh = "market_data_snapshot_full_refresh"

var _ sdk.Msg = &MsgMarketDataSnapshotFullRefresh{}

func NewMsgMarketDataSnapshotFullRefresh(creator string, mdReqID string, symbol string, noMDEntries string) *MsgMarketDataSnapshotFullRefresh {
	return &MsgMarketDataSnapshotFullRefresh{
		Creator:     creator,
		MdReqID:     mdReqID,
		Symbol:      symbol,
		NoMDEntries: noMDEntries,
	}
}

func (msg *MsgMarketDataSnapshotFullRefresh) Route() string {
	return RouterKey
}

func (msg *MsgMarketDataSnapshotFullRefresh) Type() string {
	return TypeMsgMarketDataSnapshotFullRefresh
}

func (msg *MsgMarketDataSnapshotFullRefresh) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarketDataSnapshotFullRefresh) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarketDataSnapshotFullRefresh) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgMarketDataIncremental = "market_data_incremental"

var _ sdk.Msg = &MsgMarketDataIncremental{}

func NewMsgMarketDataIncremental(creator string, mdReqID string, noMDEntries string) *MsgMarketDataIncremental {
	return &MsgMarketDataIncremental{
		Creator:     creator,
		MdReqID:     mdReqID,
		NoMDEntries: noMDEntries,
	}
}

func (msg *MsgMarketDataIncremental) Route() string {
	return RouterKey
}

func (msg *MsgMarketDataIncremental) Type() string {
	return TypeMsgMarketDataIncremental
}

func (msg *MsgMarketDataIncremental) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarketDataIncremental) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarketDataIncremental) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgMarketDataRequestReject = "market_data_request_reject"

var _ sdk.Msg = &MsgMarketDataRequestReject{}

func NewMsgMarketDataRequestReject(creator string, mdReqID string, mdReqRejReason string, text string) *MsgMarketDataRequestReject {
	return &MsgMarketDataRequestReject{
		Creator:        creator,
		MdReqID:        mdReqID,
		MdReqRejReason: mdReqRejReason,
		Text:           text,
	}
}

func (msg *MsgMarketDataRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgMarketDataRequestReject) Type() string {
	return TypeMsgMarketDataRequestReject
}

func (msg *MsgMarketDataRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarketDataRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarketDataRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
