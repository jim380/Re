package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/mic/types"
)

// RegisterMarketIdentificationCode registers user's market identification code issued by ISO standard 10383 on the chain
func (k msgServer) RegisterMarketIdentificationCode(goCtx context.Context, msg *types.MsgRegisterMarketIdentificationCode) (*types.MsgRegisterMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that MIC doesn't exist already
	val, found := k.GetMarketIdentificationCode(ctx, msg.MIC)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrMICExistsAlready, "MIC: %s", msg.MIC)
	}

	// check if the msg creator is not the same as the owner of any existing MIC
	if msg.Creator == val.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMICCreatorIsTaken, "Creator: %s", msg.Creator)
	}

	// check that these fields are not empty using this https://www.iso20022.org/sites/default/files/ISO10383_MIC/ISO10383_MIC.pdf to know the required fields
	if msg.MIC == "" {
		return nil, sdkerrors.Wrapf(types.ErrMICIsEmpty, "MIC: %s", msg.MIC)
	}
	if msg.Operating_MIC == "" {
		return nil, sdkerrors.Wrapf(types.ErrOperatingMICIsEmpty, "MIC: %s", msg.Operating_MIC)
	}
	if msg.OPRT_SGMT == "" {
		return nil, sdkerrors.Wrapf(types.ErrOPRT_SGMTIsEmpty, "MIC: %s", msg.OPRT_SGMT)
	}
	if msg.MarketName == "" {
		return nil, sdkerrors.Wrapf(types.ErrMarketNameIsEmpty, "MIC: %s", msg.MarketName)
	}
	if msg.MarketCategory == "" {
		return nil, sdkerrors.Wrapf(types.ErrMarketCategoryIsEmpty, "MIC: %s", msg.MarketCategory)
	}
	if msg.ISOCountryCode == "" {
		return nil, sdkerrors.Wrapf(types.ErrISOCountryCodeIsEmpty, "MIC: %s", msg.ISOCountryCode)
	}
	if msg.City == "" {
		return nil, sdkerrors.Wrapf(types.ErrCityIsEmpty, "MIC: %s", msg.City)
	}
	if msg.Status == "" {
		return nil, sdkerrors.Wrapf(types.ErrStatusIsEmpty, "MIC: %s", msg.Status)
	}
	if msg.CreationDate == "" {
		return nil, sdkerrors.Wrapf(types.ErrCreationDateIsEmpty, "MIC: %s", msg.CreationDate)
	}
	if msg.LastUpdateDate == "" {
		return nil, sdkerrors.Wrapf(types.ErrLastUpdateDateIsEmpty, "MIC: %s", msg.LastUpdateDate)
	}

	var marketIdentificationCode = types.MarketIdentificationCode{
		Creator:               msg.Creator,
		MIC:                   msg.MIC,
		Operating_MIC:         msg.Operating_MIC,
		OPRT_SGMT:             msg.OPRT_SGMT,
		MarketName:            msg.MarketName,
		LegalEntityName:       msg.LegalEntityName,
		LegalEntityIdentifier: msg.LegalEntityIdentifier,
		MarketCategory:        msg.MarketCategory,
		Acronym:               msg.Acronym,
		ISOCountryCode:        msg.ISOCountryCode,
		City:                  msg.City,
		Website:               msg.Website,
		Status:                msg.Status,
		CreationDate:          msg.CreationDate,
		LastUpdateDate:        msg.LastUpdateDate,
		LastValidationDate:    msg.LastValidationDate,
		ExpiryDate:            msg.ExpiryDate,
		Comments:              msg.Comments,
	}

	// set MIC Data to store
	k.SetMarketIdentificationCode(ctx, marketIdentificationCode)

	return &types.MsgRegisterMarketIdentificationCodeResponse{}, nil
}

