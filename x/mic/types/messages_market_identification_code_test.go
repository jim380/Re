package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateMarketIdentificationCode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMarketIdentificationCode
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMarketIdentificationCode{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateMarketIdentificationCode{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateMarketIdentificationCode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMarketIdentificationCode
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMarketIdentificationCode{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateMarketIdentificationCode{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteMarketIdentificationCode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteMarketIdentificationCode
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteMarketIdentificationCode{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteMarketIdentificationCode{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
