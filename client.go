package go_rubbles

import (
	"github.com/germangorelkin/http-client"
	"net/http"
	"time"
)

const (
	tokenType = "Bearer"
)

type Client struct {
	httpClient *http_client.Client
}

func NewClient(cfg ClientConfig) (*Client, error) {
	httpClient := &http.Client{
		Timeout: cfg.HTTPClientTimeout,
	}
	cl, err := http_client.New(httpClient,
		http_client.SetBaseURL(cfg.BaseURL),
		http_client.SetAuthorization(cfg.Token, tokenType))
	if err != nil {
		return nil, err
	}

	return &Client{httpClient: cl}, nil
}

type ClientConfig struct {
	HTTPClientTimeout time.Duration
	BaseURL           string
	Token             string
}

type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      int         `json:"id,omitempty"`
	JSONRPC string      `json:"jsonrpc,omitempty"`
}
