# Interaction with the network: CLI

For more information on the command usage, refer to its help screen: `red --help`.

Here is a list of useful `red` commands, including usage examples.

## Keys

### Generate Keys

You'll need an account private and public key pair \(a.k.a. `sk, pk` respectively\) to be able to receive funds, send txs, bond tx, etc.

To generate a new _secp256k1_ key:

```bash
red keys add <account_name>
```

Previously, you had to enter a password to save it to disk, but you do not currently need to.

[Setting up the keyring](https://docs.cosmos.network/v0.42/run-node/keyring.html)

If you check your private keys, you'll now see `<account_name>`:

```bash
red keys show <account_name>
```

View the validator operator's address via:

```bash
red keys show <account_name> --bech=val
```

You can see all your available keys by typing:

```bash
red keys list
```

View the validator pubkey for your node by typing:

```bash
red tendermint show-validator
```

## Bank

### Query balance

After receiving tokens to your address, you can view your account's balance by typing:

```bash
red query bank balances <account_address>
```

## Send Tokens

The following command could be used to send coins from one account to another:

```bash
red tx bank send \
  <from_key_or_address> \
  <to_address> \
  <amount> \
  --chain-id <chain_id> \
````

Now, view the updated balances of the origin and destination accounts:

```bash
red query bank balances <from_address>
red query bank balances <to_address>
```

You can also check your balance at a given block by using the `--height` flag:

```bash
red query account <account_address> --height=<block_height>
```

You can simulate a transaction without actually broadcasting it by appending the `--dry-run` flag to the command line:

```bash
red tx bank send \
  <from_key_or_address> \
  <to_address> \
  <amount> \
  --chain-id <chain_id> \
  --dry-run
```

Furthermore, you can build a transaction and print its JSON format to STDOUT by appending `--generate-only` to the list of the command line arguments:

```bash
red tx bank send \
  <from_key_or_address> \
  <to_address> \
  <amount> \
  --chain-id <chain_id> \
  --generate-only > unsignedSendTx.json
```

```bash
red tx sign \
  --chain-id <chain_id> \
  --from <key_name> \
  unsignedSendTx.json > signedSendTx.json
```


## DID

### Create(Issue) a DID

```bash
redd tx did create-did \
  --chain-id <chain-id> \
  --from <address>
```
This doesn't require any parameter except `chain-id` and `from`.
That is, it generates a Secp256k1 key-pair and derive a DID and a DID Document.
The DID Document is stored in Re.

To store the key-pair safely in your local, the command will prompt you to enter a passphrase.
The encrypted key-pair file will be stored in your `~/did_keystore` directory.

### Resolve a DID

```bash
red query did get-did <did>
```
This returns a DID Document in JSON corresponding to that DID.
If the DID doesn't exist, or was already deactivated, an error will be returned.

### Update a DID

```bash
red tx did update-did <did> <key-id> <did-doc-path> \
  --chain-id <chain-id> \
  --from <address>
```
A DID Document will be replaced to the new one written in a JSON file: `<did-doc-path>`.
To prove that you are the DID owner, you must pass a `<key-id>` that is one of `verificationMethod`s in the DID Document.
Also, that command will prompt you to enter a passphrase of that key if the key-pair is stored in your keystore: `~/did_keystore`.
The key-pair will be used to make a signature so that Re can verify that you are the DID owner.

### Deactivate a DID

```bash
red tx did deactivate-did <did> <key-id> \
  --chain-id <chain-id> \
  --from <address>
```

Like updating a DID, a `<key-id>` must be specified and the corresponding key-pair should be used to make a signature
so that Re can verify that you are the DID owner.


