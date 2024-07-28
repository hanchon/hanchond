package txbuilder

func (t *TxBuilder) GetGasLimit(method string) uint64 {
	v, ok := t.customGasLimit[method]
	if ok {
		return v
	}
	return t.defaultGasLimit
}
