#!/bin/bash

rm -r ./bin/evmone.dylib

git clone --recursive https://github.com/ethereum/evmone && cd evmone

cmake -S . -B build
cmake --build build --parallel

cd ..

mv evmone/build/lib/libevmone.0.6.0-dev.dylib ./bin/evmone.dylib

chmod +x ./bin/evmone.dylib
rm -rf evmone



