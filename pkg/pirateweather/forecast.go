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
// It takes latitude and longitude as parameters, along with optional ForecastOptions
func (c *Client) Forecast(latitude, longitude float64, options ...ForecastOption) (*models.ForecastResponse, error) {
	cacheKey := fmt.Sprintf("forecast:%f:%f:%v", latitude, longitude, options)
	if cachedForecast, found := c.Cache.Get(cacheKey); found {
		return cachedForecast.(*models.ForecastResponse), nil
	}

	url := fmt.Sprintf("%s/%s/%f,%f", c.BaseURL, c.APIKey, latitude, longitude)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Apply all provided options to the request
	for _, option := range options {
		option(req)
	}

	var resp *http.Response
	var forecast models.ForecastResponse

	// Retry logic for handling transient errors
	for i := 0; i < maxRetries; i++ {
		if !c.RateLimiter.Allow() {
			return nil, &RateLimitError{
				Message: "rate limit exceeded",
			}
		}

		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}
		defer resp.Body.Close()

		// Handle different response status codes
		switch resp.StatusCode {
		case http.StatusOK:
			if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
				return nil, &JSONError{
					Message: fmt.Sprintf("error decoding response: %v", err),
				}
			}
			c.updateRateLimiter(resp.Header)
			c.Cache.Set(cacheKey, &forecast, time.Hour) // Cache for 1 hour
			return &forecast, nil
		case http.StatusBadRequest:
			return nil, fmt.Errorf("bad request: invalid latitude or longitude")
		case http.StatusUnauthorized:
			return nil, fmt.Errorf("unauthorized: invalid API key or insufficient permissions")
		case http.StatusNotFound:
			return nil, fmt.Errorf("not found: invalid route or missing latitude/longitude")
		case http.StatusTooManyRequests:
			return nil, fmt.Errorf("rate limit exceeded: API key has hit the quota for the month")
		case http.StatusInternalServerError:
			if i == maxRetries-1 {
				return nil, fmt.Errorf("API request failed after %d retries: internal server error", maxRetries)
			}
			time.Sleep(retryDelay)
		default:
			return nil, &APIError{
				Message: fmt.Sprintf("API request failed with unexpected status code: %d", resp.StatusCode),
			}
		}
	}

	return nil, fmt.Errorf("API request failed after %d retries", maxRetries)
}

// updateRateLimiter updates the rate limiter based on the response headers
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
