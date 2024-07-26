package erc20

import (
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const erc20MinimalAbiJSON = `
[{
  "inputs": [],
  "name": "totalSupply",
  "outputs": [
    {
      "internalType": "uint256",
      "name": "",
      "type": "uint256"
    }
  ],
  "stateMutability": "view",
  "type": "function"
},{
    "inputs": [
    {
        "internalType": "address",
        "name": "account",
        "type": "address"
    }
    ],
    "name": "balanceOf",
    "outputs": [
    {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
    }
    ],
    "stateMutability": "view",
    "type": "function"
}]
`

var contract abi.ABI

func init() {
	var err error
	contract, err = abi.JSON(bytes.NewReader([]byte(erc20MinimalAbiJSON)))
	if err != nil {
		panic(err)
	}
}
