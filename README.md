# Injective Oracle Scaffold

Home of the following services:

* [injectived](/cmd/injectived)

## Building from sources

In order to build from sources you’ll need at least [Go 1.15+](https://golang.org/dl/).

```bash
# need to clone if you plan to run tests, and use Makefile
$ git clone git@github.com:InjectiveLabs/injective-oracle-scaffold.git
$ cd injective-oracle-scaffold
$ make install

# or simply do this to fetch modules and build executables
$ go install github.com/InjectiveLabs/injective-oracle-scaffold/cmd/...
```

## Injective Chain (Single Node)

### Quick Setup

The most convenient way to launch services is by running the setup script:

```bash
$ ./setup.sh
```

⚠️ This will remove any data from `~/.injective` and create a new state.

Then run an instance of the injectived node. 

```bash
$ ./injectived.sh
```

Voila! You have now successfully setup a full node on the Injective Chain. 

### Running an Injective Chain node

To start up a Injective Chain daemon node (Tendermint/Cosmos) which also has built-in gRPC server, run the following command. 

```bash
$ injectived \
	--chain-id=888 \
	--from=genesis \
	--from-passphrase="12345678" \
	start

```

You can also use script `./injectived.sh` and edit it.

## Injective Chain (Multinode Cluster)

Multinode launching script `multinode.sh` included in this repo can be used to create any Cosmos testnet, we'll be using it to create a testing Injective Chain, that will behave like a real deployment with active consensus.

### Bootstrap cluster of 3 nodes

```
$ CHAIN_ID=888 DENOM=inj ./multinode.sh injectived
```

Full list of the supported ENV variables:
* `CHAIN_ID` - specifies Cosmos Chain ID, like `888`
* `CHAIN_DIR` - is a prefix for all data dirs and logs, will be removed if `CLEANUP=1`. By default it's in `./data`
* `DENOM` - Cosmos coin denom, the default coin of the network. Examples: `uatom`, `aphoton`, `inj` etc
* `STAKE_DENOM` - Cosmos coin denom that is used for staking and governance. On the Cosmos Hub it's `stake`. Defaults to value of `DENOM` in the script.
* `SCALE_FACTOR` - Scale factor for the Cosmos coin. Defaults to 1e18 to reflect Ethereum token balances. Use `000000` to follow Cosmos uatom (1e6) style.
* `CLEANUP` - if this option set to `1`, then the `CHAIN_DIR` will be removed in the most unsafe manner.
* `LOG_LEVEL` - sets the log level of the Cosmos node configuration. Defaults to Cosmos' default (`main:info,state:info,statesync:info,*:error`).

**Important**: it is safe to run the script multiple times, it will stop nodes upon running, and optionally cleanup the state. If the state is not empty, the script will start nodes without running initialization again. So it could be used for manually retriable tests.

The nodes are running in background, easy way to access logs:

```
$ tail -f ./data/888.n0.log"
$ tail -f ./data/888.n1.log"
$ tail -f ./data/888.n2.log"
```

### Cosmos Accounts

The script imports 3 validator accounts and 1 user account, specified by mnemonics in the script itself. Each validator account accessible as `val` on the corresponding nodes, and user account is shared across all three nodes as `user`.

## Generating Injective Chain type bindings and REST and gRPC Gateway docs

First, ensure that the `Enable` and `Swagger` values are true in APIConfig set in `cmd/injectived/config/config.go`. 

Then simply run the following command to auto-generate the Swagger UI docs.

```
$ make gen
```

Then when you start the Injective Daemon, simply navigate to [http://localhost:10337/swagger/](http://localhost:10337/swagger/) for Swagger docs and REST (legacy) client.

We recommend to use GRPC endpoint instead, so get a GUI client like [BloomRPC](https://github.com/uw-labs/bloomrpc) and connect it to `localhost:9900`. Import any proto files from `third_party/proto` or `proto/injective` to query services.

