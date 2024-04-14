package app

import (
	"strings"
)

var wasmCapabilities = []string{
	"iterator",
	"staking",
	"stargate",
	"cosmwasm_1_1",
	"cosmwasm_1_2",
	"cosmwasm_1_3",
	"cosmwasm_1_4",
}

func GetWasmCapabilities() string {
	return strings.Join(wasmCapabilities, ",")
}
