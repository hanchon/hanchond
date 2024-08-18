package cosmosdaemon

import (
	"fmt"

	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/requester"
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
	Prefix      string

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

	CustomConfig func() error
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
		Prefix:      prefix,

		KeyType: algo,

		KeyringBackend: "test",
		HomeDir:        homeDir,
		BinaryName:     binaryName,
		SDKVersion:     sdkVersion,

		ChainID: chainID,
		Moniker: moniker,

		BaseDenom: denom,

		ValidatorInitialSupply: "100000000000000000000000000",

		// Maybe move this to just evmos
		GasLimit: "10000000",
		BaseFee:  "1000000000",

		Ports: nil,
	}
}

func (d *Daemon) SetBinaryPath(path string) {
	d.BinaryPath = path
}

// This is used to change the config files that are specific to a client
func (d *Daemon) SetCustomConfig(configurator func() error) {
	d.CustomConfig = configurator
}

func (d *Daemon) ExecuteCustomConfig() error {
	if d.CustomConfig == nil {
		return nil
	}
	return d.CustomConfig()
}

func (d *Daemon) SetValidatorWallet(mnemonic, wallet string) {
	d.ValMnemonic = mnemonic
	d.ValWallet = wallet
}

func (d *Daemon) NewRequester() *requester.Client {
	return requester.NewClient().
		WithUnsecureWeb3Endpoint(fmt.Sprintf("http://localhost:%d", d.Ports.P8545)).
		WithUnsecureRestEndpoint(fmt.Sprintf("http://localhost:%d", d.Ports.P1317)).
		WithUnsecureTendermintEndpoint(fmt.Sprintf("http://localhost:%d", d.Ports.P26657))
}

func (d *Daemon) NewTxBuilder(gasLimit uint64) *txbuilder.TxBuilder {
	return txbuilder.NexTxBuilder(
		map[string]txbuilder.Contract{},
		d.ValMnemonic,
		map[string]uint64{},
		gasLimit,
		d.NewRequester(),
	)
}
