package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	fixTypes "github.com/jim380/Re/x/fix/types"
	"github.com/jim380/Re/x/oracle/types"
)

func (k msgServer) CosmoshubTxs(goCtx context.Context, msg *types.MsgCosmoshubTxs) (*types.MsgCosmoshubTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}
	if msg.Creator != msg.Address {
		return nil, sdkerrors.Wrapf(types.ErrWrongAddress, "Address: %s", msg.Address)
	}

	// set cosmoshub txs, map to the fix module

	// test, update fix module from oracle
	Accounts := []fixTypes.AccountRegistration{
		{
			Address:          "re1a34ffvsapzu6f35xzrzlc2nl6ytnrly0cg9n0mhhhhpppppppp",
			CompanyName:      "CompanyAu",
			Website:          "www.companya.commm",
			SocialMediaLinks: "@CompanyppA",
			CreatedAt:        constants.CreatedAt,
		},
		{
			Address:          "re1a34ffvsapzu6f35xzrzlc2nl6ytnrly0cg9n0myuuyyyyyyyyy",
			CompanyName:      "CompanyAbbb",
			Website:          "www.companya.com",
			SocialMediaLinks: "@CompanyAaaa",
			CreatedAt:        constants.CreatedAt,
		},
		{
			Address:          "re1a34ffvsapzu6f35xzrzlc2nl6ytnrly0cg9n0mtygggvzzzzzzz",
			CompanyName:      "CompanyAbbb",
			Website:          "www.companya.commmmm",
			SocialMediaLinks: "@CompanyAui",
			CreatedAt:        constants.CreatedAt,
		},
	}

	for i, account := range Accounts {
		// set
		log.Println("When it is in range", account, i)
		k.fixKeeper.SetAccountRegistration(ctx, account.Address, account)
	}

	/*
		account := fixTypes.AccountRegistration{
			Address:          msg.Creator,
			CompanyName:      "danielllll",
			Website:          "www.daniel",
			SocialMediaLinks: "@danii",
			CreatedAt:        constants.CreatedAt,
		}

		// set
		k.fixKeeper.SetAccountRegistration(ctx, msg.Creator, account)
	*/
	oracle := types.MultiChainTxOracle{
		OracleId:  msg.OracleId,
		Address:   msg.Address,
		Timestamp: constants.CreatedAt,
	}

	// set oracle to store
	k.SetMultiChainTxOracle(ctx, msg.OracleId, oracle)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgCosmoshubTxsResponse{}, err
}
