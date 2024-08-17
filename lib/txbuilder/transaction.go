package txbuilder

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *TxBuilder) SendTxToContract(contractName string, address common.Address, privateKey *ecdsa.PrivateKey, value *big.Int, message string, args ...interface{}) (string, error) {
	var contractABI abi.ABI
	var contractAddress common.Address
	var err error

	if v, ok := t.contracts[contractName]; ok {
		contractABI = v.ABI
		contractAddress = v.address
	} else {
		return "", fmt.Errorf("invalid contract name")
	}

	var data []byte
	data, err = contractABI.Pack(message, args...)
	if err != nil {
		return "", err
	}

	gasLimit := t.GetGasLimit(message)

	return t.SendTx(address, contractAddress, value, gasLimit, data, privateKey)
}

func (t *TxBuilder) SendTx(from, to common.Address, value *big.Int, gasLimit uint64, data []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	var err error

	v, ok := t.currentNonce[from.Hex()]
	nonce := uint64(0)
	if ok {
		nonce = v
	} else if nonce, err = t.requester.GetNonce(from.Hex()); err != nil {
		return "", err
	}

	gasPrice, err := t.requester.GasPrice()
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	chainID, err := t.requester.ChanID()
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	txhash, err := t.requester.BroadcastTx(signedTx)
	if err != nil {
		return "", err
	}

	t.currentNonce[from.Hex()] = nonce + 1

	return txhash, nil
}
