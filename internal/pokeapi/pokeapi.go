package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Next     string
	Previous string
}

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) LocationAreas(config *Config, endpoint string) (*LocationAreaResp, error) {

	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("unexpected status %d from API", resp.StatusCode)
	}

	var response LocationAreaResp
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(response.Results) == 0 {
		fmt.Println("No locations found.")
		return nil, nil
	}

	config.Next = response.Next
	config.Previous = response.Previous

	return &response, nil
}
