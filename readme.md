# ZetaChain

ZetaChain is an EVM-compatible L1 blockchain that enables omnichain, generic
smart contracts and messaging between any blockchain.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.20
- [Docker](https://docs.docker.com/install/) and
  [Docker Compose](https://docs.docker.com/compose/install/) (optional, for
  running tests locally)
- [buf](https://buf.build/) (optional, for processing protocol buffer files)
- [jq](https://stedolan.github.io/jq/download/) (optional, for running scripts)

## Components of ZetaChain

ZetaChain is built with [Cosmos SDK](https://github.com/cosmos/cosmos-sdk), a
modular framework for building blockchain and
[Ethermint](https://github.com/evmos/ethermint), a module that implements
EVM-compatibility.

- [zeta-node](https://github.com/zeta-chain/zeta-node) (this repository)
  contains the source code for the ZetaChain node (`zetacored`) and the
  ZetaChain client (`zetaclientd`).
- [protocol-contracts](https://github.com/zeta-chain/protocol-contracts)
  contains the source code for the Solidity smart contracts that implement the
  core functionality of ZetaChain.

## Building the `zetacored`/`zetaclientd` binaries

Clone this repository, checkout the latest release tag, and type the following command to build the binaries:
```
make install
```
to build. 

This command will install the `zetacored` and `zetaclientd` binaries in your
`$GOPATH/bin` directory.

Verify that the version of the binaries match the release tag.  
```
zetacored version
zetaclientd version
```

## Making changes to the source code

After making changes to any of the protocol buffer files, run the following
command to run generated files generation (ProtoBuf, OpenAPI and docs):

```
make generate
```

This command will use `buf` to generate the Go files from the protocol buffer
files and move them to the correct directories inside `x/`. It will also
generate an OpenAPI spec.

This command will run a script to update the modules' documentation. The script
uses static code analysis to read the protocol buffer files and identify all
Cosmos SDK messages. It then searches the source code for the corresponding
message handler functions and retrieves the documentation for those functions.
Finally, it creates a `messages.md` file for each module, which contains the
documentation for all the messages in that module.

## Run ZetaChain localnet with Optimism devnet
- Clone branch `integrate/optimism`
  - `git clone -b integrate/optimism https://github.com/royki/node.git`
- `cd node`
- Start zetachain localnet and optimism devnet
  - `make start-zeta-localnet-with-op`
- Start monitoring
  - `make start-monitoring`
- Stop monitoring
  - `make stop-monitoring`
- Stop zetachain localnet and optimism devnet
  - `make stop-zeta-localnet-with-op`
- Start e2e tests
  - `make start-e2e-test-with-op`
- Stop/remove zetachain localnet and optimism devnet
  - `make stop-e2e-test-with-op`

## Further Reading

Find below further documentation for development and running your own ZetaChain node:

* [Run the E2E tests and interact with the localnet](docs/development/LOCAL_TESTING.md)
* [Make a new ZetaChain release](docs/development/RELEASES.md)
* [Deploy your own ZetaChain or Bitcoin node](docs/development/DEPLOY_NODES.md)


## Community

[Twitter](https://twitter.com/zetablockchain) |
[Discord](https://discord.com/invite/zetachain) |
[Telegram](https://t.me/zetachainofficial) | [Website](https://zetachain.com)


