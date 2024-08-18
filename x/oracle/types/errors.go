package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// Error Code Enums
const (
	errSample = iota + 101
	errWrongAddress
)

// x/oracle module sentinel errors
var (
	ErrSample       = sdkerrors.Register(ModuleName, errSample, "sample error")
	ErrWrongAddress = sdkerrors.Register(ModuleName, errWrongAddress, "The account address provided does not match the creator of the oracle")
)
