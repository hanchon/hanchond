package txbuilder

import (
	"github.com/hanchon/hanchond/lib/requester"
)

type TxBuilder struct {
	contracts map[string]Contract
	mnemonic  string

	customGasLimit  map[string]uint64
	defaultGasLimit uint64

	currentNonce map[string]uint64

	requester *requester.Client
}

func NexTxBuilder(
	contracts map[string]Contract,
	mnemonic string,
	customGasLimit map[string]uint64,
	defaultGasLimit uint64,
	requester *requester.Client,
) *TxBuilder {
	return &TxBuilder{
		contracts:       contracts,
		mnemonic:        mnemonic,
		customGasLimit:  customGasLimit,
		defaultGasLimit: defaultGasLimit,
		currentNonce:    map[string]uint64{},

		requester: requester,
	}
}

func NewSimpleTxBuilder(mnemonic string, web3Endpoint string) *TxBuilder {
	return NexTxBuilder(
		map[string]Contract{},
		mnemonic,
		map[string]uint64{},
		100_000,
		requester.NewClient().WithUnsecureWeb3Endpoint(web3Endpoint),
	)
}
