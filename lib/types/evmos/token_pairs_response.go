package evmos

type TokenPairsResponse struct {
	TokenPairs []struct {
		Erc20Address  string `json:"erc20_address"`
		Denom         string `json:"denom"`
		Enabled       bool   `json:"enabled"`
		ContractOwner string `json:"contract_owner"`
	} `json:"token_pairs"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}
