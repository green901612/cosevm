#!/usr/bin/env bash

rm -rf $HOME/.cosevmd
COSEVMD_BIN=$(which cosevmd)
if [ -z "$COSEVMD_BIN" ]; then
    GOBIN=$(go env GOPATH)/bin
    COSEVMD_BIN=$(which $GOBIN/cosevmd)
fi

if [ -z "$COSEVMD_BIN" ]; then
    echo "please verify cosevmd is installed"
    exit 1
fi

# configure cosevmd
$COSEVMD_BIN config set client chain-id local
$COSEVMD_BIN config set client keyring-backend test
$COSEVMD_BIN keys add alice
$COSEVMD_BIN keys add bob
$COSEVMD_BIN init test --chain-id local --default-denom cose
# update genesis
$COSEVMD_BIN genesis add-genesis-account alice 10000000cose --keyring-backend test
$COSEVMD_BIN genesis add-genesis-account bob 1000cose --keyring-backend test
# create default validator
$COSEVMD_BIN genesis gentx alice 1000000cose --chain-id local
$COSEVMD_BIN genesis collect-gentxs
