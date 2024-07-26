package erc20

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type Request struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

type ERC20 struct {
	Client *rpc.Client
}

const ethCall = "eth_call"

func NewERC20(web3Endpoint string) (*ERC20, error) {
	client, err := rpc.DialHTTP(web3Endpoint)
	if err != nil {
		return nil, err
	}

	return &ERC20{
		Client: client,
	}, err
}
