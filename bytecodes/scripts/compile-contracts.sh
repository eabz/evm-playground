#!/bin/bash

./bin/solc -o ./erc20/compiled --bin-runtime --ast-json --asm --bin ./erc20/base/ERC20.sol