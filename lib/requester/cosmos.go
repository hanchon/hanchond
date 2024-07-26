package requester

import (
	cosmostypes "github.com/hanchon/vivi/lib/types/cosmos"
)

func (c *Client) GetBlockCosmos(height string) (*cosmostypes.CosmosBlockResult, error) {
	// TODO: add pagination support
	var result cosmostypes.CosmosBlockResult
	return &result, c.SendGetRequestEasyJSON(
		c.CosmosRestEndpoint,
		"/cosmos/tx/v1beta1/txs/block/"+height+"?pagination.count_total=true",
		&result,
		c.CosmosRestAuth,
	)
}

func (c *Client) GetCosmosTx(hash string) (*cosmostypes.TxRestResponseForEvents, error) {
	// TODO: add the 0x prefix if not included in the hash string
	var result cosmostypes.TxRestResponseForEvents
	return &result, c.SendGetRequestEasyJSON(
		c.CosmosRestEndpoint,
		"/cosmos/tx/v1beta1/txs/"+hash,
		&result,
		c.CosmosRestAuth,
	)
}
