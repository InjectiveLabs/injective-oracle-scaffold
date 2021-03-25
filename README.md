# Injective Oracle Scaffold

Home of the following services:

* [injectived](/cmd/injectived)

### Building from sources

In order to build from sources you’ll need at least [Go 1.15+](https://golang.org/dl/).

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

Voila! You have now successfully setup a full node on the Injective Chain. 

## Running an Injective Chain node

To start up a Injective Chain daemon node (Tendermint/Cosmos) which also has built-in gRPC server, run the following command. 

```bash
$ injectived \
	--chain-id=888 \
	--from=genesis \
	--from-passphrase="12345678" \
	start

```

You can also use script `./injectived.sh` and edit it.

## Generating Injective Chain type bindings and REST and gRPC Gateway docs
First, ensure that the `Enable` and `Swagger` values are true in APIConfig set in `cmd/injectived/config/config.go`. 

Then simply run the following command to auto-generate the Swagger UI docs.
```
$ make gen
```
Then when you start the Injective Daemon, simply navigate to [http://localhost:10337/swagger/](http://localhost:10337/swagger/).  
