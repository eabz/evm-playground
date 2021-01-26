package vm

import (
	"github.com/dgraph-io/badger"
	"github.com/eabz/evm-playground/csmt"
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
)

type Host struct {
	db csmt.Tree
}

func (h Host) AccountExists(addr evmc.Address) bool {
	panic("implement me AccountExists")
}

func (h Host) GetStorage(addr evmc.Address, key evmc.Hash) evmc.Hash {
	panic("implement me GetStorage")
}

func (h Host) SetStorage(addr evmc.Address, key evmc.Hash, value evmc.Hash) evmc.StorageStatus {
	panic("implement me SetStorage")
}

func (h Host) GetBalance(addr evmc.Address) evmc.Hash {
	panic("implement me GetBalance")
}

func (h Host) GetCodeSize(addr evmc.Address) int {
	panic("implement me GetCodeSize")
}

func (h Host) GetCodeHash(addr evmc.Address) evmc.Hash {
	panic("implement me GetCodeHash")
}

func (h Host) GetCode(addr evmc.Address) []byte {
	panic("implement me GetCode")
}

func (h Host) Selfdestruct(addr evmc.Address, beneficiary evmc.Address) {
	panic("implement me Selfdestruct")
}

func (h Host) GetTxContext() evmc.TxContext {
	panic("implement me GetTxContext")
}

func (h Host) GetBlockHash(number int64) evmc.Hash {
	panic("implement me GetBlockHash")
}

func (h Host) EmitLog(addr evmc.Address, topics []evmc.Hash, data []byte) {
	panic("implement me EmitLog")
}

func (h Host) Call(kind evmc.CallKind, destination evmc.Address, sender evmc.Address, value evmc.Hash, input []byte, gas int64, depth int, static bool, salt evmc.Hash) (output []byte, gasLeft int64, createAddr evmc.Address, err error) {
	panic("implement me Call")
}

func NewHost() (evmc.HostContext, error) {
	opts := badger.Options{
		Dir: "./database",
	}
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	bdb := csmt.NewBadgerTreeDB(db)
	return &Host{
		db: csmt.NewTree(bdb),
	}, nil
}
