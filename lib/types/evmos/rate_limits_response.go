package evmos

type RateLimitsResponse struct {
	RateLimits []struct {
		Path struct {
			Denom     string `json:"denom"`
			ChannelID string `json:"channel_id"`
		} `json:"path"`
		Quota struct {
			MaxPercentSend string `json:"max_percent_send"`
			MaxPercentRecv string `json:"max_percent_recv"`
			DurationHours  string `json:"duration_hours"`
		} `json:"quota"`
		Flow struct {
			Inflow       string `json:"inflow"`
			Outflow      string `json:"outflow"`
			ChannelValue string `json:"channel_value"`
		} `json:"flow"`
	} `json:"rate_limits"`
}
