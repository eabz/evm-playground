package vm

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/eabz/evm-playground/csmt"
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"sync"
)

type MockState struct {
	balances map[[20]byte]uint64
	nonces   map[[20]byte]uint64
	lock     sync.Mutex
}

type Host struct {
	tree      csmt.Tree
	db        *leveldb.DB
	mockState *MockState
}

// AccountExists checks if a specif contract account exits withint the bytecode storage
func (h Host) AccountExists(addr evmc.Address) bool {
	fmt.Println(addr)
	return true
}

func (h Host) GetStorage(addr evmc.Address, key evmc.Hash) evmc.Hash {
	fmt.Println(addr, key)
	return evmc.Hash{}
}

func (h Host) SetStorage(addr evmc.Address, key evmc.Hash, value evmc.Hash) evmc.StorageStatus {
	fmt.Println(addr, key, value)
	return evmc.StorageAdded
}

func (h Host) GetBalance(addr evmc.Address) evmc.Hash {
	fmt.Println(addr)
	return evmc.Hash{}
}

func (h Host) GetCodeSize(addr evmc.Address) int {
	fmt.Println(addr)
	return 0
}

func (h Host) GetCodeHash(addr evmc.Address) evmc.Hash {
	fmt.Println(addr)
	return evmc.Hash{}
}

func (h Host) GetCode(addr evmc.Address) []byte {
	fmt.Println(addr)
	return []byte{}
}

func (h Host) Selfdestruct(addr evmc.Address, beneficiary evmc.Address) {
	fmt.Println(addr, beneficiary)
	return
}

func (h Host) GetTxContext() evmc.TxContext {
	return evmc.TxContext{}
}

func (h Host) GetBlockHash(number int64) evmc.Hash {
	fmt.Println(number)
	return evmc.Hash{}
}

func (h Host) EmitLog(addr evmc.Address, topics []evmc.Hash, data []byte) {
	fmt.Println(addr, topics, data)
}

func (h Host) Call(kind evmc.CallKind, destination evmc.Address, sender evmc.Address, value evmc.Hash, input []byte, gas int64, depth int, static bool, salt evmc.Hash) (output []byte, gasLeft int64, createAddr evmc.Address, err error) {
	fmt.Println(kind, destination, sender, value, input, gas, depth, static, salt)
	return nil, 0, [20]byte{}, err
}

func NewHost() (evmc.HostContext, error) {
	levelopts := &opt.Options{
		ErrorIfExist:           false,
		Strict:                 opt.DefaultStrict,
		Compression:            opt.SnappyCompression,
		Filter:                 filter.NewBloomFilter(10),
		DisableSeeksCompaction: true,
	}

	db, err := leveldb.OpenFile("./data/bytecodes", levelopts)
	if err != nil {
		return nil, err
	}

	tree, err := badger.Open(badger.DefaultOptions("./data/tree").WithLogger(nil))
	if err != nil {
		return nil, err
	}

	bdb := csmt.NewBadgerTreeDB(tree)

	return &Host{
		db:   db,
		tree: csmt.NewTree(bdb),
	}, nil

}
