package pirateweather_test

import (
	"net/http"
	"testing"

	"github.com/jdotcurs/pirateweather-go/pkg/models"
	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/stretchr/testify/require"
)

func TestForecast(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return &models.ForecastResponse{
				Latitude:  45.42,
				Longitude: -75.69,
				Timezone:  "America/Toronto",
				Currently: &models.DataPoint{
					Time:        1620000000,
					Temperature: 20.5,
					FireIndex:   5.0,
				},
				Hourly: &models.DataBlock{
					Data: make([]models.DataPoint, 168),
				},
				Alerts: []models.Alert{
					{
						Title:       "Test Alert",
						Description: "This is a test alert",
					},
				},
			}, nil
		},
	}

	forecast, err := mockClient.Forecast(45.42, -75.69,
		pirateweather.WithUnits("si"),
		pirateweather.WithExclude([]string{"minutely"}),
		pirateweather.WithExtend("hourly"),
		pirateweather.WithVersion(2),
	)

	require.NoError(t, err)
	require.NotNil(t, forecast)
	require.Equal(t, 45.42, forecast.Latitude)
	require.Equal(t, -75.69, forecast.Longitude)
	require.Equal(t, "America/Toronto", forecast.Timezone)
	require.NotNil(t, forecast.Currently)
	require.Equal(t, int64(1620000000), forecast.Currently.Time)
	require.Equal(t, 20.5, forecast.Currently.Temperature)
	require.Equal(t, 5.0, forecast.Currently.FireIndex)
	require.Equal(t, 168, len(forecast.Hourly.Data))
	require.Equal(t, 1, len(forecast.Alerts))
	require.Equal(t, "Test Alert", forecast.Alerts[0].Title)
}

func TestForecastRateLimit(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return nil, &pirateweather.RateLimitError{
				Message: "rate limit exceeded",
			}
		},
	}

	_, err := mockClient.Forecast(45.42, -75.69)
	require.Error(t, err)
	require.Contains(t, err.Error(), "rate limit exceeded")
}

func TestForecastAPIFailure(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return nil, &pirateweather.APIError{
				Message: "API request failed after 5 attempts",
			}
		},
	}

	_, err := mockClient.Forecast(45.42, -75.69)
	require.Error(t, err)
	require.Contains(t, err.Error(), "API request failed after")
}

func TestForecastVersion2(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return &models.ForecastResponse{
				Latitude:  45.42,
				Longitude: -75.69,
				Timezone:  "America/Toronto",
				Currently: &models.DataPoint{
					Time:        1620000000,
					Temperature: 20.5,
					FireIndex:   5.0,
					Smoke:       10.0,
					DawnTime:    1619950000,
					DuskTime:    1620000000,
				},
			}, nil
		},
	}

	forecast, err := mockClient.Forecast(45.42, -75.69, pirateweather.WithVersion(2))
	require.NoError(t, err)
	require.NotNil(t, forecast)
	require.Equal(t, 5.0, forecast.Currently.FireIndex)
	require.Equal(t, 10.0, forecast.Currently.Smoke)
	require.Equal(t, int64(1619950000), forecast.Currently.DawnTime)
	require.Equal(t, int64(1620000000), forecast.Currently.DuskTime)
}

func TestForecastWithAllOptions(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return &models.ForecastResponse{
				Latitude:  45.42,
				Longitude: -75.69,
				Timezone:  "America/Toronto",
				Currently: &models.DataPoint{
					Time:        1620000000,
					Temperature: 68.0,
					FireIndex:   5.0,
					Smoke:       10.0,
				},
			}, nil
		},
	}

	forecast, err := mockClient.Forecast(45.42, -75.69,
		pirateweather.WithUnits("us"),
		pirateweather.WithExclude([]string{"minutely", "hourly"}),
		pirateweather.WithExtend("daily"),
		pirateweather.WithVersion(2),
	)

	require.NoError(t, err)
	require.NotNil(t, forecast)
	require.Equal(t, 45.42, forecast.Latitude)
	require.Equal(t, -75.69, forecast.Longitude)
	require.Equal(t, "America/Toronto", forecast.Timezone)
	require.NotNil(t, forecast.Currently)
	require.Equal(t, int64(1620000000), forecast.Currently.Time)
	require.Equal(t, 68.0, forecast.Currently.Temperature)
	require.Equal(t, 5.0, forecast.Currently.FireIndex)
	require.Equal(t, 10.0, forecast.Currently.Smoke)
}

func TestForecastInvalidJSON(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return nil, &pirateweather.JSONError{
				Message: "error decoding response: invalid character 'i' looking for beginning of value",
			}
		},
	}

	_, err := mockClient.Forecast(45.42, -75.69)
	require.Error(t, err)
	require.Contains(t, err.Error(), "error decoding response")
}

func TestForecastWithMockClient(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return &models.ForecastResponse{
				Latitude:  45.42,
				Longitude: -75.69,
				Timezone:  "America/Toronto",
				Currently: &models.DataPoint{
					Time:        1620000000,
					Temperature: 20.5,
					FireIndex:   5.0,
				},
				Hourly: &models.DataBlock{
					Data: make([]models.DataPoint, 168),
				},
				Alerts: []models.Alert{
					{
						Title:       "Test Alert",
						Description: "This is a test alert",
					},
				},
			}, nil
		},
	}

	forecast, err := mockClient.Forecast(45.42, -75.69,
		pirateweather.WithUnits("si"),
		pirateweather.WithExclude([]string{"minutely"}),
		pirateweather.WithExtend("hourly"),
		pirateweather.WithVersion(2),
	)

	require.NoError(t, err)
	require.NotNil(t, forecast)
	require.Equal(t, 45.42, forecast.Latitude)
	require.Equal(t, -75.69, forecast.Longitude)
	require.Equal(t, "America/Toronto", forecast.Timezone)
	require.NotNil(t, forecast.Currently)
	require.Equal(t, int64(1620000000), forecast.Currently.Time)
	require.Equal(t, 20.5, forecast.Currently.Temperature)
	require.Equal(t, 5.0, forecast.Currently.FireIndex)
	require.Equal(t, 168, len(forecast.Hourly.Data))
	require.Equal(t, 1, len(forecast.Alerts))
	require.Equal(t, "Test Alert", forecast.Alerts[0].Title)
}

func TestMockClientUpdateRateLimiter(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		UpdateRateLimiterFunc: func(headers http.Header) {
			// Simulate updating rate limiter
			headers.Set("X-RateLimit-Limit", "100")
			headers.Set("X-RateLimit-Remaining", "99")
			headers.Set("X-RateLimit-Reset", "1620000000")
		},
	}

	headers := make(http.Header)
	mockClient.UpdateRateLimiter(headers)

	require.Equal(t, "100", headers.Get("X-RateLimit-Limit"))
	require.Equal(t, "99", headers.Get("X-RateLimit-Remaining"))
	require.Equal(t, "1620000000", headers.Get("X-RateLimit-Reset"))
}
