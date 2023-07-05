# Overview

See below for an overview of tendermint.

- [What is tendermint](https://github.com/PikeEcosystem/tendermint/blob/main/docs/en/01-overview.md)
- [Consensus](https://github.com/PikeEcosystem/tendermint/blob/main/docs/en/02-consensus.md)
- [Transaction Sharing](https://github.com/PikeEcosystem/tendermint/blob/main/docs/en/03-tx-sharing.md)

## Optimization

tendermint has the following optimizations to improve performance:

- Fixed each reactor to process messages asynchronously in separate threads.
  - https://github.com/PikeEcosystem/tendermint/issues/128
  - https://github.com/PikeEcosystem/tendermint/pull/135
- Fixed some ABCI methdos to be executed concurrently.
  - https://github.com/PikeEcosystem/tendermint/pull/160
  - https://github.com/PikeEcosystem/tendermint/pull/163
