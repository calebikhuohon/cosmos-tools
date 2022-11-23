# cosmos-tools

A CLI tool that generates CSV data with the specified Cosmos chain information and provides an upgrade analysis (incomplete) of the Osmosis chain.

## Prerequisites for running the CLI
* Go (min. version 1.19)
* CMake
* [golang CI Linter](https://github.com/golangci/golangci-lint)

## How to Run the Application
* Clone this repository
* Run `make install`
* Run `validator-status generate --chain=<chain name>`. The Cosmos chain e.g evmos, is passed in here. 
* Optional: Run `make lint` to lint the various Go files.

## Possible Improvements to the Application
* Addition of unit test in the `client` package, using [gock](gopkg.in/h2non/gock.v1) to mock HTTP requests.
* Use of batch work processing to generate CSV data on multiple worker threads.
* Robust error logging with [Sentry](https://sentry.io/welcome/)
