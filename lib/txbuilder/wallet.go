package txbuilder

import (
	"crypto/sha256"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/ethereum/go-ethereum/accounts"
	evmoshd "github.com/evmos/evmos/v18/crypto/hd"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// WalletFromMnemonicWithPath requires the complete mnemonic path to generate the wallet
func WalletFromMnemonicWithPath(mnemonic string, mnemonicPath string) (*hdwallet.Wallet, accounts.Account, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, accounts.Account{}, err
	}

	path := hdwallet.MustParseDerivationPath(mnemonicPath)
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, accounts.Account{}, err
	}
	return wallet, account, nil
}

// WalletFromMnemonicWithAccountID returns the wallet with id=accountID, always using the main path of derivation
func WalletFromMnemonicWithAccountID(mnemonic string, accountID int) (*hdwallet.Wallet, accounts.Account, error) {
	return WalletFromMnemonicWithPath(mnemonic, fmt.Sprintf("m/44'/60'/0'/0/%d", accountID))
}

// WalletFromMnemonic returns the first wallet generated by the mnemonic
func WalletFromMnemonic(mnemonic string) (*hdwallet.Wallet, accounts.Account, error) {
	return WalletFromMnemonicWithPath(mnemonic, "m/44'/60'/0'/0/0")
}

func NewMnemonicFromEntropy(entropy string) (string, error) {
	if len(entropy) < 43 {
		return "", fmt.Errorf("256-bits is 43 characters in Base-64, and 100 in Base-6. You entered %v, and probably want more", len(entropy))
	}
	hashedEntropy := sha256.Sum256([]byte(entropy))
	return bip39.NewMnemonic(hashedEntropy[:])
}

func NewMnemonic() (string, error) {
	const mnemonicEntropySize = 256
	// read entropy seed straight from crypto.Rand
	var err error
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropySeed)
}

func MnemonicToAddressWithPath(mnemonic string, hdPath string, prefix string, algo keyring.SignatureAlgo) (string, error) {
	derivedPriv, err := algo.Derive()(mnemonic, "", hdPath)
	if err != nil {
		return "", err
	}
	privKey := algo.Generate()(derivedPriv)
	return sdk.Bech32ifyAddressBytes(prefix, privKey.PubKey().Address().Bytes())
}

func MnemonicToCosmosAddressWithPath(mnemonic string, hdPath string, prefix string) (string, error) {
	return MnemonicToAddressWithPath(mnemonic, hdPath, prefix, hd.Secp256k1)
}

func MnemonicToCosmosAddress(mnemonic string, prefix string) (string, error) {
	hdPath := hd.CreateHDPath(118, 0, 0).String()
	return MnemonicToCosmosAddressWithPath(mnemonic, hdPath, prefix)
}

func MnemonicToEthereumAddressWithPath(mnemonic string, hdPath string, prefix string) (string, error) {
	return MnemonicToAddressWithPath(mnemonic, hdPath, prefix, evmoshd.EthSecp256k1)
}

func MnemonicToEthereumAddress(mnemonic string, prefix string) (string, error) {
	hdPath := hd.CreateHDPath(60, 0, 0).String()
	return MnemonicToEthereumAddressWithPath(mnemonic, hdPath, prefix)
}
