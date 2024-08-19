#!/usr/bin/env bash

BINARY=$(which red)
NODE_HOME_DIR=$HOME/.re
CHAIN_ID=re-local
DENOM=stake

function add_keys() {
    local name=$1
    local amount=$2
    $BINARY keys add $name --keyring-backend test
    $BINARY genesis add-genesis-account $name $amount --keyring-backend test
}

# reset the chain
rm -r ~/.re || true

# configure the chain
$BINARY config set client chain-id $CHAIN_ID
$BINARY config set client keyring-backend test

# initialize a single node local network
$BINARY init node --chain-id $CHAIN_ID --default-denom $DENOM

# add keys
add_keys "alice" "1000000000stake"
add_keys "bob" "1000000000stake"

# create default validator
$BINARY genesis gentx alice 1000000${DENOM} --chain-id $CHAIN_ID
if ! $BINARY genesis validate --trace; then
    echo "invalid genesis file"
    exit 1
fi

$BINARY genesis collect-gentxs

# modify minimum-gas-prices and API settings
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    sed -i '' "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0${DENOM}\"/" $NODE_HOME_DIR/config/app.toml
    sed -i '' '/^\[api\]/,/^\[/ s/enable = false/enable = true/' $NODE_HOME_DIR/config/app.toml
    sed -i '' '/^\[api\]/,/^\[/ s/swagger = false/swagger = true/' $NODE_HOME_DIR/config/app.toml
    sed -i '' '/^\[api\]/,/^\[/ s|address = "tcp://localhost:1317"|address = "tcp://0.0.0.0:1317"|' $NODE_HOME_DIR/config/app.toml
else
    # Linux
    sed -i "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0${DENOM}\"/" $NODE_HOME_DIR/config/app.toml
    sed -i '/^\[api\]/,/^\[/ s/enable = false/enable = true/' $NODE_HOME_DIR/config/app.toml
    sed -i '/^\[api\]/,/^\[/ s/swagger = false/swagger = true/' $NODE_HOME_DIR/config/app.toml
    sed -i '/^\[api\]/,/^\[/ s|address = "tcp://localhost:1317"|address = "tcp://0.0.0.0:1317"|' $NODE_HOME_DIR/config/app.toml
fi

# start lcnd    
$BINARY start || echo "Failed to start the chain"