package utils

import (
	"github.com/ethereum/evmc/v7/bindings/go/evmc"
	"golang.org/x/crypto/blake2b"
	"math/big"
)

var wei = big.NewInt(1e18)

func StringToHash(s string) evmc.Hash {
	hash := blake2b.Sum256([]byte(s))
	return hash
}

func NumToWei(number int64) string {
	b := big.NewInt(number)
	b.Mul(b, wei)
	return b.String()
}