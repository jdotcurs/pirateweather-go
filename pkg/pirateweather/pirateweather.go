package pirateweather

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://api.pirateweather.net/forecast"
)

type Client struct {
	APIKey      string
	HTTPClient  *http.Client
	BaseURL     string
	RateLimiter *RateLimiter
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
		BaseURL:     baseURL,
		RateLimiter: NewRateLimiter(1000), // Assuming a default limit of 1000 requests per day
	}
}
