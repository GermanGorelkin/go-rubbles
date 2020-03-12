package go_rubbles

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (cl *Client) GetPredict(ctx context.Context, products []Product) (*PredictResponse, error) {
	r := RPCRequest{
		Method: "predict",
		Params: ProductsPredict{Products: products},
	}

	// create new request
	req, err := cl.httpClient.NewRequest("POST", "", r)
	if err != nil {
		return nil, fmt.Errorf("error create new request:%w", err)
	}

	// send request
	buf := new(bytes.Buffer)
	_, err = cl.httpClient.Do(ctx, req, buf)
	if err != nil {
		return nil, fmt.Errorf("error send request:%w", err)
	}

	// prepare data for decoding json
	b := buf.Bytes()
	ReplaceAll(b, byte('\''), byte('"'))
	/*
		 Replace True with true
			-> 'ready': True
	*/
	ReplaceAll(b, byte('T'), byte('t'))

	// decode json
	predict := new(PredictResponse)
	err = json.NewDecoder(buf).Decode(predict)
	if err != nil {
		return nil, fmt.Errorf("error decode json(body):%w\n\n%s", err, string(b))
	}

	return predict, nil
}

//
func ReplaceAll(s []byte, old, new byte) {
	for i := 0; i < len(s); i++ {
		if s[i] == old {
			s[i] = new
		}
	}
}
