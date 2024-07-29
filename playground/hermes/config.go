package hermes

import (
	"fmt"
	"strings"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (h *Hermes) GetConfigFile() string {
	// If the dir already existed it will return error, but that is fine
	_ = filesmanager.CreateHermesFolder()
	return filesmanager.GetHermesPath() + "/config.toml"
}

func (h *Hermes) AddEvmosChain(chainID string, p26657 int64, p9090 int64, keyname string, mnemonic string) error {
	values := fmt.Sprintf(`
[[chains]]
id = '%s'
rpc_addr = 'http://127.0.0.1:%d'
grpc_addr = 'http://127.0.0.1:%d'
event_source = { mode = 'pull', interval = '1s' }
rpc_timeout = '10s'
account_prefix = 'evmos'
key_name = '%s'
key_store_folder = '%s'
store_prefix = 'ibc'
default_gas = 100000
max_gas = 3000000
clock_drift = '15s'
max_block_time = '10s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
gas_price = { price = 800000000, denom = 'aevmos' }
address_type = { derivation = 'ethermint', proto_type = { pk_type = '/ethermint.crypto.v1.ethsecp256k1.PubKey' } }
`, chainID, p26657, p9090, keyname, filesmanager.GetHermesPath()+"/keyring"+chainID)

	configFile, err := filesmanager.ReadFile(h.GetConfigFile())
	if err != nil {
		return err
	}

	configFileString := string(configFile)
	// If the chain was already included in the config file, do nothing
	if strings.Contains(configFileString, chainID) {
		// Maybe update ports if needed
		return nil
	}
	configFileString += values
	err = filesmanager.SaveFile([]byte(configFileString), h.GetConfigFile())
	if err != nil {
		panic(err)
	}

	err = h.AddRelayerKey(chainID, mnemonic)
	if err != nil {
		panic(err)
	}
	return nil
}

func (h *Hermes) initHermesConfig() {
	// Init the file only if it does not exist
	if filesmanager.DoesFileExist(h.GetConfigFile()) {
		return
	}

	basicConfig := `
[global]
log_level = 'trace'

[mode]

[mode.clients]
enabled = true
refresh = true

[mode.connections]
enabled = false

[mode.channels]
enabled = false

[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = true

[rest]
enabled = false
host = '127.0.0.1'
port = 3000

[telemetry]
enabled = false
host = '127.0.0.1'
port = 3001
`
	err := filesmanager.SaveFile([]byte(basicConfig), h.GetConfigFile())
	if err != nil {
		panic(err)
	}
}
