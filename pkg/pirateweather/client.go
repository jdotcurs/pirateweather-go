// Package pirateweather provides a client for the Pirate Weather API
package pirateweather

import (
	"net/http"
	"time"
)

var timeNow = time.Now

func SetTimeNow(f func() time.Time) {
	timeNow = f
}

func ResetTimeNow() {
	timeNow = time.Now
}

const (
	baseURL = "https://api.pirateweather.net/forecast"
)

// Client represents a Pirate Weather API client
type Client struct {
	APIKey      string
	HTTPClient  *http.Client
	BaseURL     string
	RateLimiter *RateLimiter
	Cache       *Cache
}

// NewClient creates a new Pirate Weather API client with the given API key
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
		BaseURL:     baseURL,
		RateLimiter: NewRateLimiter(10000 / 30), // Default limit of 10000 requests per month
		Cache:       NewCache(),
	}
}
