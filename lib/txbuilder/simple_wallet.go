package txbuilder

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type SimpleWeb3Wallet struct {
	Mnemonic  string
	Address   common.Address
	PrivKey   *ecdsa.PrivateKey
	TxBuilder *TxBuilder
}

func NewSimpleWeb3WalletFromMnemonic(mnemonic string, web3Endpoint string) *SimpleWeb3Wallet {
	w, a, err := WalletFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}
	privKey, err := w.PrivateKey(a)
	if err != nil {
		panic(err)
	}

	return &SimpleWeb3Wallet{
		Mnemonic:  mnemonic,
		Address:   a.Address,
		PrivKey:   privKey,
		TxBuilder: NewSimpleTxBuilder(mnemonic, web3Endpoint),
	}
}

func NewSimpleWeb3Wallet(web3Endpoint string) *SimpleWeb3Wallet {
	mnemonic, err := NewMnemonic()
	if err != nil {
		panic(err)
	}
	return NewSimpleWeb3WalletFromMnemonic(mnemonic, web3Endpoint)
}
