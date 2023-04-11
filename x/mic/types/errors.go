package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mic module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrMICExistsAlready  = sdkerrors.Register(ModuleName, 1, "MIC exists already")
	ErrMICCreatorIsTaken = sdkerrors.Register(ModuleName, 2, "Account Address is used Already")
	ErrMICCreatorIsWrong = sdkerrors.Register(ModuleName, 3, "Incorrect Account Address")
	ErrMICIsNotFound     = sdkerrors.Register(ModuleName, 4, "MIC does not exist")
	ErrMICIsEmpty        = sdkerrors.Register(ModuleName, 4, "MIC is empty")
	ErrNameIsEmpty       = sdkerrors.Register(ModuleName, 5, "Name used for MIC registration is empty")
	ErrLocationIsEmpty   = sdkerrors.Register(ModuleName, 6, "Location used for MIC registration is empty")
	ErrAssetClassIsEmpty = sdkerrors.Register(ModuleName, 7, "AssetClass used for MIC registration is empty")
	ErrCurrencyIsEmpty   = sdkerrors.Register(ModuleName, 8, "Currency used for MIC registration is empty")
)
