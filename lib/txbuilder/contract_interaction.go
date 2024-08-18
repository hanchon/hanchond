package txbuilder

import "math/big"

func (t *TxBuilder) InteractWithContract(
	contractName string,
	accountID int,
	value *big.Int,
	message string,
	args ...interface{},
) (string, error) {
	wallet, account, err := WalletFromMnemonicWithAccountID(t.mnemonic, accountID)
	if err != nil {
		return "", err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return "", err
	}

	return t.SendTxToContract(contractName, account.Address, privateKey, value, message, args...)
}
