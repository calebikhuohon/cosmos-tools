# Chain Upgrade Analysis 

A chain upgrade analysis for the cosmos-hub and Osmosis chains.

This report provides a detailed outlook on the last three chain upgrades in the aforementioned Cosmos chains

## Osmosis
We begin with an indepth look at Osmosis.

The last upgrade to the Osmosis chain, v12.0.0 (Oxygen upgrade), occurred at block height 6246000 on approximately September 30th. 

This upgrade featured the following improvements:
* Time weighted average prices for all AMM pools, exposing an on-chain oracle for each pool which allows CosmWasm apps to use this data.
* Addition of more CosmWasm contract developer features. Examples include the ability to enable queries for CosmWasm contracts, and upgrade to wasmd v0.28.0, and many more.
* Enablement of Interchain Accounts. ICA enabled chains can now carry out transactions with Osmosis and vice versa. This also enabled platforms like Quasar to manage assets/contracts on different chains from accounts on Osmosis, and vice-versa.
* Expedited proposals. This allows governance to have proposals that execute in faster time windows given higher quorums (2/3rds)

A complete change log can be found [here](https://raw.githubusercontent.com/osmosis-labs/osmosis/v12.x/CHANGELOG.md)

As a result of this change, chain information such as names of binary executables, schema files, and semantic versioning had to be updated as well to match the new upgrade version.

To achieve the upgrade, the following steps were taken:
* All validator nodes were to upgrade to v12 but keep their nodes offline until a defined time to allow for a coordinated restart of the network.
* All validator start their nodes at the predefined time
* Once 67% or more of the voting power was online, block 6246000 was reached, along with the upgrade at this height. 

After chain upgrade was achieved, the team noted that there was a [chain halt](https://link.medium.com/gKc4lUYdcvb ) for 5 hours on October 3rd due to a bug in the [EndBlocker](https://github.com/cosmos/cosmos-sdk/blob/214b11dcbaa129f7b4c0013b2103db9d54b85e9e/docs/docs/building-modules/05-beginblock-endblock.md) logic. This resulted in non-deterministic chain states which broke consensus. 

_______________


Further work on this report will include details relating to code commits made to facilitate the upgrade, possible security issues, areas for improvement, possible runtime benchmarks comparing upgrades, work on the previous two upgrades (v11 and v10), and work on the Cosmos-hub's last three upgrades.



