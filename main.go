package main

import (
	"encoding/hex"
	"fmt"
	"github.com/eabz/evm-playground/vm"
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
	"github.com/olympus-protocol/ogen/pkg/bls"
	"github.com/olympus-protocol/ogen/pkg/primitives"
	"io/ioutil"
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
	code, _ := ioutil.ReadFile("./bytecodes/storage/compiled/Storage.bin")

	//input, _ := hex.DecodeString("a9059cbb0000000000000000000000007e747efe1e22e6eb0a92c3a12e334229bf20fd56000000000000000000000000000000000000000000000690836c0af5f5600000")

	host := vm.NewHost(primitives.Execution{}, primitives.Block{})

	output, gasLeft, err := m.Execute(host, evmc.Istanbul, evmc.Create2, false, 1, 10000000, addr, addr, nil, h, code, h)
	fmt.Println(output, gasLeft, err)
}
