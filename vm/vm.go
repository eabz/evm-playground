package vm

import (
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
)

type VM struct {
	instance *evmc.VM
}

func (v *VM) Version() string {
	return v.instance.Version()
}

func (v *VM) Name() string {
	return v.instance.Name()
}

func (v *VM) Execute(ctx evmc.HostContext, rev evmc.Revision,
	kind evmc.CallKind, static bool, depth int, gas int64,
	destination evmc.Address, sender evmc.Address, input []byte, value evmc.Hash,
	code []byte, create2Salt evmc.Hash) (output []byte, gasLeft int64, err error) {
	return v.instance.Execute(ctx, rev, kind, static, depth, gas, destination, sender, input, value, code, create2Salt)
}

func NewVM() (*VM, error) {
	ins, err := evmc.LoadAndConfigure("./bin/evmone.dylib")
	if err != nil {
		return nil, err
	}

	vm := &VM{
		instance: ins,
	}

	return vm, nil
}
