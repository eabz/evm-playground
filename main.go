package main

import (
	"encoding/hex"
	"fmt"
	"github.com/eabz/evm-playground/vm"
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
	"github.com/olympus-protocol/ogen/pkg/bls"
)

var senderSecret, _ = hex.DecodeString("29ca7497918b5526c355f627aa6219d2c8fd5dba93a735274e457b7489240df2")
var sender, _ = bls.SecretKeyFromBytes(senderSecret)

func main() {
	m, err := vm.NewVM()
	if err != nil {
		panic(err)
	}

	addr := evmc.Address{}
	h := evmc.Hash{}
	code := []byte("\x43\x60\x00\x52\x59\x60\x00\xf3")

	output, gasLeft, err := m.Execute(evmc.Istanbul, evmc.Call, false, 1, 999, addr, addr, code, h, nil, h)
	fmt.Println(output, gasLeft, err)
}
