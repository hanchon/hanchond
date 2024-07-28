package web3

type SendRawTransactionResponse struct {
	Jsonrpc string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Result  string         `json:"result"`
	Error   BroadcastError `json:"error"`
}

type BroadcastError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
