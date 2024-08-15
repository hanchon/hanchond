package cosmosdaemon

import (
	"fmt"
	"regexp"
	"strings"
)

func (d *Daemon) UpdateConfigFile(withSlowBlocks bool) error {
	// ConfigFile
	configFile, err := d.openConfigFile()
	if err != nil {
		return err
	}

	if withSlowBlocks {
		configFile = d.enableSlowBlocks(configFile)
	}

	configFile = d.allowDuplicateIP(configFile)

	return d.saveConfigFile(configFile)
}

func (d *Daemon) UpdateAppFile() error {
	//  Pruning
	appFile, err := d.OpenAppFile()
	if err != nil {
		return err
	}

	// No pruning to use archive queries
	appFile = d.SetPruningInAppFile(false, appFile)
	appFile = d.SetMinGasPricesInAppFile(appFile)
	return d.SaveAppFile(appFile)
}

func (d *Daemon) CreateGenTx() error {
	validatorAddr, err := d.GetValidatorAddress()
	if err != nil {
		return err
	}
	if err := d.AddGenesisAccount(validatorAddr); err != nil {
		return err
	}
	if err := d.ValidatorGenTx(); err != nil {
		return err
	}
	return nil
}

func (d *Daemon) InitGenesis() error {
	if err := d.setBank(); err != nil {
		return err
	}

	if err := d.CollectGenTxs(); err != nil {
		return err
	}

	if err := d.ValidateGenesis(); err != nil {
		return err
	}

	// Make backups of the AppFile and ConfigFile
	if err := d.backupConfigFiles(); err != nil {
		return err
	}

	return nil
}

func (d *Daemon) allowDuplicateIP(configFile []byte) []byte {
	configValues := string(configFile)
	// Tendermint Values
	configValues = strings.Replace(
		configValues,
		"allow_duplicate_ip = false",
		"allow_duplicate_ip = true",
		1,
	)
	return []byte(configValues)
}

func (d *Daemon) enableSlowBlocks(configFile []byte) []byte {
	configValues := string(configFile)
	// Tendermint Values
	configValues = strings.Replace(
		configValues,
		"timeout_propose = \"3s\"",
		"timeout_propose = \"10s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_propose_delta = \"500ms\"",
		"timeout_propose_delta = \"1s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_prevote = \"1s\"",
		"timeout_prevote = \"5s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_prevote_delta = \"500ms\"",
		"timeout_prevote_delta = \"1s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_precommit = \"1s\"",
		"timeout_precommit = \"5s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_precommit_delta = \"500ms\"",
		"timeout_precommit_delta = \"1s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_commit = \"3s\"",
		"timeout_commit = \"5s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_broadcast_tx_commit = \"10s\"",
		"timeout_broadcast_tx_commit = \"15s\"",
		1,
	)
	return []byte(configValues)
}

func (d *Daemon) EnableWeb3API(config []byte) []byte {
	configValues := string(config)
	configValues = strings.Replace(
		configValues,
		`# Enable defines if the JSONRPC server should be enabled.
enable = false`,
		"enable = true",
		1,
	)
	return []byte(configValues)
}

func (d *Daemon) SetMinGasPricesInAppFile(config []byte) []byte {
	configValues := string(config)
	configValues = strings.Replace(
		configValues,
		"minimum-gas-prices = \"\"",
		"minimum-gas-prices = \"0.00001"+d.BaseDenom+"\"",
		1,
	)
	return []byte(configValues)
}

func (d *Daemon) SetPruningInAppFile(pruningEnabled bool, config []byte) []byte {
	configValues := string(config)
	if pruningEnabled {
		configValues = strings.Replace(
			configValues,
			"pruning = \"default\"",
			"pruning = \"custom\"",
			1,
		)
		configValues = strings.Replace(
			configValues,
			"pruning-keep-recent = \"0\"",
			"pruning-keep-recent = \"2\"",
			1,
		)
		configValues = strings.Replace(
			configValues,
			"pruning-interval = \"0\"",
			"pruning-interval = \"10\"",
			1,
		)
		return []byte(configValues)
	}

	// NoPrunning
	configValues = strings.Replace(
		configValues,
		"pruning = \"default\"",
		"pruning = \"nothing\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"pruning = \"custom\"",
		"pruning = \"nothing\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"pruning = \"everything\"",
		"pruning = \"nothing\"",
		1,
	)
	return []byte(configValues)
}

func (d *Daemon) GetPeerInfo() (string, error) {
	nodeID, err := d.GetNodeID()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@127.0.0.1:%d", nodeID, d.Ports.P26656), nil
}

func (d *Daemon) AddPersistenPeers(peers []string) error {
	filtered := []string{}
	for k := range peers {
		// Exclude ourself from the list
		if !strings.Contains(peers[k], fmt.Sprintf("%d", d.Ports.P26656)) {
			filtered = append(filtered, peers[k])
		}
	}

	configFile, err := d.openConfigFile()
	if err != nil {
		return err
	}
	regex := regexp.MustCompile(`persistent_peers\s*=\s*".*"`)
	configFile = regex.ReplaceAll(
		configFile,
		[]byte(
			fmt.Sprintf(
				"persistent_peers = \"%s\"",
				strings.Join(filtered, ","),
			),
		),
	)

	if err := d.saveConfigFile(configFile); err != nil {
		return err
	}

	return nil
}

func (d *Daemon) UpdateConfigPorts() error {
	appFile, err := d.OpenAppFile()
	if err != nil {
		return err
	}
	app := string(appFile)
	app = strings.Replace(app, "1317", fmt.Sprint(d.Ports.P1317), 1)
	app = strings.Replace(app, "8080", fmt.Sprint(d.Ports.P8080), 1)
	app = strings.Replace(app, "9090", fmt.Sprint(d.Ports.P9090), 1)
	app = strings.Replace(app, "9091", fmt.Sprint(d.Ports.P9091), 1)
	app = strings.Replace(app, "8545", fmt.Sprint(d.Ports.P8545), 1)
	app = strings.Replace(app, "8546", fmt.Sprint(d.Ports.P8546), 1)
	app = strings.Replace(app, "6065", fmt.Sprint(d.Ports.P6065), 1)
	if err := d.SaveAppFile([]byte(app)); err != nil {
		return err
	}

	configFile, err := d.openConfigFile()
	if err != nil {
		return err
	}

	config := string(configFile)
	config = strings.Replace(config, "26656", fmt.Sprint(d.Ports.P26656), 1)
	config = strings.Replace(config, "26657", fmt.Sprint(d.Ports.P26657), 1)
	config = strings.Replace(config, "26658", fmt.Sprint(d.Ports.P26658), 1)
	config = strings.Replace(config, "26660", fmt.Sprint(d.Ports.P26660), 1)
	config = strings.Replace(config, "6060", fmt.Sprint(d.Ports.P6060), 1)
	if err := d.saveConfigFile([]byte(config)); err != nil {
		return err
	}

	return nil
}
