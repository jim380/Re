package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module errors
var (
	ErrDIDExists                               = sdkerrors.Register(ModuleName, 2, "DID already exists")
	ErrInvalidDID                              = sdkerrors.Register(ModuleName, 3, "Invalid DID")
	ErrInvalidDIDDocument                      = sdkerrors.Register(ModuleName, 4, "Invalid DID Document")
	ErrDIDNotFound                             = sdkerrors.Register(ModuleName, 5, "DID not found")
	ErrInvalidSignature                        = sdkerrors.Register(ModuleName, 6, "Invalid signature")
	ErrInvalidVerificationMethodID             = sdkerrors.Register(ModuleName, 7, "Invalid VerificationMethodID")
	ErrVerificationMethodIDNotFound            = sdkerrors.Register(ModuleName, 8, "VerificationMethodID not found")
	ErrSigVerificationFailed                   = sdkerrors.Register(ModuleName, 9, "DID signature verification was failed")
	ErrInvalidSecp256k1PublicKey               = sdkerrors.Register(ModuleName, 10, "Invalid Secp256k1 public key")
	ErrInvalidNetworkID                        = sdkerrors.Register(ModuleName, 11, "Invalid network ID")
	ErrInvalidDIDDocumentWithSeq               = sdkerrors.Register(ModuleName, 12, "Invalid DIDDocumentWithSeq")
	ErrDIDDeactivated                          = sdkerrors.Register(ModuleName, 13, "DID was already deactivated")
	CodeInvalidKeyController                   = sdkerrors.Register(ModuleName, 14, "Invalid key controller")
	ErrVerificationMethodKeyTypeNotImplemented = sdkerrors.Register(ModuleName, 15, "Verification not implemented with key type")
	ErrAccountExists                           = sdkerrors.Register(ModuleName, 16, "DID already exists with this Account Address")
	ErrNotUserAccount                          = sdkerrors.Register(ModuleName, 17, "This is not the Account Address that created the DID")
	ErrNotTheSameDID                           = sdkerrors.Register(ModuleName, 18, "DID is unique, hence it can not be changed")
	ErrDIDNotDeactivated                       = sdkerrors.Register(ModuleName, 19, "DID is not deactivated")
)
