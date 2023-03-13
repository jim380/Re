package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/jim380/Re/x/mic/types"
)

func TestMarketIdentificationCodeMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateMarketIdentificationCode(ctx, &types.MsgCreateMarketIdentificationCode{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestMarketIdentificationCodeMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMarketIdentificationCode
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateMarketIdentificationCode{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMarketIdentificationCode{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMarketIdentificationCode{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateMarketIdentificationCode(ctx, &types.MsgCreateMarketIdentificationCode{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateMarketIdentificationCode(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMarketIdentificationCodeMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMarketIdentificationCode
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteMarketIdentificationCode{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteMarketIdentificationCode{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteMarketIdentificationCode{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateMarketIdentificationCode(ctx, &types.MsgCreateMarketIdentificationCode{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteMarketIdentificationCode(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
