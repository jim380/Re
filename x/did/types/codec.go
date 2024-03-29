package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateDID{}, "did/CreateDID", nil)
	cdc.RegisterConcrete(&MsgUpdateDID{}, "did/UpdateDID", nil)
	cdc.RegisterConcrete(&MsgDeactivateDID{}, "did/DeactivateDID", nil)
	cdc.RegisterConcrete(&MsgReActivateDID{}, "did/ReActivateDID", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDID{},
		&MsgUpdateDID{},
		&MsgDeactivateDID{},
		&MsgReActivateDID{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
