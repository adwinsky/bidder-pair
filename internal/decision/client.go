package decision

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ClientConfig struct {
	BaseURL string
	Client  *http.Client
}

type Client struct {
	baseURL string
	c       *http.Client
}

func NewClient(cfg ClientConfig) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		c:       cfg.Client,
	}
}

func (c *Client) Decide(ctx context.Context, req DecisionRequest) (DecisionResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return DecisionResponse{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/decision", bytes.NewReader(b))
	if err != nil {
		return DecisionResponse{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.c.Do(httpReq)
	if err != nil {
		return DecisionResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body)
		return DecisionResponse{}, fmt.Errorf("decision status=%d body=%s", resp.StatusCode, string(raw))
	}

	var out DecisionResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return DecisionResponse{}, err
	}
	return out, nil
}
