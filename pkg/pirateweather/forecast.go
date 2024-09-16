package pirateweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/models"
)

const (
	maxRetries = 3
	retryDelay = time.Second * 2
)

// Forecast retrieves the weather forecast for a given location
func (c *Client) Forecast(latitude, longitude float64, options ...ForecastOption) (*models.ForecastResponse, error) {
	url := fmt.Sprintf("%s/%s/%f,%f", c.BaseURL, c.APIKey, latitude, longitude)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for _, option := range options {
		option(req)
	}

	var resp *http.Response
	var forecast models.ForecastResponse

	for i := 0; i < maxRetries; i++ {
		if !c.RateLimiter.Allow() {
			return nil, fmt.Errorf("rate limit exceeded")
		}

		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			break
		}

		if resp.StatusCode != http.StatusInternalServerError {
			return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
		}

		time.Sleep(retryDelay)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed after %d retries", maxRetries)
	}

	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	c.updateRateLimiter(resp.Header)

	return &forecast, nil
}

func (c *Client) updateRateLimiter(headers http.Header) {
	limit, _ := strconv.Atoi(headers.Get("Ratelimit-Limit"))
	remaining, _ := strconv.Atoi(headers.Get("Ratelimit-Remaining"))
	reset, _ := strconv.ParseInt(headers.Get("Ratelimit-Reset"), 10, 64)

	c.RateLimiter.UpdateFromHeaders(limit, remaining, time.Unix(reset, 0))
}

// ForecastOption represents an option for the Forecast method
type ForecastOption func(*http.Request)

// WithUnits sets the units for the forecast request
func WithUnits(units string) ForecastOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("units", units)
		req.URL.RawQuery = q.Encode()
	}
}

// WithExclude sets the exclude parameter for the forecast request
func WithExclude(exclude []string) ForecastOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("exclude", strings.Join(exclude, ","))
		req.URL.RawQuery = q.Encode()
	}
}

// WithExtend sets the extend parameter for the forecast request
func WithExtend(extend string) ForecastOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("extend", extend)
		req.URL.RawQuery = q.Encode()
	}
}

// WithVersion sets the version parameter for the forecast request
func WithVersion(version int) ForecastOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("version", strconv.Itoa(version))
		req.URL.RawQuery = q.Encode()
	}
}
