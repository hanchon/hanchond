package requester

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

func (c *Client) SendPostRequestEasyJSON(endpoint string, body []byte, res easyjson.Unmarshaler, auth string) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(endpoint)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	if auth != "" {
		req.Header.Add("Authorization", auth)
	}
	req.SetBodyRaw(body)
	resp := fasthttp.AcquireResponse()
	err := c.Client.DoTimeout(req, resp, c.Client.ReadTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		return err
	}

	statusCode := resp.StatusCode()
	if statusCode != http.StatusOK {
		return fmt.Errorf("status code is not ok: " + strconv.Itoa(statusCode))
	}

	respBody := resp.Body()
	return easyjson.Unmarshal(respBody, res)
}
