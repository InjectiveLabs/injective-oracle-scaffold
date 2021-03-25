package metrics

func SidechainSubmitBatchSize(size int, tags ...Tags) {
	clientMux.RLock()
	defer clientMux.RUnlock()

	if client == nil {
		return
	}

	tagSpec := joinTags(tags...)

	client.Count("sidechain.submit_batch.size"+tagSpec, size)
}

func EVMRevertedTx(tags ...Tags) {
	clientMux.RLock()
	defer clientMux.RUnlock()

	if client == nil {
		return
	}

	tagSpec := joinTags(tags...)

	client.Count("evm.reverted_tx"+tagSpec, 1)
}

func EVMGasConsumed(used uint64, tags ...Tags) {
	clientMux.RLock()
	defer clientMux.RUnlock()

	if client == nil {
		return
	}

	tagSpec := joinTags(tags...)

	client.Count("evm.gas_consumed"+tagSpec, used)
}
