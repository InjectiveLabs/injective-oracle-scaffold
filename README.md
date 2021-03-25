# Injective Oracle Scaffold

Home of the following services:

* [injectived](/cmd/injectived)
* [injective-api](/cmd/injective-api)

## Architecture

<img alt="diagram-injective.png" src="https://docs.injectiveprotocol.com/images/ContractArchitecture-03726ae4.png" width="700px"/>

## Installation

### Building from sources

In order to build from sources you’ll need at least [Go 1.13+](https://golang.org/dl/).

Install solc version manager so you can do "solc use 0.6.7" and another day "solc use 0.7.1", yet we try to stick to ^0.6.x everywhere for now.

```bash
# need to clone if you plan to run tests, and use Makefile
$ git clone git@github.com:InjectiveLabs/injective-oracle-scaffold.git
$ cd injective-oracle-scaffold
$ make install

# or simply do this to fetch modules and build executables
$ go install github.com/InjectiveLabs/injective-oracle-scaffold/cmd/...
```
### Quick Setup 
The most convenient way to launch services is by running the setup script:
```bash
$ ./setup.sh
```
Then run an instance of the injectived node. 
```bash
$ ./injectived.sh
```
Open a new terminal tab and run the injective-cli EVM RPC node. 
```bash
# Make sure to edit the flag to the most recent deployment
$ ./injective-cli.sh
```
The Injective Chain services expect the [0x exchange contract suite](https://github.com/InjectiveLabs/0x-exchange-omnibus), [Injective Coordinator Contract](https://github.com/InjectiveLabs/injective-coordinator-contract), and the [Injective Derivatives Protocol](https://github.com/InjectiveLabs/injective-futures) contracts to be deployed before all of the services can run.  

If these contracts haven't been deployed yet (e.g. if this is the first time you're setting up the Injective Chain locally), run the following steps: 

```bash
./deploy_contracts.sh
```

Now, the injective-exchange service can be run. 
```bash
./injective-exchange.sh
```

Voila! You have now successfully setup a full node on the Injective Chain. 

### Using Docker

Another convenient way to launch services is by using the provided Docker container:

```bash
$ docker run --rm docker.injective.dev/core injectived -h
$ docker run --rm docker.injective.dev/core injective-api -h
```

### Setting up a node

```bash
$ injectived init [your_custom_moniker] --chain-id testnet
```
The home directory of this node is `~/.injectived` by default.

If you receive an error that `genesis.json` already exists in your `~/.injectived/config/genesis.json` then you can either delete your `./injectived` directory and run the init command again or proceed.  

You can also edit this `moniker` later, in the `~/.injectived/config/config.toml` file:

```toml
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

## Running an Injective Chain node

To start up a Injective Chain daemon node (Tendermint/Cosmos) which also has built-in gRPC server, run the following command. 

Note: this will cause the node to attempt to connect with the [0x exchange contract suite](https://github.com/InjectiveLabs/0x-exchange-omnibus), [Injective Coordinator Contract](https://github.com/InjectiveLabs/injective-coordinator-contract), and the [Injective Derivatives Protocol](https://github.com/InjectiveLabs/injective-futures) contracts at your local Injective EVM instance at `localhost:1317`.      

```bash
$ injectived \
	--chain-id=888 \
	--from=genesis \
	--from-passphrase="12345678" \
	--eth-node-http="http://localhost:1317" \
	--eth-node-ws="http://localhost:1318" \
	start

# Or use those for a private keyfile:
#	--eth-from="0x0" \
#	--eth-from-passphrase="1234" \
```

You can also use script `./injectived.sh` and edit it. By default injectived starts with these options:

* `--chain-id` specifies Cosmos chain id and must match the name set at genesis stage;
* `--from` sets the name of wallet that will be used to sign Cosmos transactions;
* `--eth-node-http` is the Ethereum network node HTTP RPC API endpoint;
* `--eth-node-ws` is the Ethereum network node WS RPC API endpoint;

You can also launch the node in read-only mode with the following:
```bash
$ injectived --chain-id=888 start
```

## Generating REST and gRPC Gateway docs
First, ensure that the `Enable` and `Swagger` values are true in APIConfig set in `cmd/injectived/config/config.go`. 

Then simply run the following command to auto-generate the Swagger UI docs.
```
$ make proto-swagger-gen
```
Then when you start the Injective Daemon, simply navigate to [http://localhost:10337/swagger/](http://localhost:10337/swagger/).  

## Generating Injective Chain API gRPC Typescript bindings

```
$ make gen
```
Then when you start the Injective Daemon, simply navigate to [http://localhost:10337/swagger/](http://localhost:10337/swagger/).  

## Exposing Relayer REST API

By default `injectived` RPC server listens on `grpc://localhost:9900`.

Notice that by default this HTTP server listens on `http://localhost:4444`. When Injective API server starts, it connects to the Injective daemon using gRPC protocol, you should see no warnings about connectivity and also in `injectived` logs the following lines are mandatory:

```
I[2019-10-19|18:11:04.383] [RPC]                                        module=main id=UZuQBaJU method=/injective_daemon.InjectiveDaemon/Ping bytes=0
I[2019-10-19|18:11:04.383] [RPC]                                        module=main id=UZuQBaJU status=OK bytes=0 time=97.233µs
```


<!--## Running a validator that submits to Ethereum -->

<!--In order to report to Ethereum, a validator must publish itself as an online peer. To do that, a key needs to be specified and `chain-id` to send Tednermint Txs. After start, validator starts to send MsgPing and will shutdown with a MsgLogOff tx. That will allow to participate in leaders selection to submit current trades into Ethereum.-->

<!--All encrypted keyfiles must be placed into `~/.injectived/eth/keystore` in order to be found. The name of the file doesn't matter. -->

<!--**Example of new private keyfile creation**:-->

<!--geth account new --keystore ~/.injectived/eth/keystore/ -->

```
injective-exchange \
	--injectived-rpc-addr="localhost:9900" \
	--evm-node-http="http://localhost:1317" \
	--evm-node-ws="ws://localhost:1318" \
```

## Maintenance

To run all unit tests:

```bash
$ make test-unit
```

To check the coverage:

```bash
$ make cover

# opens browser with coverage report
```
