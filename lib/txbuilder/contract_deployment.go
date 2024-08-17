package txbuilder

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (t *TxBuilder) DeployContract(
	accountID int,
	bytecode []byte,
	gasLimit uint64,
) (string, error) {
	value := big.NewInt(0)
	wallet, account, err := WalletFromMnemonicWithAccountID(t.mnemonic, accountID)
	if err != nil {
		return "", err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return "", err
	}

	return t.SendTx(account.Address, common.Address{}, value, gasLimit, bytecode, privateKey)
}
