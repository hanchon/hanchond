package requester

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	web3types "github.com/hanchon/hanchond/lib/types/web3"
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

func (c *Client) GetTransactionReceiptWithRetry(hash string, retryInSeconds int) (receipt *web3types.TxReceipt, err error) {
	retry := 0
	for retry < retryInSeconds {
		receipt, err = c.GetTransactionReceipt(hash)
		if err != nil || receipt.Result.BlockHash == "" {
			time.Sleep(time.Second)
			continue
		}
		return
	}

	return nil, err
}

func (c *Client) GetNonce(address string) (uint64, error) {
	var resp web3types.NonceResponse
	if err := c.SendPostRequestEasyJSON(
		c.Web3Endpoint,
		[]byte(`{"method":"eth_getTransactionCount","params":["`+address+`", "latest"],"id":1,"jsonrpc":"2.0"}`),
		&resp,
		c.Web3Auth,
	); err != nil {
		return 0, err
	}
	return strconv.ParseUint(resp.Result, 0, 64)
}

func (c *Client) GasPrice() (*big.Int, error) {
	var resp web3types.GasPriceResponse
	if err := c.SendPostRequestEasyJSON(
		c.Web3Endpoint,
		[]byte(`{"method":"eth_gasPrice","params":[],"id":1,"jsonrpc":"2.0"}`),
		&resp,
		c.Web3Auth,
	); err != nil {
		return nil, err
	}

	supply := new(big.Int)
	supply.SetString(resp.Result[2:], 16)
	return supply, nil
}

func (c *Client) ChanID() (*big.Int, error) {
	var resp web3types.NetVersionResponse
	if err := c.SendPostRequestEasyJSON(
		c.Web3Endpoint,
		[]byte(`{"method":"net_version","params":[],"id":1,"jsonrpc":"2.0"}`),
		&resp,
		c.Web3Auth,
	); err != nil {
		return nil, err
	}

	version := new(big.Int)
	if _, ok := version.SetString(resp.Result, 10); !ok {
		return nil, fmt.Errorf("invalid chainID %q", resp.Result)
	}

	return version, nil
}

// BroadcastTx returns txhash and error
func (c *Client) BroadcastTx(tx *coretypes.Transaction) (string, error) {
	var resp web3types.SendRawTransactionResponse
	data, err := tx.MarshalBinary()
	if err != nil {
		return "", err
	}

	if err := c.SendPostRequestEasyJSON(
		c.Web3Endpoint,
		[]byte(`{"method":"eth_sendRawTransaction","params":["`+hexutil.Encode(data)+`"],"id":1,"jsonrpc":"2.0"}`),
		&resp,
		c.Web3Auth,
	); err != nil {
		return "", err
	}

	// Success
	if resp.Result != "" {
		return resp.Result, nil
	}

	return "", fmt.Errorf("%s", resp.Error.Message)
}

// Eth_call
func (c *Client) EthCall(address string, data string, height string) ([]byte, error) {
	heightString, err := heigthToQueryParam(height)
	if err != nil {
		return nil, err
	}
	return c.SendPostRequest(
		c.Web3Endpoint,
		[]byte(`{"method":"eth_call","params":[{"to":"`+address+`","data":"`+data+`"},"`+heightString+`"],"id":1,"jsonrpc":"2.0"}`),
		c.Web3Auth,
	)
}

// Eth_code
func (c *Client) EthCode(address string, height string) ([]byte, error) {
	heightString, err := heigthToQueryParam(height)
	if err != nil {
		return nil, err
	}
	return c.SendPostRequest(
		c.Web3Endpoint,
		[]byte(`{"method":"eth_getCode","params":["`+address+`","`+heightString+`"],"id":1,"jsonrpc":"2.0"}`),
		c.Web3Auth,
	)
}

func heigthToQueryParam(height string) (string, error) {
	heightString := "latest"
	if height != "latest" {
		temp, err := strconv.ParseInt(height, 10, 64)
		if err != nil {
			return "", fmt.Errorf("invalid height: %s", err.Error())
		}

		heightString = fmt.Sprintf("0x%x", temp)
	}
	return heightString, nil
}
