package txbuilder

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (t *TxBuilder) SendCoins(to string, amount *big.Int) (string, error) {
	wallet, account, err := WalletFromMnemonicWithAccountID(t.mnemonic, 0)
	if err != nil {
		return "", err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(to)
	return t.SendTx(
		account.Address,
		&toAddress,
		amount,
		25_000,
		[]byte{},
		privateKey,
	)
}
