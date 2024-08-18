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

	// TODO: add support for constructors args
	// parsedABI, err := abi.JSON(bytes.NewReader(abiFile))
	// if err != nil {
	//     log.Fatalf("Failed to parse ABI: %v", err)
	// }
	// input, err := parsedABI.Pack("", "Hello, Ethereum!")
	// if err != nil {
	//     log.Fatalf("Failed to pack input parameters: %v", err)
	// }
	//
	// // Combine bytecode + constructor arguments
	// fullBytecode := append(bytecode, input...)

	return t.SendTx(account.Address, nil, value, gasLimit, bytecode, privateKey)
}
