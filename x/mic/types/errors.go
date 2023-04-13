package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mic module sentinel errors
var (
	ErrSample                = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrMICExistsAlready      = sdkerrors.Register(ModuleName, 1, "MIC exists already")
	ErrMICCreatorIsTaken     = sdkerrors.Register(ModuleName, 2, "Account Address is used Already")
	ErrMICCreatorIsWrong     = sdkerrors.Register(ModuleName, 3, "Incorrect Account Address")
	ErrMICIsNotFound         = sdkerrors.Register(ModuleName, 4, "MIC does not exist")
	ErrMICIsEmpty            = sdkerrors.Register(ModuleName, 5, "MIC is empty")
	ErrOperatingMICIsEmpty   = sdkerrors.Register(ModuleName, 6, "Operating_MIC is empty")
	ErrOPRT_SGMTIsEmpty      = sdkerrors.Register(ModuleName, 7, "OPRT_SGMT is empty")
	ErrMarketNameIsEmpty     = sdkerrors.Register(ModuleName, 8, "MarketName is empty")
	ErrMarketCategoryIsEmpty = sdkerrors.Register(ModuleName, 9, "MarketCategory is empty")
	ErrISOCountryCodeIsEmpty = sdkerrors.Register(ModuleName, 10, "ISOCountryCode is empty")
	ErrCityIsEmpty           = sdkerrors.Register(ModuleName, 11, "City is empty")
	ErrStatusIsEmpty         = sdkerrors.Register(ModuleName, 12, "Status is empty")
	ErrCreationDateIsEmpty   = sdkerrors.Register(ModuleName, 13, "CreationDate is empty")
	ErrLastUpdateDateIsEmpty = sdkerrors.Register(ModuleName, 14, "LastUpdateDate is empty")
)
