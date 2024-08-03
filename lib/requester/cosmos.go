package requester

import (
	cosmostypes "github.com/hanchon/hanchond/lib/types/cosmos"
	"github.com/hanchon/hanchond/lib/types/evmos"
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

func (c *Client) GetEvmosERC20TokenPairs() (*evmos.TokenPairsResponse, error) {
	var result evmos.TokenPairsResponse
	return &result, c.SendGetRequestEasyJSON(
		c.CosmosRestEndpoint,
		"/evmos/erc20/v1/token_pairs",
		&result,
		c.CosmosRestAuth,
	)
}

func (c *Client) GetIBCRateLimits() (*evmos.RateLimitsResponse, error) {
	var result evmos.RateLimitsResponse
	return &result, c.SendGetRequestEasyJSON(
		c.CosmosRestEndpoint,
		"/Stride-Labs/ibc-rate-limiting/ratelimit/ratelimits",
		&result,
		c.CosmosRestAuth,
	)
}
