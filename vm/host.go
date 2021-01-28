package vm

import (
	"fmt"
	"github.com/eabz/evm-playground/utils"
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
	"github.com/olympus-protocol/ogen/pkg/primitives"
	"time"
)

type Host struct {
	txContext    primitives.Execution
	blockContext primitives.Block
}

// AccountExists checks if a specif contract account exits withint the bytecode storage
func (h Host) AccountExists(addr evmc.Address) bool {
	fmt.Println("AccountExists")
	fmt.Println(addr)
	return true
}

func (h Host) GetStorage(addr evmc.Address, key evmc.Hash) evmc.Hash {
	fmt.Println("GetStorage")
	fmt.Println(addr, key)
	return evmc.Hash{}
}

func (h Host) SetStorage(addr evmc.Address, key evmc.Hash, value evmc.Hash) evmc.StorageStatus {
	fmt.Println("SetStorage")
	fmt.Println(addr, key, value)
	return evmc.StorageAdded
}

func (h Host) GetBalance(addr evmc.Address) evmc.Hash {
	balance := utils.NumToWei(10000000)
	fmt.Println("GetBalance")
	return utils.StringToHash(balance)
}

func (h Host) GetCodeSize(addr evmc.Address) int {
	fmt.Println("GetCodeSize")
	fmt.Println(addr)
	return 0
}

func (h Host) GetCodeHash(addr evmc.Address) evmc.Hash {
	fmt.Println("GetCodeHash")
	fmt.Println(addr)
	return evmc.Hash{}
}

func (h Host) GetCode(addr evmc.Address) []byte {
	fmt.Println("GetCode")
	fmt.Println(addr)
	return []byte{}
}

func (h Host) Selfdestruct(addr evmc.Address, beneficiary evmc.Address) {
	fmt.Println("Selfdestruct")
	fmt.Println(addr, beneficiary)
	return
}

func (h Host) GetTxContext() evmc.TxContext {
	ctx := evmc.TxContext{
		GasPrice:   utils.StringToHash("100"),
		Origin:     evmc.Address{},
		Coinbase:   evmc.Address{},
		Number:     1,
		Timestamp:  time.Now().Unix(),
		GasLimit:   1000000,
		Difficulty: evmc.Hash{},
		ChainID:    utils.StringToHash("1"),
	}
	fmt.Println("GetTxContext")
	return ctx
}

func (h Host) GetBlockHash(number int64) evmc.Hash {
	fmt.Println("GetBlockHash")
	return evmc.Hash{}
}

func (h Host) EmitLog(addr evmc.Address, topics []evmc.Hash, data []byte) {
	fmt.Println("EmitLog")
}

func (h Host) Call(kind evmc.CallKind, destination evmc.Address, sender evmc.Address, value evmc.Hash, input []byte, gas int64, depth int, static bool, salt evmc.Hash) (output []byte, gasLeft int64, createAddr evmc.Address, err error) {
	fmt.Println("Call")
	return output, 0, [20]byte{}, err
}

func NewHost(tx primitives.Execution, block primitives.Block) evmc.HostContext {

	return &Host{
		txContext:    tx,
		blockContext: block,
	}

}
