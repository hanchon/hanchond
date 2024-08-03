package requester

import (
	tenderminttypes "github.com/hanchon/hanchond/lib/types/tendermint"
)

func (c *Client) GetChainStatus() (*tenderminttypes.StatusResponse, error) {
	var result tenderminttypes.StatusResponse
	return &result, c.SendGetRequestEasyJSON(
		c.TendermintRestEndpoint,
		"/status",
		&result,
		c.TendermintRestAuth,
	)
}

func (c *Client) GetCurrentHeight() (string, error) {
	status, err := c.GetChainStatus()
	if err != nil {
		return "", err
	}
	return status.Result.SyncInfo.LatestBlockHeight, nil
}
