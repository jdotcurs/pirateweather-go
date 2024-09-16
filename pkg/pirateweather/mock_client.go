package pirateweather

import (
	"net/http"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/models"
)

type MockClient struct {
	ForecastFunc          func(latitude, longitude float64, options ...ForecastOption) (*models.ForecastResponse, error)
	TimeMachineFunc       func(latitude, longitude float64, time time.Time, options ...ForecastOption) (*models.ForecastResponse, error)
	UpdateRateLimiterFunc func(headers http.Header)
}

func (m *MockClient) Forecast(latitude, longitude float64, options ...ForecastOption) (*models.ForecastResponse, error) {
	return m.ForecastFunc(latitude, longitude, options...)
}

func (m *MockClient) TimeMachine(latitude, longitude float64, time time.Time, options ...ForecastOption) (*models.ForecastResponse, error) {
	return m.TimeMachineFunc(latitude, longitude, time, options...)
}

func (m *MockClient) UpdateRateLimiter(headers http.Header) {
	if m.UpdateRateLimiterFunc != nil {
		m.UpdateRateLimiterFunc(headers)
	}
}
