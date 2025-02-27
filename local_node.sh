#!/usr/bin/env bash

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
# TODO: exocore testnet chainid is still under consideration and need to be finalized later
CHAINID="exocoretestnet_233-1"
MONIKER="localtestnet"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
ALGO="eth_secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the exocored instance
HOMEDIR="$HOME/.tmp-exocored"
# to trace evm
#TRACE="--trace"
TRACE=""

# make the validator consensus key consistent
CONSENSUS_KEY_MNEMONIC="wonder quality resource ketchup occur stadium vicious output situate plug second monkey harbor vanish then myself primary feed earth story real soccer shove like"
# the account below acts as both initial operator and local consistent faucet.
# pk: D196DCA836F8AC2FFF45B3C9F0113825CCBB33FA1B39737B948503B263ED75AE
# 0x3e108c058e8066DA635321Dc3018294cA82ddEdf == exo18cggcpvwspnd5c6ny8wrqxpffj5zmhklprtnph
LOCAL_MNEMONIC="knock benefit magnet slogan normal broken frequent level video focus spell utility"
LOCAL_NAME="local_funded_account"

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# validate dependencies are installed
command -v jq >/dev/null 2>&1 || {
	echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
	exit 1
}

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
make install

# User prompt if an existing local node configuration is found.
if [ -d "$HOMEDIR" ]; then
	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
	echo "Overwrite the existing configuration and start a new local node? [y/n]"
	read -r overwrite
else
	overwrite="Y"
fi

# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# Remove the previous folder
	rm -rf "$HOMEDIR"

	# Set client config
	exocored config keyring-backend $KEYRING --home "$HOMEDIR"
	exocored config chain-id $CHAINID --home "$HOMEDIR"

	# If keys exist they should be deleted
	for KEY in "${KEYS[@]}"; do
		exocored keys add "$KEY" --keyring-backend $KEYRING --algo $ALGO --home "$HOMEDIR"
	done

	# Use recover so that there is always a consistent address funded in the localnet genesis.
	echo "${LOCAL_MNEMONIC}" | exocored --home "$HOMEDIR" --keyring-backend $KEYRING keys add "${LOCAL_NAME}" --recover

	# Set moniker and chain-id for Evmos (Moniker can be anything, chain-id must be an integer)
	# Use recover to use a consistent consensus key for validator.
	echo "${CONSENSUS_KEY_MNEMONIC}" | exocored init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR" --recover

	# Change parameter token denominations to aexo
	jq '.app_state["crisis"]["constant_fee"]["denom"]="aexo"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aexo"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	# When upgrade to cosmos-sdk v0.47, use gov.params to edit the deposit params
	jq '.app_state["gov"]["params"]["min_deposit"][0]["denom"]="aexo"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["evm"]["params"]["evm_denom"]="aexo"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Set gas limit in genesis
	jq '.consensus_params["block"]["max_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# x/assets
	jq '.app_state["assets"]["client_chains"][0]["name"]="Example EVM chain"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["client_chains"][0]["layer_zero_chain_id"]="101"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["tokens"][0]["asset_basic_info"]["address"]="0xdAC17F958D2ee523a2206206994597C13D831ec7"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["tokens"][0]["asset_basic_info"]["layer_zero_chain_id"]="101"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["tokens"][0]["asset_basic_info"]["total_supply"]="40022689732746729"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["tokens"][0]["staking_total_amount"]="0"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["deposits"][0]["staker"]="0x3e108c058e8066da635321dc3018294ca82ddedf_0x65"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["deposits"][0]["deposits"][0]["asset_id"]="0xdac17f958d2ee523a2206206994597c13d831ec7_0x65"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["deposits"][0]["deposits"][0]["info"]["total_deposit_amount"]="5000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["deposits"][0]["deposits"][0]["info"]["withdrawable_amount"]="5000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["assets"]["deposits"][0]["deposits"][0]["info"]["wait_unbonding_amount"]="0"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# x/operator
	jq '.app_state["operator"]["operators"][0]["earnings_addr"]="exo18cggcpvwspnd5c6ny8wrqxpffj5zmhklprtnph"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["operator"]["operators"][0]["operator_meta_info"]="operator1"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["operator"]["operators"][0]["commission"]["commission_rates"]["rate"]="0.0"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["operator"]["operators"][0]["commission"]["commission_rates"]["max_rate"]="0.0"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["operator"]["operators"][0]["commission"]["commission_rates"]["max_change_rate"]="0.0"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["operator"]["operator_records"][0]["operator_address"]="exo18cggcpvwspnd5c6ny8wrqxpffj5zmhklprtnph"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq ".app_state.operator.operator_records[0].chains[0].chain_id = \"${CHAINID}\"" "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq ".app_state.operator.operator_records[0].chains[0].consensus_key = \"0xf0f6919e522c5b97db2c8255bff743f9dfddd7ad9fc37cb0c1670b480d0f9914\"" "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# x/delegation
	jq '.app_state["delegation"]["delegations"][0]["staker_id"]="0x3e108c058e8066da635321dc3018294ca82ddedf_0x65"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["delegation"]["delegations"][0]["delegations"][0]["asset_id"]="0xdac17f958d2ee523a2206206994597c13d831ec7_0x65"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["delegation"]["delegations"][0]["delegations"][0]["per_operator_amounts"][0]["key"]="exo18cggcpvwspnd5c6ny8wrqxpffj5zmhklprtnph"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["delegation"]["delegations"][0]["delegations"][0]["per_operator_amounts"][0]["value"]["amount"]="5000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# x/dogfood
	jq '.app_state["dogfood"]["initial_val_set"][0]["public_key"]="0xf0f6919e522c5b97db2c8255bff743f9dfddd7ad9fc37cb0c1670b480d0f9914"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["dogfood"]["initial_val_set"][0]["power"]="5000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	if [[ $1 == "pending" ]]; then
		if [[ "$OSTYPE" == "darwin"* ]]; then
			sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
			sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
			sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
			sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$CONFIG"
			sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
		else
			sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
			sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
			sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
			sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$CONFIG"
			sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
		fi
	fi

	# remove evmos seeds for localnet
	if [[ "$OSTYPE" == "darwin"* ]]; then
		sed -i '' 's/seeds = "[^"]*"/seeds = ""/' "$CONFIG"
	else
		sed -i 's/seeds = "[^"]*"/seeds = ""/' "$CONFIG"
	fi

	# enable prometheus metrics
	if [[ "$OSTYPE" == "darwin"* ]]; then
		sed -i '' 's/prometheus = false/prometheus = true/' "$CONFIG"
		sed -i '' 's/prometheus-retention-time = 0/prometheus-retention-time  = 1000000000000/g' "$APP_TOML"
		sed -i '' 's/enabled = false/enabled = true/g' "$APP_TOML"
	else
		sed -i 's/prometheus = false/prometheus = true/' "$CONFIG"
		sed -i 's/prometheus-retention-time  = "0"/prometheus-retention-time  = "1000000000000"/g' "$APP_TOML"
		sed -i 's/enabled = false/enabled = true/g' "$APP_TOML"
	fi

	# Change proposal periods to pass within a reasonable time for local testing
	sed -i.bak 's/"max_deposit_period": "172800s"/"max_deposit_period": "30s"/g' "$HOMEDIR"/config/genesis.json
	sed -i.bak 's/"voting_period": "172800s"/"voting_period": "30s"/g' "$HOMEDIR"/config/genesis.json

	# set custom pruning settings
	sed -i.bak 's/pruning = "default"/pruning = "custom"/g' "$APP_TOML"
	sed -i.bak 's/pruning-keep-recent = "0"/pruning-keep-recent = "2"/g' "$APP_TOML"
	sed -i.bak 's/pruning-interval = "0"/pruning-interval = "10"/g' "$APP_TOML"

	# make sure the localhost IP is 0.0.0.0
	sed -i.bak 's/localhost/0.0.0.0/g' "$CONFIG"
	sed -i.bak 's/localhost/0.0.0.0/g' "$APP_TOML"
	sed -i.bak 's/127.0.0.1/0.0.0.0/g' "$APP_TOML"

	# Allocate genesis accounts (cosmos formatted addresses)
	for KEY in "${KEYS[@]}"; do
		exocored add-genesis-account "$KEY" 100000000000000000000000000aexo --keyring-backend $KEYRING --home "$HOMEDIR"
	done
	exocored add-genesis-account "${LOCAL_NAME}" 100000000000000000000000000aexo --keyring-backend $KEYRING --home "$HOMEDIR"

	# bc is required to add these big numbers
	# note the extra +1 is for LOCAL_NAME
	total_supply=$(echo "(${#KEYS[@]} + 1) * 100000000000000000000000000" | bc)
	jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	exocored validate-genesis --home "$HOMEDIR"

	if [[ $1 == "pending" ]]; then
		echo "pending mode is on, please wait for the first block committed."
	fi
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
exocored start --metrics "$TRACE" --log_level $LOGLEVEL --minimum-gas-prices=0.0001aexo --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --json-rpc.enable true --home "$HOMEDIR" --chain-id "$CHAINID"
