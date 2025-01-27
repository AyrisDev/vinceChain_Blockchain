<!--
parent:
  order: false
-->

<div align="center">
  <img src="https://user-images.githubusercontent.com/102999403/161656793-7a826432-de47-46ea-b212-72907f462b7d.gif" />
  <h1> vince </h1>
</div>

<!-- TODO: add banner -->
<!-- ![banner](docs/ethermint.jpg) -->

<div align="center">
  <a href="https://github.com/AyrisDev/vinceChain_Blockchain/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/vincefoundation/vince.svg" />
  </a>
  <a href="https://github.com/AyrisDev/vinceChain_Blockchain/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/vincefoundation/vince.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/AyrisDev/vinceChain_Blockchain">
    <img alt="GoDoc" src="https://godoc.org/github.com/AyrisDev/vinceChain_Blockchain?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/AyrisDev/vinceChain_Blockchain">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/AyrisDev/vinceChain_Blockchain"/>
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/ArXNfK99ae">
    <img alt="Discord" src="https://img.shields.io/discord/962917488180490250.svg" />
  </a>
  <a href="https://github.com/AyrisDev/vinceChain_Blockchain/actions?query=branch%3Amain+workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/AyrisDev/vinceChain_Blockchain/actions/workflows/lint.yml/badge.svg?branch=main" />
  </a>
  <a href="https://codecov.io/gh/vincefoundation/vince">
    <img alt="Code Coverage" src="https://codecov.io/gh/vincefoundation/vince/branch/main/graph/badge.svg" />
  </a>
  <a href="https://twitter.com/vinceFDN">
    <img alt="Twitter Follow vince" src="https://img.shields.io/twitter/follow/vinceFDN"/>
  </a>
</div>

vince (ECH) is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and
interoperable with Ethereum and Cosmos. It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine with [EVMOS Ethermint](https://github.com/evmos/ethermint).

**Note**: Requires [Go 1.18+](https://golang.org/dl/)

## Installation

For prerequisites and detailed build instructions please read the [Installation](https://docs.ech.network) instructions. Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/AyrisDev/vinceChain_Blockchain/releases).

## Genesis

To get onto our mainnet (vince_5000-1) genesis download the genesis.json here

`wget https://gist.githubusercontent.com/vincefoundation/ee862f58850fc1b5ee6a6fdccc3130d2/raw/55c2c4ea2fee8a9391d0dc55b2c272adb804054a/genesis.json`

and then move it into the vinced config (after you have initilized your node)

`vinced init <nodename> --chain-id vince_5000-1`

`mv genesis.json ~/.vinced/config/`

## Quick Start

To learn how the vince works from a high-level perspective, go to the [Introduction](https://docs.ech.network) section from the documentation. You can also check the instructions to [Run a Node](https://docs.ech.network). You can also read the Cosmos SDK and familiarize yourself with Cosmos SDK, Tendermint, and Ethermint.

## Community

The following chat channels and forums are a great spot to ask questions about vince:

- [vince Twitter](https://twitter.com/vinceFDN)
- [vince Discord](https://discord.gg/ArXNfK99ae)
- [vince Telegram](https://t.me/vinceANN)
- [Official Website](https://ech.network)
- [Official Dapp](https://app.ech.network)

## Contributing

Looking for a good place to start contributing? Check out some [`good first issues`](https://github.com/AyrisDev/vinceChain_Blockchain/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

Original codebase forked from [EVMOS](https://github.com/evmos/evmos)

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.
