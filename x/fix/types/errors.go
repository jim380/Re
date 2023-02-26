package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fix module sentinel errors
var (
	ErrInvalidDidDocument = sdkerrors.Register(ModuleName, 2, "Invalid DID Document")
	ErrAccountIsTaken     = sdkerrors.Register(ModuleName, 3, "Account Address is used already")
	ErrNotDIDCreator      = sdkerrors.Register(ModuleName, 4, "Account creator must be the same with DID creator")
	ErrDIDIsTaken         = sdkerrors.Register(ModuleName, 5, "DID exists with an account already")
	ErrCompanyNameIsTaken = sdkerrors.Register(ModuleName, 6, "Company Name exists with an account already")
	ErrWebsiteIstaken     = sdkerrors.Register(ModuleName, 7, "Website is used already")
	ErrNotAccountCreator  = sdkerrors.Register(ModuleName, 8, "Incorrect Account Owner")

	ErrSessionNameFound = sdkerrors.Register(ModuleName, 9, "Session Name exists")
	ErrSessionSameDID   = sdkerrors.Register(ModuleName, 10, "Session can not use same DID for senderCompID and targetCompID")
)