// UpdateMarketIdentificationCode updates user's market identification code issued by ISO standard 10383 on the chain
func (k msgServer) UpdateMarketIdentificationCode(goCtx context.Context, msg *types.MsgUpdateMarketIdentificationCode) (*types.MsgUpdateMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check that the old MIC exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.Old_MIC)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMICIsNotFound, "MIC: %s", msg.Old_MIC)
	}

	// check that new MIC doesn't exist already
	_, found = k.GetMarketIdentificationCode(ctx, msg.New_MIC)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrMICExistsAlready, "MIC: %s", msg.New_MIC)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMICCreatorIsWrong, "Creator: %s", msg.Creator)
	}

	// check that these fields are not empty using this https://www.iso20022.org/sites/default/files/ISO10383_MIC/ISO10383_MIC.pdf to know the required fields
	if msg.New_MIC == "" {
		return nil, sdkerrors.Wrapf(types.ErrMICIsEmpty, "MIC: %s", msg.New_MIC)
	}
	if msg.Operating_MIC == "" {
		return nil, sdkerrors.Wrapf(types.ErrOperatingMICIsEmpty, "MIC: %s", msg.Operating_MIC)
	}
	if msg.OPRT_SGMT == "" {
		return nil, sdkerrors.Wrapf(types.ErrOPRT_SGMTIsEmpty, "MIC: %s", msg.OPRT_SGMT)
	}
	if msg.MarketName == "" {
		return nil, sdkerrors.Wrapf(types.ErrMarketNameIsEmpty, "MIC: %s", msg.MarketName)
	}
	if msg.MarketCategory == "" {
		return nil, sdkerrors.Wrapf(types.ErrMarketCategoryIsEmpty, "MIC: %s", msg.MarketCategory)
	}
	if msg.ISOCountryCode == "" {
		return nil, sdkerrors.Wrapf(types.ErrISOCountryCodeIsEmpty, "MIC: %s", msg.ISOCountryCode)
	}
	if msg.City == "" {
		return nil, sdkerrors.Wrapf(types.ErrCityIsEmpty, "MIC: %s", msg.City)
	}
	if msg.Status == "" {
		return nil, sdkerrors.Wrapf(types.ErrStatusIsEmpty, "MIC: %s", msg.Status)
	}
	if msg.CreationDate == "" {
		return nil, sdkerrors.Wrapf(types.ErrCreationDateIsEmpty, "MIC: %s", msg.CreationDate)
	}
	if msg.LastUpdateDate == "" {
		return nil, sdkerrors.Wrapf(types.ErrLastUpdateDateIsEmpty, "MIC: %s", msg.LastUpdateDate)
	}

	var editedMarketIdentificationCode = types.MarketIdentificationCode{
		Creator:               msg.Creator,
		MIC:                   msg.New_MIC,
		Operating_MIC:         msg.Operating_MIC,
		OPRT_SGMT:             msg.OPRT_SGMT,
		MarketName:            msg.MarketName,
		LegalEntityName:       msg.LegalEntityName,
		LegalEntityIdentifier: msg.LegalEntityIdentifier,
		MarketCategory:        msg.MarketCategory,
		Acronym:               msg.Acronym,
		ISOCountryCode:        msg.ISOCountryCode,
		City:                  msg.City,
		Website:               msg.Website,
		Status:                msg.Status,
		CreationDate:          msg.CreationDate,
		LastUpdateDate:        msg.LastUpdateDate,
		LastValidationDate:    msg.LastValidationDate,
		ExpiryDate:            msg.ExpiryDate,
		Comments:              msg.Comments,
	}

	// remove old MIC from store
	k.RemoveMarketIdentificationCode(ctx, val.MIC)

	// set edited MIC Data to store
	k.SetMarketIdentificationCode(ctx, editedMarketIdentificationCode)

	return &types.MsgUpdateMarketIdentificationCodeResponse{}, nil
}

// DeleteMarketIdentificationCode deletes user's market identification code from the store
func (k msgServer) DeleteMarketIdentificationCode(goCtx context.Context, msg *types.MsgDeleteMarketIdentificationCode) (*types.MsgDeleteMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.MIC)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMICIsEmpty, "MIC: %s", msg.MIC)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMICCreatorIsWrong, "Creator: %s", msg.Creator)
	}

	//remove MIC from store
	k.RemoveMarketIdentificationCode(ctx, msg.MIC)

	return &types.MsgDeleteMarketIdentificationCodeResponse{}, nil
}
