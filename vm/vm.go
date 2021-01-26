package vm

import (
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
)

type VM struct {
	instance *evmc.VM
	host     evmc.HostContext
}

func (v *VM) Version() string {
	return v.instance.Version()
}

func (v *VM) Name() string {
	return v.instance.Name()
}

func (v *VM) Execute(rev evmc.Revision,
	kind evmc.CallKind, static bool, depth int, gas int64,
	destination evmc.Address, sender evmc.Address, input []byte, value evmc.Hash,
	code []byte, create2Salt evmc.Hash) (output []byte, gasLeft int64, err error) {

	return v.instance.Execute(v.host, rev, kind, static, depth, gas, destination, sender, input, value, code, create2Salt)
}

func NewVM() (*VM, error) {
	h, err := NewHost()
	if err != nil {
		return nil, err
	}

	ins, err := evmc.LoadAndConfigure("./bin/evmone.dylib")
	if err != nil {
		return nil, err
	}

	vm := &VM{
		instance: ins,
		host:     h,
	}

	return vm, nil
}
