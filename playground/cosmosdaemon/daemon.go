package cosmosdaemon

import "github.com/hanchon/hanchond/lib/txbuilder"

type Dameon struct {
	ValKeyName  string
	ValMnemonic string
	KeyType     string

	KeyringBackend string
	HomeDir        string
	Version        string

	ChainID string
	Moniker string

	BaseDenom string
	GasLimit  string
	BaseFee   string

	ValidatorInitialSupply string

	Ports Ports
}

const (
	// --key-type string          Key signing algorithm to generate keys for (default "")
	EthAlgo    = "eth_secp256k1"
	CosmosAlgo = "secp256k1"
)

func NewDameon(version string, homeDir string, chainID string, keyName string, algo string, denom string) *Dameon {
	mnemonic, _ := txbuilder.NewMnemonic()
	return &Dameon{
		ValKeyName:  keyName,
		ValMnemonic: mnemonic,
		KeyType:     algo,

		KeyringBackend: "test",
		HomeDir:        homeDir,
		Version:        version,

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
