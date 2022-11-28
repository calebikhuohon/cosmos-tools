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

In order to enable a smooth upgrade experience, a [permissions update](https://github.com/osmosis-labs/osmosis/commit/286027916efa89a96a83c073c802b8857c3679f6) was made in the Osmosis daemon for all nodes

```shell
- dasel put string -f $GENESIS '.app_state.wasm.params.code_upload_access.permission' "Nobody"
+ dasel put string -f $GENESIS '.app_state.wasm.params.code_upload_access.permission' "Everybody"
    ...
    ...
    ...
- osmosisd add-genesis-account osmo12smx2wdlyttvyzvzg54y2vnqwq2qjateuf7thj 100000000000uosmo,100000000000uion --home $OSMOSIS_HOME     
+ osmosisd add-genesis-account osmo12smx2wdlyttvyzvzg54y2vnqwq2qjateuf7thj 100000000000uosmo,100000000000uion,100000000000stake --home $OSMOSIS_HOME 
```
From the above code snippet, we see that the wasm permissions were updated, and extra denoms were added for genesis accounts.

Some other changes outside features and fixes made to facilitate the upgrade included a [testnetify refactor and localOsmosis Dockerfile](https://github.com/osmosis-labs/osmosis/commit/bf9de2d9c18e6b269a345632f90720b48cdc9419) update, inclusion of the Osmosis [pool query](https://github.com/osmosis-labs/osmosis/commit/2487690324265496850c132b76a23bfc6d559065), amongst others. The details of all changes made for the Osmosis Oxygen (v12.0.0) upgrade can be found [here](https://github.com/osmosis-labs/osmosis/compare/v11.0.1...v12.0.0).

After chain upgrade was achieved, the team noted that there was a [chain halt](https://link.medium.com/gKc4lUYdcvb ) for 5 hours on October 3rd due to a bug in the [EndBlocker](https://github.com/cosmos/cosmos-sdk/blob/214b11dcbaa129f7b4c0013b2103db9d54b85e9e/docs/docs/building-modules/05-beginblock-endblock.md) logic. This resulted in non-deterministic chain states which broke consensus. 

Some possible improvements would be a revival of IBC client to Dig after the expiration caused by the Chain halt. This is currently being [voted on](https://wallet.keplr.app/chains/osmosis/proposals/367).


## Conclusion

So far, we have covered the the Osmosis V12.0.0 upgrade, specific steps required to enable it for all nodes, upgrade-specific code changes, the chain halt bug, and a possible improvement currently in consideration from the chain governance. Further work on the Osmosis v11.0.0 and v10.0.0, and Cosmos-Hub upgrades will follow similar pattern, including possible benchmarks between upgrades and some notes on recorded malicious attacks.




