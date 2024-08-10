package cosmosdaemon

import (
	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/txbuilder"
)

type SignatureAlgo string

const (
	EthAlgo    SignatureAlgo = "eth_secp256k1"
	CosmosAlgo SignatureAlgo = "secp256k1"
)

type SDKVersion string

const (
	// NOTE: there are some differences in the namespace while interacting with the CLI, like the genesis namespace
	GaiaSDK  SDKVersion = "gaiaSDK"
	EvmosSDK SDKVersion = "evmosSDK"
)

type Daemon struct {
	ValKeyName  string
	ValMnemonic string
	ValWallet   string
	KeyType     SignatureAlgo

	KeyringBackend string
	HomeDir        string
	BinaryName     string
	SDKVersion     SDKVersion

	ChainID string
	Moniker string

	BaseDenom string
	GasLimit  string
	BaseFee   string

	ValidatorInitialSupply string

	Ports *Ports

	BinaryPath string
}

func NewDameon(
	moniker string,
	binaryName string,
	homeDir string,
	chainID string,
	keyName string,
	algo SignatureAlgo,
	denom string,
	prefix string,
	sdkVersion SDKVersion,
) *Daemon {
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
		SDKVersion:     sdkVersion,

		ChainID: chainID,
		Moniker: moniker,

		BaseDenom: denom,

		// Maybe move this to just evmos
		GasLimit:               "10000000",
		BaseFee:                "1000000000",
		ValidatorInitialSupply: "100000000000000000000000000",

		Ports: nil,
	}
}

func (d *Daemon) SetBinaryPath(path string) {
	d.BinaryPath = path
}
