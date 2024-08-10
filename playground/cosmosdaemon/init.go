package cosmosdaemon

func (d *Daemon) InitNode() error {
	if err := d.ConfigKeyring(); err != nil {
		return err
	}

	if err := d.ConfigChainID(); err != nil {
		return err
	}

	if err := d.NodeInit(); err != nil {
		return err
	}

	if err := d.AddValidatorKey(); err != nil {
		return err
	}
	return nil
}
