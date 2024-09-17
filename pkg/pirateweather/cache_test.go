package pirateweather_test

import (
	"testing"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/models"
	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/stretchr/testify/require"
)

func TestForecastCaching(t *testing.T) {
	mockTime := time.Now()
	pirateweather.SetTimeNow(func() time.Time {
		return mockTime
	})
	defer pirateweather.ResetTimeNow()

	callCount := 0
	mockClient := &pirateweather.MockClient{
		ForecastFunc: func(latitude, longitude float64, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			callCount++
			return &models.ForecastResponse{
				Latitude:  latitude,
				Longitude: longitude,
				Currently: &models.DataPoint{
					Time:        mockTime.Unix(),
					Temperature: 20.5,
				},
			}, nil
		},
	}

	// First call
	forecast1, err := mockClient.Forecast(45.42, -75.69)
	require.NoError(t, err)
	require.Equal(t, 1, callCount)

	// Second call (should use cache)
	forecast2, err := mockClient.Forecast(45.42, -75.69)
	require.NoError(t, err)
	require.Equal(t, 1, callCount)
	require.Equal(t, forecast1, forecast2)

	// Advance mock time by 16 minutes (just past cache expiration)
	mockTime = mockTime.Add(16 * time.Minute)

	// Third call (should not use cache)
	forecast3, err := mockClient.Forecast(45.42, -75.69)
	require.NoError(t, err)
	require.Equal(t, 2, callCount)
	require.NotEqual(t, forecast1.Currently.Time, forecast3.Currently.Time)
}

func TestTimeMachineCaching(t *testing.T) {
	mockTime := time.Now()
	pirateweather.SetTimeNow(func() time.Time {
		return mockTime
	})
	defer pirateweather.ResetTimeNow()

	callCount := 0
	mockClient := &pirateweather.MockClient{
		TimeMachineFunc: func(latitude, longitude float64, time time.Time, options ...pirateweather.ForecastOption) (*models.ForecastResponse, error) {
			callCount++
			return &models.ForecastResponse{
				Latitude:  latitude,
				Longitude: longitude,
				Currently: &models.DataPoint{
					Time:        mockTime.Unix(),
					Temperature: 20.5,
				},
			}, nil
		},
	}

	// First call
	forecast1, err := mockClient.TimeMachine(45.42, -75.69, mockTime)
	require.NoError(t, err)
	require.Equal(t, 1, callCount)

	// Second call (should use cache)
	forecast2, err := mockClient.TimeMachine(45.42, -75.69, mockTime)
	require.NoError(t, err)
	require.Equal(t, 1, callCount)
	require.Equal(t, forecast1, forecast2)

	// Advance mock time by 61 minutes (just past cache expiration)
	mockTime = mockTime.Add(61 * time.Minute)

	// Third call (should not use cache)
	forecast3, err := mockClient.TimeMachine(45.42, -75.69, mockTime)
	require.NoError(t, err)
	require.Equal(t, 2, callCount)
	require.NotEqual(t, forecast1.Currently.Time, forecast3.Currently.Time)
}
