package pirateweather_test

import (
	"testing"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/models"
	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/stretchr/testify/require"
)

func TestTimeMachine(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		TimeMachineFunc: func(latitude, longitude float64, time time.Time, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return &models.ForecastResponse{
				Latitude:  45.42,
				Longitude: -75.69,
				Timezone:  "America/Toronto",
				Currently: &models.DataPoint{
					Time:        1620000000,
					Temperature: 18.5,
				},
				Hourly: &models.DataBlock{
					Data: make([]models.DataPoint, 24),
				},
			}, nil
		},
	}

	timeMachineTimestamp := time.Unix(1620000000, 0)
	forecast, err := mockClient.TimeMachine(45.42, -75.69, timeMachineTimestamp, pirateweather.WithUnits("si"))

	require.NoError(t, err)
	require.NotNil(t, forecast)
	require.Equal(t, 45.42, forecast.Latitude)
	require.Equal(t, -75.69, forecast.Longitude)
	require.Equal(t, "America/Toronto", forecast.Timezone)
	require.NotNil(t, forecast.Currently)
	require.Equal(t, int64(1620000000), forecast.Currently.Time)
	require.Equal(t, 18.5, forecast.Currently.Temperature)
	require.Equal(t, 24, len(forecast.Hourly.Data))
}

func TestTimeMachineWithAllOptions(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		TimeMachineFunc: func(latitude, longitude float64, time time.Time, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return &models.ForecastResponse{
				Latitude:  45.42,
				Longitude: -75.69,
				Timezone:  "America/Toronto",
				Currently: &models.DataPoint{
					Time:        1620000000,
					Temperature: 68.0,
				},
			}, nil
		},
	}

	timeMachineTimestamp := time.Unix(1620000000, 0)
	forecast, err := mockClient.TimeMachine(45.42, -75.69, timeMachineTimestamp,
		pirateweather.WithUnits("us"),
		pirateweather.WithExclude([]string{"minutely", "hourly"}),
	)

	require.NoError(t, err)
	require.NotNil(t, forecast)
	require.Equal(t, 45.42, forecast.Latitude)
	require.Equal(t, -75.69, forecast.Longitude)
	require.Equal(t, "America/Toronto", forecast.Timezone)
	require.NotNil(t, forecast.Currently)
	require.Equal(t, int64(1620000000), forecast.Currently.Time)
	require.Equal(t, 68.0, forecast.Currently.Temperature)
}

func TestTimeMachineInvalidJSON(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		TimeMachineFunc: func(latitude, longitude float64, time time.Time, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return nil, &pirateweather.JSONError{
				Message: "error decoding response: invalid character 'i' looking for beginning of value",
			}
		},
	}

	timeMachineTimestamp := time.Unix(1620000000, 0)
	_, err := mockClient.TimeMachine(45.42, -75.69, timeMachineTimestamp)
	require.Error(t, err)
	require.Contains(t, err.Error(), "error decoding response")
}

func TestTimeMachineWithMockClient(t *testing.T) {
	mockClient := &pirateweather.MockClient{
		TimeMachineFunc: func(latitude, longitude float64, time time.Time, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			return &models.ForecastResponse{
				Latitude:  45.42,
				Longitude: -75.69,
				Timezone:  "America/Toronto",
				Currently: &models.DataPoint{
					Time:        1620000000,
					Temperature: 18.5,
				},
				Hourly: &models.DataBlock{
					Data: make([]models.DataPoint, 24),
				},
			}, nil
		},
	}

	timeMachineTimestamp := time.Unix(1620000000, 0)
	forecast, err := mockClient.TimeMachine(45.42, -75.69, timeMachineTimestamp, pirateweather.WithUnits("si"))

	require.NoError(t, err)
	require.NotNil(t, forecast)
	require.Equal(t, 45.42, forecast.Latitude)
	require.Equal(t, -75.69, forecast.Longitude)
	require.Equal(t, "America/Toronto", forecast.Timezone)
	require.NotNil(t, forecast.Currently)
	require.Equal(t, int64(1620000000), forecast.Currently.Time)
	require.Equal(t, 18.5, forecast.Currently.Temperature)
	require.Equal(t, 24, len(forecast.Hourly.Data))
}
