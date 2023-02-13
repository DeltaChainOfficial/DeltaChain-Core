# DeltaChain

_UPDATE: DeltaChainCore featureset is frozen for LTS, see issue https://github.com/DeltaChainOfficial <br/>
_This is the latest stable release used by DeltaChainhub-4, version 0.34.24_<br/>
_The previous main branch (v0.38.xx) can now be found under "main_backup"_<br/>

![banner](docs/DeltaChain-core-image.jpg)

[Byzantine-Fault Tolerant][bft] [State Machine Replication][smr]. Or
[Blockchain], for short.

[![Version][version-badge]][version-url]
[![API Reference][api-badge]][api-url]
[![Go version][go-badge]][go-url]
[![Discord chat][discord-badge]][discord-url]
[![License][license-badge]][license-url]
[![Sourcegraph][sg-badge]][sg-url]

| Branch | Tests                              | Linting                         |
|--------|------------------------------------|---------------------------------|
| main   | [![Tests][tests-badge]][tests-url] | [![Lint][lint-badge]][lint-url] |

DeltaChain Core is a Byzantine Fault Tolerant (BFT) middleware that takes a
state transition machine - written in any programming language - and securely
replicates it on many machines.

For protocol details, refer to the [DeltaChain Specification](./spec/README.md).

For detailed analysis of the consensus protocol, including safety and liveness
proofs, read our paper, "[The latest gossip on BFT
consensus](https://arxiv.org/abs/1807.04938)".

## Documentation

Complete documentation can be found on the
[website](https://docs.DeltaChain.com/).

## Releases

Please do not depend on `main` as your production branch. Use
[releases](https://github.com/DeltaChainOfficial) instead.

DeltaChain has been in the production of private and public environments, most
notably the blockchains of the DeltaChain Network. we haven't released v1.0 yet
since we are making breaking changes to the protocol and the APIs. See below for
more details about [versioning](#versioning).

In any case, if you intend to run DeltaChain in production, we're happy to help.
You can contact us [over email](contact@deltachain.network)

More on how releases are conducted can be found [here](./RELEASES.md).

## Security

To report a security vulnerability, see our [bug bounty
program](https://hackerone.com/DeltaChain). For examples of the kinds of bugs we're
looking for, see [our security policy](SECURITY.md).

We also maintain a dedicated mailing list for security updates. We will only
ever use this mailing list to notify you of vulnerabilities and fixes in
DeltaChain Core. You can subscribe [here](http://eepurl.com/gZ5hQD).

## Minimum requirements

| Requirement | Notes             |
|-------------|-------------------|
| Go version  | Go 1.18 or higher |

### Install

See the [install instructions](./docs/introduction/install.md).

### Quick Start

- [Single node](./docs/introduction/quick-start.md)
- [Local cluster using docker-compose](./docs/tools/docker-compose.md)
- [Remote cluster using Terraform and Ansible](./docs/tools/terraform-and-ansible.md)

## Contributing

Please abide by the [Code of Conduct](CODE_OF_CONDUCT.md) in all interactions.

Before contributing to the project, please take a look at the [contributing
guidelines](CONTRIBUTING.md) and the [style guide](STYLE_GUIDE.md). You may also
find it helpful to read the [specifications](./spec/README.md), and familiarize
yourself with our [Architectural Decision Records
(ADRs)](./docs/architecture/README.md) and
[Request For Comments (RFCs)](./docs/rfc/README.md).

## Versioning

### Semantic Versioning

DeltaChain uses [Semantic Versioning](http://semver.org/) to determine when and
how the version changes. According to SemVer, anything in the public API can
change at any time before version 1.0.0

To provide some stability to users of 0.X.X versions of DeltaChain, the MINOR
version is used to signal breaking changes across DeltaChain's API. This API
includes all publicly exposed types, functions, and methods in non-internal Go
packages as well as the types and methods accessible via the DeltaChain RPC
interface.

Breaking changes to these public APIs will be documented in the CHANGELOG.

### Upgrades

In an effort to avoid accumulating technical debt prior to 1.0.0, we do not
guarantee that breaking changes (ie. bumps in the MINOR version) will work with
existing DeltaChain blockchains. In these cases you will have to start a new
blockchain, or write something custom to get the old data into the new chain.
However, any bump in the PATCH version should be compatible with existing
blockchain histories.

For more information on upgrading, see [UPGRADING.md](./UPGRADING.md).

### Supported Versions

Because we are a small core team, we only ship patch updates, including security
updates, to the most recent minor release and the second-most recent minor
release. Consequently, we strongly recommend keeping DeltaChain up-to-date.
Upgrading instructions can be found in [UPGRADING.md](./UPGRADING.md).

## Resources

### Libraries

- [DeltaChain SDK](https://github.com/DeltaChainOfficial); A framework for building
  applications in Golang
- [DeltaChain in Rust](https://github.com/DeltaChainOfficial)
- [ABCI Tower](https://github.com/DeltaChainOfficial)