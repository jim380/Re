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

	docWithSeq := types.NewDIDDocumentWithSeq(msg.Document, uint64(seq), msg.FromAddress)

	getDIDDocument := keeper.GetDIDDocumentWithCreator(ctx, docWithSeq)
	if !getDIDDocument.Empty() {
		if getDIDDocument.Deactivated() {
			return nil, sdkerrors.Wrapf(types.ErrDIDDeactivated, "DID: %s", msg.Did)
		}
		return nil, sdkerrors.Wrapf(types.ErrDIDExists, "DID: %s", msg.Did)
	}
	if getDIDDocument.Creator == msg.FromAddress {
		return nil, sdkerrors.Wrapf(types.ErrAccountExists, "AccountAddress: %s", msg.FromAddress)
	}

	keeper.SetDIDDocument(ctx, msg.Did, docWithSeq)
	keeper.SetDIDDocumentWithCreator(ctx, docWithSeq, docWithSeq)
	return &types.MsgCreateDIDResponse{}, nil
}

// UpdateDID updates DID document on RE Protocol
func (m msgServer) UpdateDID(goCtx context.Context, msg *types.MsgUpdateDID) (*types.MsgUpdateDIDResponse, error) {
	keeper := m.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)

	docWithSeq := keeper.GetDIDDocument(ctx, msg.Did)
	if docWithSeq.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrDIDNotFound, "DID: %s", msg.Did)
	}
	if docWithSeq.Deactivated() {
		return nil, sdkerrors.Wrapf(types.ErrDIDDeactivated, "DID: %s", msg.Did)
	}

	newSeq, err := VerifyDIDOwnership(msg.Document, docWithSeq.Sequence, docWithSeq.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, err
	}

	newDocWithSeq := types.NewDIDDocumentWithSeq(msg.Document, newSeq, msg.FromAddress)
	if docWithSeq.Document.Id != newDocWithSeq.Document.Id {
		return nil, sdkerrors.Wrapf(types.ErrNotTheSameDID, "DID: %s", msg.Did)
	}

	//only account owner can update DID
	if docWithSeq.Creator != newDocWithSeq.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	//verification method's Public Key and Type for updating must be the same at DID creation
	for i := range docWithSeq.Document.VerificationMethods {
		if docWithSeq.Document.VerificationMethods[i].PublicKeyBase58 != newDocWithSeq.Document.VerificationMethods[i].PublicKeyBase58 && docWithSeq.Document.VerificationMethods[i].Type != newDocWithSeq.Document.VerificationMethods[i].Type {
			return nil, sdkerrors.Wrapf(types.ErrSigVerificationFailed, "Verification Methods: %s", newDocWithSeq.Document.VerificationMethods[i])
		}
	}

	keeper.SetDIDDocument(ctx, newDocWithSeq.Document.Id, newDocWithSeq)
	return &types.MsgUpdateDIDResponse{}, nil
}

// DeactivateDID deactivates DID document on RE Protocol
func (m msgServer) DeactivateDID(goCtx context.Context, msg *types.MsgDeactivateDID) (*types.MsgDeactivateDIDResponse, error) {
	keeper := m.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)

	docWithSeq := keeper.GetDIDDocument(ctx, msg.Did)
	if docWithSeq.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrDIDNotFound, "DID: %s", msg.Did)
	}
	if docWithSeq.Deactivated() {
		return nil, sdkerrors.Wrapf(types.ErrDIDDeactivated, "DID: %s", msg.Did)
	}

	doc := types.DIDDocument{
		Id: msg.Did,
	}

	if docWithSeq.Creator != msg.FromAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	newSeq, err := VerifyDIDOwnership(&doc, docWithSeq.Sequence, docWithSeq.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, err
	}

	keeper.SetDeactivatedDIDDocument(ctx, msg.FromAddress, docWithSeq)
	keeper.SetDIDDocument(ctx, msg.Did, docWithSeq.Deactivate(newSeq, msg.FromAddress))
	return &types.MsgDeactivateDIDResponse{}, nil

}

// ReActivateDID reactivates DID document on RE Protocol
func (m msgServer) ReactivateDID(goCtx context.Context, msg *types.MsgReActivateDID) (*types.MsgReActivateDIDResponse, error) {
	keeper := m.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)

	docWithSeq := keeper.GetDeactivatedDIDDocument(ctx, msg.FromAddress)

	if docWithSeq.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	//get the DID from the GetDeactivatedDIDDocument to prevent DID that is not deactivated
	activedocWithSeq := keeper.GetDIDDocument(ctx, docWithSeq.Document.Id)

	if !activedocWithSeq.Document.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrDIDNotDeactivated, "DID: %s", activedocWithSeq.Document.Id)
	}
	if docWithSeq.Creator != msg.FromAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotUserAccount, "Account Address: %s", msg.FromAddress)
	}

	// if all checks passes, reset the DIDDocument
	keeper.SetDIDDocument(ctx, docWithSeq.Document.Id, docWithSeq)

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
