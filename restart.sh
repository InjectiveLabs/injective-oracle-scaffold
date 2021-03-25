cd cmd/injectived
ulimit -n 65000
yes 12345678 | go run commands.go genaccounts.go gentx.go main.go root.go start.go util.go \
--log-level "main:info,state:info,statesync:info,*:error" \
--chain-id=888 start
