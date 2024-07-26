package web3

type Tx struct {
	From        string `json:"from"`
	To          string `json:"to"`
	GasUsed     string `json:"gas"`
	ReturnValue string `json:"returnValue"`
	Error       string `json:"error"`
}

type ResultTraceBlock struct {
	ID      int64              `json:"id"`
	Jsonrpc string             `json:"jsonrpc"`
	Result  []TraceBlockResult `json:"result"`
}
type TraceBlockResult struct {
	Result TraceBlockValue `json:"result"`
}
type TraceBlockValue struct {
	Failed      bool   `json:"failed"`
	Gas         int64  `json:"gas"`
	ReturnValue string `json:"returnValue"`
	StructLogs  []Logs `json:"structLogs"`
}

type Logs struct {
	Depth   int64    `json:"depth"`
	Gas     int64    `json:"gas"`
	GasCost int64    `json:"gasCost"`
	Op      string   `json:"op"`
	Pc      int64    `json:"pc"`
	Stack   []string `json:"stack"`
}
