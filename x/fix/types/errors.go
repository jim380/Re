package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fix module sentinel errors
var (
	ErrInvalidDidDocument      = sdkerrors.Register(ModuleName, 2, "Invalid DID Document")
	ErrAccountIsEmpty          = sdkerrors.Register(ModuleName, 3, "Account is empty")
	ErrAccountUserIsNotTheSame = sdkerrors.Register(ModuleName, 4, "Account creator must be the same with DID creator")
	ErrDIDIsTaken              = sdkerrors.Register(ModuleName, 5, "DID exists with an account already")
	ErrCompanyNameIsTaken      = sdkerrors.Register(ModuleName, 5, "Company Name exists with an account already")
)
