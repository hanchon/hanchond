package requester

import (
	"time"

	"github.com/hanchon/hanchond/lib/smartcontract/erc20"
	"github.com/valyala/fasthttp"
)

type Client struct {
	Client                 *fasthttp.Client
	Web3Endpoint           string
	CosmosRestEndpoint     string
	TendermintRestEndpoint string

	Web3Auth           string
	CosmosRestAuth     string
	TendermintRestAuth string

	ERC20Client *erc20.ERC20
}

const (
	defaultWeb3Endpoint           = "https://proxy.evmos.org/web3"
	defaultCosmosRestEndpoint     = "https://proxy.evmos.org/cosmos"
	defaultTendermintRestEndpoint = "https://proxy.evmos.org/tendermint"

	defaultWeb3Auth           = ""
	defaultCosmosRestAuth     = ""
	defaultTendermintRestAuth = ""

	defaultRequestTimeout = time.Minute
	defaultReadTimeout    = time.Minute
	defaultWriteTimeout   = time.Minute
	defaultMaxIdleConn    = time.Hour
)

func NewClient() *Client {
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
	return &Client{
		Client: client,

		Web3Endpoint:           defaultWeb3Endpoint,
		CosmosRestEndpoint:     defaultCosmosRestEndpoint,
		TendermintRestEndpoint: defaultTendermintRestEndpoint,

		Web3Auth:           defaultWeb3Auth,
		CosmosRestAuth:     defaultCosmosRestAuth,
		TendermintRestAuth: defaultTendermintRestAuth,

		ERC20Client: nil,
	}
}

func (c *Client) WithUnsecureWeb3Endpoint(endpoint string) *Client {
	c.Web3Endpoint = endpoint
	c.Web3Auth = ""
	return c
}

func (c *Client) WithUnsecureRestEndpoint(endpoint string) *Client {
	c.CosmosRestEndpoint = endpoint
	c.CosmosRestAuth = ""
	return c
}

func (c *Client) WithUnsecureTendermintEndpoint(endpoint string) *Client {
	c.TendermintRestEndpoint = endpoint
	c.TendermintRestAuth = ""
	return c
}

func (c *Client) WithSecureWeb3Endpoint(endpoint string, auth string) *Client {
	c.Web3Endpoint = endpoint
	c.Web3Auth = auth
	return c
}

func (c *Client) WithSecureRestEndpoint(endpoint string, auth string) *Client {
	c.CosmosRestEndpoint = endpoint
	c.CosmosRestAuth = auth
	return c
}

func (c *Client) WithSecureTendermintEndpoint(endpoint string, auth string) *Client {
	c.TendermintRestEndpoint = endpoint
	c.TendermintRestAuth = auth
	return c
}
