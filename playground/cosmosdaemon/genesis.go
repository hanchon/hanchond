package cosmosdaemon

func (e *Daemon) initNode() error {
	if err := e.ConfigKeyring(); err != nil {
		return err
	}

	if err := e.ConfigChainID(); err != nil {
		return err
	}

	if err := e.NodeInit(); err != nil {
		return err
	}
	return nil
}

func (e *Daemon) InitGenesis() error {
	if err := e.initNode(); err != nil {
		return err
	}

	if err := e.AddValidatorKey(); err != nil {
		return err
	}

	genesis, err := e.openGenesisFile()
	if err != nil {
		return err
	}
	// Update the genesis
	e.UpdateGenesisDenom(genesis)
	e.SetGenesisGasLimit(genesis)
	e.SetGenesisBaseFee(genesis)
	e.SetGenesisFastProposals(genesis)

	if err := e.saveGenesisFile(genesis); err != nil {
		return err
	}

	// ConfigFile
	configFile, err := e.openConfigFile()
	if err != nil {
		return err
	}
	// TODO: renable this if we want slow blocks
	// configFile = e.UpdateConfigFile(configFile)
	if err := e.saveConfigFile(configFile); err != nil {
		return err
	}

	//  Pruning
	appFile, err := e.openAppFile()
	if err != nil {
		return err
	}
	appFile = e.SetPruningInAppFile(true, appFile)
	if err := e.saveAppFile(appFile); err != nil {
		return err
	}

	if err := e.AddGenesisAccount(); err != nil {
		return err
	}

	if err := e.UpdateTotalSupply(); err != nil {
		return err
	}

	if err := e.ValidatorGenTx(); err != nil {
		return err
	}

	if err := e.CollectGenTxs(); err != nil {
		return err
	}

	if err := e.ValidateGenesis(); err != nil {
		return err
	}

	// Make backups of the AppFile and ConfigFile
	if err := e.backupConfigFiles(); err != nil {
		return err
	}

	return nil
}
