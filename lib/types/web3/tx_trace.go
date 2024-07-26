package web3

type Calls struct {
	From    string  `json:"from"`
	Gas     string  `json:"gas"`
	GasUsed string  `json:"gasUsed"`
	Input   string  `json:"input"`
	Output  string  `json:"output"`
	To      string  `json:"to"`
	Type    string  `json:"type"`
	Value   string  `json:"value"`
	Error   string  `json:"error"`
	Calls   []Calls `json:"calls"`
}

type TraceTransactionResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Calls  `json:"result"`
}
