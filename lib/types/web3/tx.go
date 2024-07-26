package web3

type TxReceipt struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  TxReceiptResult `json:"result"`
}

type TxReceiptResult struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	From              string `json:"from"`
	GasUsed           string `json:"gasUsed"`
	Logs              []struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      string   `json:"blockNumber"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
		BlockHash        string   `json:"blockHash"`
		LogIndex         string   `json:"logIndex"`
		Removed          bool     `json:"removed"`
	} `json:"logs"`
	LogsBloom        string `json:"logsBloom"`
	Status           string `json:"status"`
	To               string `json:"to"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
}
