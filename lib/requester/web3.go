package requester

import (
	"fmt"

	web3types "github.com/hanchon/vivi/lib/types/web3"
)

func (c *Client) GetBlockByNumber(height string, withTransactions bool) (*web3types.BlockByNumberWithTransactions, error) {
	var receipt web3types.BlockByNumberWithTransactions
	return &receipt, c.SendPostRequestEasyJSON(
		c.Web3Endpoint,
		[]byte(fmt.Sprintf(`{"method":"eth_getBlockByNumber","params":["%s",%t],"id":1,"jsonrpc":"2.0"}`, height, withTransactions)),
		&receipt,
		c.Web3Auth,
	)
}

func (c *Client) GetTransactionTrace(hash string) (*web3types.TraceTransactionResult, error) {
	var trace web3types.TraceTransactionResult
	return &trace, c.SendPostRequestEasyJSON(
		c.Web3Endpoint,
		[]byte(`{"method":"debug_traceTransaction","params":["`+hash+`", {"tracer": "callTracer"}],"id":1,"jsonrpc":"2.0"}`),
		&trace,
		c.Web3Auth,
	)
}

func (c *Client) GetTransactionReceipt(hash string) (*web3types.TxReceipt, error) {
	var receipt web3types.TxReceipt
	return &receipt, c.SendPostRequestEasyJSON(
		c.Web3Endpoint,
		[]byte(`{"method":"eth_getTransactionReceipt","params":["`+hash+`"],"id":1,"jsonrpc":"2.0"}`),
		&receipt,
		c.Web3Auth,
	)
}
