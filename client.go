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
	http *http_client.Client
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

	return &Client{http: cl}, nil
}

type ClientConfig struct {
	HTTPClientTimeout time.Duration
	BaseURL           string
	Token             string
}
