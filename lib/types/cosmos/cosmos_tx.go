package cosmos

type TxRestResponseForEvents struct {
	TxResponse struct {
		Height string `json:"height"`
		Txhash string `json:"txhash"`
		Code   int    `json:"code"`
		Logs   []struct {
			MsgIndex int    `json:"msg_index"`
			Log      string `json:"log"`
			Events   []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"logs"`
		GasWanted string `json:"gas_wanted"`
		GasUsed   string `json:"gas_used"`
		Events    []struct {
			Type       string `json:"type"`
			Attributes []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
				Index bool   `json:"index"`
			} `json:"attributes"`
		} `json:"events"`
	} `json:"tx_response"`
}
