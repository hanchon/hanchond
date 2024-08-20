package txbuilder

import (
	"math/big"
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

	return t.SendTx(account.Address, nil, value, gasLimit, bytecode, privateKey)
}
