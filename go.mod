module github.com/InjectiveLabs/injective-oracle-scaffold

go 1.15

require (
	github.com/InjectiveLabs/sdk-go v1.16.7
	github.com/alexcesaro/statsd v2.0.0+incompatible
	github.com/aristanetworks/goarista v0.0.0-20201012165903-2cb20defcd66 // indirect
	github.com/aws/aws-sdk-go v1.29.15 // indirect
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.42.3
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/ethereum/go-ethereum v1.9.25
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/improbable-eng/grpc-web v0.13.0
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/miguelmota/go-ethereum-hdwallet v0.0.0-20200123000308-a60dcd172b4c
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/onsi/ginkgo v1.15.1
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/cors v1.7.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
	github.com/tyler-smith/go-bip39 v1.0.2
	github.com/xlab/closer v0.0.0-20190328110542-03326addb7c2
	github.com/xlab/suplog v1.3.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210308170721-88b6017d0656 // indirect
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.35.0
	gopkg.in/alexcesaro/statsd.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
