package erc20

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"

	"github.com/hanchon/hanchond/lib/converter"
)

const Latest int = math.MaxInt

func (e *ERC20) GetTotalSupply(contractAddress string, height int) (*big.Int, error) {
	contractAddress = converter.RemoveHexPrefixFromAddress(contractAddress)

	method := "0x" + hex.EncodeToString(contract.Methods["totalSupply"].ID)
	req := Request{"0x" + contractAddress, method}
	var result string

	var args []interface{}
	args = append(args, req)

	heightInHex := "latest"
	if height != Latest {
		heightInHex = fmt.Sprintf("0x%x", height)
	}
	args = append(args, heightInHex)

	if err := e.Client.Call(&result, ethCall, args...); err != nil {
		return nil, err
	}

	supply := new(big.Int)
	supply.SetString(result[2:], 16)
	return supply, nil
}

func (e *ERC20) GetBalance(contractAddress string, wallet string, height int) (*big.Int, error) {
	contractAddress = converter.RemoveHexPrefixFromAddress(contractAddress)
	wallet = converter.RemoveHexPrefixFromAddress(wallet)

	method := hex.EncodeToString(contract.Methods["balanceOf"].ID)
	params := method + "000000000000000000000000" + wallet

	type Request struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}
	req := Request{"0x" + contractAddress, "0x" + params}
	var result string

	var args []interface{}
	args = append(args, req)

	heightInHex := "latest"
	if height != Latest {
		heightInHex = fmt.Sprintf("0x%x", height)
	}
	args = append(args, heightInHex)

	if err := e.Client.Call(&result, ethCall, args...); err != nil {
		return nil, err
	}

	balance := new(big.Int)
	balance.SetString(result[2:], 16)
	return balance, nil
}
