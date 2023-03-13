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
	ErrMICIsEmpty        = sdkerrors.Register(ModuleName, 3, "MIC does not exist")
	ErrMICCreatorIsWrong = sdkerrors.Register(ModuleName, 4, "Incorrect Account Address")
)
