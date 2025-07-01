#!/usr/bin/env bash

HOMEDIR="$HOME/.cosevmd"
CHAINID="cosevm_929-1"

COSEVMD_BIN=$(which cosevmd)
if [ -z "$COSEVMD_BIN" ]; then
    GOBIN=$(go env GOPATH)/bin
    COSEVMD_BIN=$(which $GOBIN/cosevmd)
fi

if [ -z "$COSEVMD_BIN" ]; then
    echo "please verify cosevmd is installed"
    exit 1
fi

$COSEVMD_BIN start \
	--minimum-gas-prices=0.0001ucose \
	--home "$HOMEDIR" \
	--json-rpc.api eth,txpool,personal,net,debug,web3 \
	--chain-id "$CHAINID"