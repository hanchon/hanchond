package requester

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
	Client       *fasthttp.Client
	Web3Endpoint string
	RestEndpoint string
	Web3Auth     string
	RestAuth     string
}

const (
	defaultWeb3Endpoint = "https://proxy.evmos.org/web3"
	defaultRestEndpoint = "https://proxy.evmos.org/cosmos"
	defaultWeb3Auth     = ""
	defaultRestAuth     = ""

	defaultRequestTimeout = time.Minute
	defaultReadTimeout    = time.Minute
	defaultWriteTimeout   = time.Minute
	defaultMaxIdleConn    = time.Hour
)

func NewClient() Client {
	client := &fasthttp.Client{
		ReadTimeout:                   defaultReadTimeout,
		WriteTimeout:                  defaultWriteTimeout,
		MaxIdleConnDuration:           defaultMaxIdleConn,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}
	return Client{
		Client:       client,
		Web3Endpoint: defaultWeb3Endpoint,
		RestEndpoint: defaultRestEndpoint,
		Web3Auth:     defaultWeb3Auth,
		RestAuth:     defaultRestAuth,
	}
}

func (c *Client) WithUnsecureWeb3Endpoint(endpoint string) *Client {
	c.Web3Endpoint = endpoint
	c.Web3Auth = ""
	return c
}

func (c *Client) WithUnsecureRestEndpoint(endpoint string) *Client {
	c.RestEndpoint = endpoint
	c.RestAuth = ""
	return c
}

func (c *Client) WithSecureWeb3Endpoint(endpoint string, auth string) *Client {
	c.Web3Endpoint = endpoint
	c.Web3Auth = auth
	return c
}

func (c *Client) WithSecureRestEndpoint(endpoint string, auth string) *Client {
	c.RestEndpoint = endpoint
	c.RestAuth = auth
	return c
}
