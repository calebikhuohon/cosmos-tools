# cosmos-tools

A CLI tool that generates CSV data with the specified Cosmos chain information and provides an upgrade analysis (incomplete) of the Osmosis chain.

## Prerequisites for running the CLI
* Go (min. version 1.19)
* CMake
* [golang CI Linter](https://github.com/golangci/golangci-lint)

## How to Run the Application
* Clone this repository
* Run `make install` to install the Validator Status and Vesting Account Analysis CLIs.
* Run `validator-status generate --chain=<chain name>` to generate status data for the passed in chain. The Cosmos chain e.g evmos, is passed in here.
* Run `vesting-account analyze` to generate a JSON file (`unlock-schedule.json`) containing Unlock Schedules for the various Cosmos vesting accounts in the Umee genesis file.
* Optional: Run `make lint` to lint the various Go files.

## Possible Improvements to the Application
* Addition of unit test in the `client` package, using [gock](https://gopkg.in/h2non/gock.v1) to mock HTTP requests, and [mockery](https://github.com/vektra/mockery) for the created I/O interfaces.
* Use of batch work processing to generate CSV data on multiple worker threads.
* Robust error logging with [Sentry](https://sentry.io/welcome/)
