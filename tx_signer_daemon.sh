#!/bin/bash

# Automated Transaction Signer Daemon Guide
# -----------------------------------------
# Description:
# This script is designed to automate sending transactions on Re Protocol using the `red tx oracle` command.
# It uses the `expect` utility to automate passphrase input when the command prompts for it.
#
# Prerequisites:
# 1. Ensure you have the `expect` tool installed.
# 2. Ensure you have the necessary permissions to execute scripts.
#
# Steps to Run:
# 1. Set Up the Script:
#    - Replace `YOUR_PASSPHRASE_HERE` with your actual passphrase.
# 2. Make the Script Executable:
#    chmod +x tx_signer_daemon.sh
# 3. Run the Script:
#    ./tx_signer_daemon.sh
#
# Notes:
# - Be cautious with hardcoded passphrases in scripts.
# - Always ensure you understand the behavior and consequences of automated scripts.
# -----------------------------------------

# Ensure the script exits on errors
set -e

PASSPHRASE="YOUR_PASSPHRASE_HERE"

while true; do
    # Use expect to automatically provide the passphrase for the first command
    expect << EOF
    spawn red tx oracle cosmoshub-txs <address> --from alice --chain-id test --keyring-backend test --gas 12899100
    expect "Enter keyring passphrase:"
    send "${PASSPHRASE}\r"
    expect eof
EOF

    # Sleep briefly between commands
    sleep 5

    # Use expect to automatically provide the passphrase for the second command
    expect << EOF
    spawn red tx oracle bitcoin-txs <address> --from bob --chain-id test --keyring-backend test --gas 12899100
    expect "Enter keyring passphrase:"
    send "${PASSPHRASE}\r"
    expect eof
EOF

    # Sleep before repeating the loop
    sleep 5
done

