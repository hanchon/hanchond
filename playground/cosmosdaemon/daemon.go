package cosmosdaemon

import (
	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/txbuilder"
)

type Daemon struct {
	ValKeyName  string
	ValMnemonic string
	ValWallet   string
	KeyType     string

	KeyringBackend string
	HomeDir        string
	BinaryName     string

	ChainID string
	Moniker string

	BaseDenom string
	GasLimit  string
	BaseFee   string

	ValidatorInitialSupply string

	Ports Ports

	BinaryPath string
}

const (
	EthAlgo    = "eth_secp256k1"
	CosmosAlgo = "secp256k1"
)

func NewDameon(binaryName string, homeDir string, chainID string, keyName string, algo string, denom string, prefix string) *Daemon {
	mnemonic, _ := txbuilder.NewMnemonic()
	wallet := ""
	switch algo {
	case EthAlgo:
		_, temp, _ := txbuilder.WalletFromMnemonic(mnemonic)
		wallet, _ = converter.HexToBech32(temp.Address.Hex(), prefix)
	case CosmosAlgo:
		wallet, _ = txbuilder.MnemonicToCosmosAddress(mnemonic, prefix)
	}

	return &Daemon{
		ValKeyName:  keyName,
		ValMnemonic: mnemonic,
		ValWallet:   wallet,

		KeyType: algo,

		KeyringBackend: "test",
		HomeDir:        homeDir,
		BinaryName:     binaryName,

		ChainID: chainID,
		Moniker: "moniker",

		BaseDenom: denom,

		// Maybe move this to just evmos
		GasLimit:               "10000000",
		BaseFee:                "1000000000",
		ValidatorInitialSupply: "100000000000000000000000000",

		Ports: *NewPorts(),
	}
}

func (d *Daemon) SetBinaryPath(path string) {
	d.BinaryPath = path
}
