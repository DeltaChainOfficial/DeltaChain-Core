---
order: 3
---

# Application Architecture Guide

Here we provide a brief guide on the recommended architecture of a
DeltaChain blockchain application.

The following diagram provides a superb example:

![DeltaChain-DeltaChain-stack](../imgs/DeltaChain-DeltaChain-stack-4k.jpg)

We distinguish here between two forms of "application". The first is the
end-user application, like a desktop-based wallet app that a user downloads,
which is where the user actually interacts with the system. The other is the
ABCI application, which is the logic that actually runs on the blockchain.
Transactions sent by an end-user application are ultimately processed by the ABCI
application after being committed by the DeltaChain consensus.

The end-user application in this diagram is the [Lunie](https://lunie.io/) app, located at the bottom
left. Lunie communicates with a REST API exposed by the application.
The application with DeltaChain nodes and verifies DeltaChain light-client proofs
through the DeltaChain Core RPC. The DeltaChain Core process communicates with
a local ABCI application, where the user query or transaction is actually
processed.

The ABCI application must be a deterministic result of the DeltaChain
consensus - any external influence on the application state that didn't
come through DeltaChain could cause a consensus failure. Thus _nothing_
should communicate with the ABCI application except DeltaChain via ABCI.

If the ABCI application is written in Go, it can be compiled into the
DeltaChain binary. Otherwise, it should use a unix socket to communicate
with DeltaChain. If it's necessary to use TCP, extra care must be taken
to encrypt and authenticate the connection.

All reads from the ABCI application happen through the DeltaChain `/abci_query`
endpoint. All writes to the ABCI application happen through the DeltaChain
`/broadcast_tx_*` endpoints.

The Light-Client Daemon is what provides light clients (end users) with
nearly all the security of a full node. It formats and broadcasts
transactions, and verifies proofs of queries and transaction results.
Note that it need not be a daemon - the Light-Client logic could instead
be implemented in the same process as the end-user application.

Note for those ABCI applications with weaker security requirements, the
functionality of the Light-Client Daemon can be moved into the ABCI
application process itself. That said, exposing the ABCI application process
to anything besides DeltaChain over ABCI requires extreme caution, as
all transactions, and possibly all queries, should still pass through
DeltaChain.

See the following for more extensive documentation:

- [Interchain Standard for the Light-Client REST API](https://github.com/DeltaChain/DeltaChain-sdk/pull/1028)
- [DeltaChain RPC Docs](https://docs.DeltaChain.com/v0.34/rpc/)
- [DeltaChain in Production](../DeltaChain-core/running-in-production.md)
- [ABCI spec](https://github.com/DeltaChain/spec/tree/95cf253b6df623066ff7cd4074a94e7a3f147c7a/spec/abci)
