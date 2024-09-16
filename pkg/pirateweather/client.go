// Package pirateweather provides a client for the Pirate Weather API
package pirateweather

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://api.pirateweather.net/forecast"
)

// Client represents a Pirate Weather API client
type Client struct {
	APIKey      string
	HTTPClient  *http.Client
	BaseURL     string
	RateLimiter *RateLimiter
}

// NewClient creates a new Pirate Weather API client with the given API key
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
		BaseURL:     baseURL,
		RateLimiter: NewRateLimiter(10000 / 30), // Default limit of 1000 requests per day
	}
}
