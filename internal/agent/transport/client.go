package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/asim9115/incidentix/internal/agent/collector"
)

type Client struct {
	httpClient *http.Client
	baseUrl string
	//token string
}


func New(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) SendMetrics(ctx context.Context, metrics collector.Metrics) error {
	data, err := json.Marshal(metrics)
	if err != nil {
		return fmt.Errorf("failed converting to json : %v", err)
	}

	resp, err := c.httpClient.Post(c.baseUrl+"/v1/agents/metrics", "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed sending metrics: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
    	return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	return nil
}