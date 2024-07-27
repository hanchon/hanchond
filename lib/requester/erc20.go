package requester

import (
	"fmt"
	"math/big"

	"github.com/hanchon/hanchond/lib/smartcontract/erc20"
)

func (c *Client) GetERC20Client() (*erc20.ERC20, error) {
	var err error
	if c.ERC20Client != nil {
		return c.ERC20Client, nil
	}

	if c.Web3Auth != "" {
		// TODO: update the ERC20 module to support basic auth
		return nil, fmt.Errorf("erc20 pkg only supports unsecured endpoints")
	}

	c.ERC20Client, err = erc20.NewERC20(c.Web3Endpoint)
	if err != nil {
		return nil, err
	}
	return c.ERC20Client, nil
}

func (c *Client) GetTotalSupply(contractAddress string, height int) (*big.Int, error) {
	client, err := c.GetERC20Client()
	if err != nil {
		return nil, err
	}
	return client.GetTotalSupply(contractAddress, height)
}

func (c *Client) GetContractBalance(contractAddress string, wallet string, height int) (*big.Int, error) {
	client, err := c.GetERC20Client()
	if err != nil {
		return nil, err
	}
	return client.GetContractBalance(contractAddress, wallet, height)
}
