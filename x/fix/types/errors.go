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

	ErrSessionNameFound  = sdkerrors.Register(ModuleName, 9, "Session Name exists")
	ErrSessionSameDID    = sdkerrors.Register(ModuleName, 10, "Session can not use same DID for senderCompID and targetCompID")
	ErrEmptySession      = sdkerrors.Register(ModuleName, 11, "Session does not Exist")
	ErrWrongSession      = sdkerrors.Register(ModuleName, 12, "The Session provided does not tally with account")
	ErrIncorrectDID      = sdkerrors.Register(ModuleName, 13, "senderCompID and targetCompID does not match in session")
	ErrSessionIsAccepted = sdkerrors.Register(ModuleName, 14, "session accepted already")

	ErrQuoteIsEmpty         = sdkerrors.Register(ModuleName, 15, "Quote Is Empty")
	ErrQuoteSession         = sdkerrors.Register(ModuleName, 16, "No established FIX session between parties")
	ErrQuoteReqIDIsTaken    = sdkerrors.Register(ModuleName, 17, "QuoteReqID exists already for a given Quote Request")
	ErrQuoteReqIDIsEmpty    = sdkerrors.Register(ModuleName, 18, "QuoteReqID is empty")
	ErrQuoteSymbolIsEmpty   = sdkerrors.Register(ModuleName, 19, "Quote Symbol is empty")
	ErrQuoteSideIsEmpty     = sdkerrors.Register(ModuleName, 20, "Quote Side is empty")
	ErrQuoteBidPxIsEmpty    = sdkerrors.Register(ModuleName, 21, "Quote BidPx is empty")
	ErrQuoteCurrencyIsEmpty = sdkerrors.Register(ModuleName, 22, "Quote Currency is empty")
	ErrQuoteTypeIsEmpty     = sdkerrors.Register(ModuleName, 23, "QuoteType is empty")
)
