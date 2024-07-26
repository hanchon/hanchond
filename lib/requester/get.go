package requester

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

func (c *Client) SendGetRequestEasyJSON(endpoint string, url string, res easyjson.Unmarshaler, auth string) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(endpoint + url)
	req.Header.SetMethod(fasthttp.MethodGet)
	if auth != "" {
		req.Header.Add("Authorization", auth)
	}
	resp := fasthttp.AcquireResponse()
	err := c.Client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New("incorrect status code: " + strconv.Itoa(resp.StatusCode()))
	}

	return easyjson.Unmarshal(resp.Body(), res)
}

func (c *Client) SendGetRequest(endpoint string, url string, auth string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(endpoint + url)
	req.Header.SetMethod(fasthttp.MethodGet)
	if auth != "" {
		req.Header.Add("Authorization", auth)
	}
	resp := fasthttp.AcquireResponse()
	err := c.Client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode() != http.StatusOK {
		return []byte{}, errors.New("incorrect status code: " + strconv.Itoa(resp.StatusCode()))
	}

	ret := make([]byte, len(resp.Body()))
	copy(ret, resp.Body())
	return ret, nil
}
