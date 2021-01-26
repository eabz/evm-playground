package main

import (
	"fmt"
	"github.com/eabz/evm-playground/vm"
)

func main() {
	m, err := vm.NewVM()
	if err != nil {
		panic(err)
	}

	ver := m.Version()
	name := m.Name()

	fmt.Println(ver)
	fmt.Println(name)

}
