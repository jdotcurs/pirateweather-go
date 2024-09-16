package pirateweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/models"
)

// TimeMachine retrieves historical weather data for a given location and time
func (c *Client) TimeMachine(latitude, longitude float64, timestamp time.Time, options ...ForecastOption) (*models.ForecastResponse, error) {
	url := fmt.Sprintf("%s/%s/%f,%f,%d", c.BaseURL, c.APIKey, latitude, longitude, timestamp.Unix())

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
