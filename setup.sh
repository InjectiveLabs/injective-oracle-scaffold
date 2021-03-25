#!/bin/bash

set -e

make install

killall injectived &>/dev/null || true
rm -rf ~/.injectived

CHAINID=888
MONIKER="injective"
PASSPHRASE="12345678"

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
injectived init $MONIKER --chain-id $CHAINID
perl -i -pe 's/^timeout_commit = ".*?"/timeout_commit = "1s"/' ~/.injectived/config/config.toml

cat $HOME/.injectived/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="inj"' > $HOME/.injectived/config/tmp_genesis.json && mv $HOME/.injectived/config/tmp_genesis.json $HOME/.injectived/config/genesis.json
cat $HOME/.injectived/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="inj"' > $HOME/.injectived/config/tmp_genesis.json && mv $HOME/.injectived/config/tmp_genesis.json $HOME/.injectived/config/genesis.json
cat $HOME/.injectived/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="inj"' > $HOME/.injectived/config/tmp_genesis.json && mv $HOME/.injectived/config/tmp_genesis.json $HOME/.injectived/config/genesis.json
echo "NOTE: Setting Governance Voting Period to 10 seconds for easy testing"
cat $HOME/.injectived/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="10s"' > $HOME/.injectived/config/tmp_genesis.json && mv $HOME/.injectived/config/tmp_genesis.json $HOME/.injectived/config/genesis.json
cat $HOME/.injectived/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="inj"' > $HOME/.injectived/config/tmp_genesis.json && mv $HOME/.injectived/config/tmp_genesis.json $HOME/.injectived/config/genesis.json

yes $PASSPHRASE | injectived keys add genesis
yes $PASSPHRASE | injectived add-genesis-account $(yes $PASSPHRASE | injectived keys show genesis -a) 1000000000000000000000000inj,1000000000000000000000000atom
# zero address account
yes $PASSPHRASE | injectived add-genesis-account inj1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqe2hm49 1inj

VAL_KEY="localkey"
VAL_MNEMONIC="gesture inject test cycle original hollow east ridge hen combine junk child bacon zero hope comfort vacuum milk pitch cage oppose unhappy lunar seat"

USER1_KEY="user1"
USER1_MNEMONIC="copper push brief egg scan entry inform record adjust fossil boss egg comic alien upon aspect dry avoid interest fury window hint race symptom"

USER2_KEY="user2"
USER2_MNEMONIC="maximum display century economy unlock van census kite error heart snow filter midnight usage egg venture cash kick motor survey drastic edge muffin visual"

USER3_KEY="user3"
USER3_MNEMONIC="keep liar demand upon shed essence tip undo eagle run people strong sense another salute double peasant egg royal hair report winner student diamond"

USER4_KEY="user4"
USER4_MNEMONIC="pony glide frown crisp unfold lawn cup loan trial govern usual matrix theory wash fresh address pioneer between meadow visa buffalo keep gallery swear"

NEWLINE=$'\n'

# Import keys from mnemonics
yes "$VAL_MNEMONIC$NEWLINE$PASSPHRASE" | injectived keys add $VAL_KEY --recover
yes "$USER1_MNEMONIC$NEWLINE$PASSPHRASE" | injectived keys add $USER1_KEY --recover
yes "$USER2_MNEMONIC$NEWLINE$PASSPHRASE" | injectived keys add $USER2_KEY --recover
yes "$USER3_MNEMONIC$NEWLINE$PASSPHRASE" | injectived keys add $USER3_KEY --recover
yes "$USER4_MNEMONIC$NEWLINE$PASSPHRASE" | injectived keys add $USER4_KEY --recover

# Allocate genesis accounts (cosmos formatted addresses)
yes $PASSPHRASE | injectived add-genesis-account $(injectived keys show $VAL_KEY -a) 1000000000000000000000inj
yes $PASSPHRASE | injectived add-genesis-account $(injectived keys show $USER1_KEY -a) 1000000000000000000000inj
yes $PASSPHRASE | injectived add-genesis-account $(injectived keys show $USER2_KEY -a) 1000000000000000000000inj
yes $PASSPHRASE | injectived add-genesis-account $(injectived keys show $USER3_KEY -a) 1000000000000000000000inj
yes $PASSPHRASE | injectived add-genesis-account $(injectived keys show $USER4_KEY -a) 1000000000000000000000inj

echo "Signing genesis transaction"
# Sign genesis transaction
yes $PASSPHRASE | injectived gentx genesis --chain-id $CHAINID --amount 1000000000000000000000inj

echo "Collecting genesis transaction"
# Collect genesis tx
yes $PASSPHRASE | injectived collect-gentxs

echo "Validating genesis"
# Run this to ensure everything worked and that the genesis file is setup correctly
injectived validate-genesis

echo "Setup done!"
