#!/bin/bash
ulimit -n 120000
yes 12345678 | injectived --chain-id=888 \
--log-level "main:info,state:info,statesync:info,*:error" \
--rpc.laddr "tcp://0.0.0.0:26657" \
start