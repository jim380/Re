package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/did/internal/secp256k1util"
	"github.com/jim380/Re/x/did/types"
)

// CreateDID registers DID document on RE Protocol
func (m msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	keeper := m.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)

	seq := types.InitialSequence
	_, err := VerifyDIDOwnership(msg.Document, seq, msg.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, err
	}

	newDocWithSeq := types.NewDIDDocumentWithSeq(msg.Document, uint64(seq), msg.FromAddress)

	for _, did := range keeper.ListDIDs(ctx) {
		existingDIDDocument := keeper.GetDIDDocument(ctx, did)
		if !existingDIDDocument.Empty() {
			if existingDIDDocument.Deactivated() {
				return nil, sdkerrors.Wrapf(types.ErrDIDDeactivated, "DID: %s", msg.Did)
			}
			if existingDIDDocument.Document.Id == msg.Did {
				return nil, sdkerrors.Wrapf(types.ErrDIDExists, "DID: %s", msg.Did)
			}
		}
		if existingDIDDocument.Creator == msg.FromAddress {
			return nil, sdkerrors.Wrapf(types.ErrAccountExists, "AccountAddress: %s", msg.FromAddress)
		}
	}

	// set new document to store
	keeper.SetDIDDocument(ctx, msg.Did, newDocWithSeq)
	return &types.MsgCreateDIDResponse{}, nil
}

// UpdateDID updates DID document on RE Protocol
func (m msgServer) UpdateDID(goCtx context.Context, msg *types.MsgUpdateDID) (*types.MsgUpdateDIDResponse, error) {
	keeper := m.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)

	existingDocWithSeq := keeper.GetDIDDocument(ctx, msg.Did)
	if existingDocWithSeq.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrDIDNotFound, "DID: %s", msg.Did)
	}
	if existingDocWithSeq.Deactivated() {
		return nil, sdkerrors.Wrapf(types.ErrDIDDeactivated, "DID: %s", msg.Did)
	}

	newSeq, err := VerifyDIDOwnership(msg.Document, existingDocWithSeq.Sequence, existingDocWithSeq.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, err
	}

	editedDocWithSeq := types.NewDIDDocumentWithSeq(msg.Document, newSeq, msg.FromAddress)
	if existingDocWithSeq.Document.Id != editedDocWithSeq.Document.Id {
		return nil, sdkerrors.Wrapf(types.ErrNotTheSameDID, "DID: %s", msg.Did)
	}

	// only account owner can update DID
	if existingDocWithSeq.Creator != editedDocWithSeq.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	// verification method's Public Key and Type for updating must be the same at DID creation
	for i := range existingDocWithSeq.Document.VerificationMethods {
		if existingDocWithSeq.Document.VerificationMethods[i].PublicKeyBase58 != editedDocWithSeq.Document.VerificationMethods[i].PublicKeyBase58 && existingDocWithSeq.Document.VerificationMethods[i].Type != editedDocWithSeq.Document.VerificationMethods[i].Type {
			return nil, sdkerrors.Wrapf(types.ErrSigVerificationFailed, "Verification Methods: %s", editedDocWithSeq.Document.VerificationMethods[i])
		}
	}

	keeper.SetDIDDocument(ctx, editedDocWithSeq.Document.Id, editedDocWithSeq)
	return &types.MsgUpdateDIDResponse{}, nil
}

// DeactivateDID deactivates DID document on RE Protocol
func (m msgServer) DeactivateDID(goCtx context.Context, msg *types.MsgDeactivateDID) (*types.MsgDeactivateDIDResponse, error) {
	keeper := m.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)

	existingDocWithSeq := keeper.GetDIDDocument(ctx, msg.Did)
	if existingDocWithSeq.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrDIDNotFound, "DID: %s", msg.Did)
	}
	if existingDocWithSeq.Deactivated() {
		return nil, sdkerrors.Wrapf(types.ErrDIDDeactivated, "DID: %s", msg.Did)
	}

	doc := types.DIDDocument{
		Id: msg.Did,
	}

	if existingDocWithSeq.Creator != msg.FromAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	newSeq, err := VerifyDIDOwnership(&doc, existingDocWithSeq.Sequence, existingDocWithSeq.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, err
	}

	keeper.SetDeactivatedDIDDocument(ctx, msg.FromAddress, existingDocWithSeq)
	keeper.SetDIDDocument(ctx, msg.Did, existingDocWithSeq.Deactivate(newSeq, msg.FromAddress))
	return &types.MsgDeactivateDIDResponse{}, nil

}

// ReActivateDID reactivates DID document on RE Protocol
func (m msgServer) ReactivateDID(goCtx context.Context, msg *types.MsgReActivateDID) (*types.MsgReActivateDIDResponse, error) {
	keeper := m.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)

	existingDocWithSeq := keeper.GetDeactivatedDIDDocument(ctx, msg.FromAddress)

	if existingDocWithSeq.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	// get the DID from the GetDeactivatedDIDDocument to prevent DID that is not deactivated
	activedocWithSeq := keeper.GetDIDDocument(ctx, existingDocWithSeq.Document.Id)

	if !activedocWithSeq.Document.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrDIDNotDeactivated, "DID: %s", activedocWithSeq.Document.Id)
	}
	if existingDocWithSeq.Creator != msg.FromAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	// if all checks passes, reset the DIDDocument
	keeper.SetDIDDocument(ctx, existingDocWithSeq.Document.Id, existingDocWithSeq)

	return &types.MsgReActivateDIDResponse{}, nil

}

func VerifyDIDOwnership(signData *types.DIDDocument, seq uint64, doc *types.DIDDocument, verificationMethodID string, sig []byte) (uint64, error) {
	verificationMethod, ok := doc.VerificationMethodFrom(doc.Authentications, verificationMethodID)
	if !ok {
		return 0, sdkerrors.Wrapf(types.ErrVerificationMethodIDNotFound, "VerificationMethodId: %s", verificationMethodID)
	}

	// TODO: Currently, only ES256K1 is supported to verify DID ownership.
	//       It makes sense for now, since a DID is derived from a Secp256k1 public key.
	//       But, need to support other key types (according to verificationMethod.Type).
	if verificationMethod.Type != types.ES256K_2019 && verificationMethod.Type != types.ES256K_2018 {
		return 0, sdkerrors.Wrapf(types.ErrVerificationMethodKeyTypeNotImplemented, "VerificationMethod: %v", verificationMethod.Type)
	}

	pubKeySecp256k1, err := secp256k1util.PubKeyFromBase58(verificationMethod.PublicKeyBase58)
	if err != nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidSecp256k1PublicKey, "PublicKey: %v", verificationMethod.PublicKeyBase58)
	}

	newSeq, ok := types.Verify(sig, signData, seq, pubKeySecp256k1)
	if !ok {
		return 0, types.ErrSigVerificationFailed
	}

	return newSeq, nil
}
