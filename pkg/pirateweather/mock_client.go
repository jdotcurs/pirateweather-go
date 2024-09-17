package pirateweather

import (
	"fmt"
	"net/http"
	mocktime "time"

	"github.com/jdotcurs/pirateweather-go/pkg/models"
)

type MockClient struct {
	ForecastFunc          func(latitude, longitude float64, options ...ForecastOption) (*models.ForecastResponse, error)
	TimeMachineFunc       func(latitude, longitude float64, time mocktime.Time, options ...ForecastOption) (*models.ForecastResponse, error)
	UpdateRateLimiterFunc func(headers http.Header)
	Cache                 *Cache
}

func (m *MockClient) initCache() {
	if m.Cache == nil {
		m.Cache = NewCache()
	}
}

func (m *MockClient) Forecast(latitude, longitude float64, options ...ForecastOption) (*models.ForecastResponse, error) {
	m.initCache()
	cacheKey := fmt.Sprintf("forecast:%f:%f:%v", latitude, longitude, options)
	if cachedForecast, found := m.Cache.Get(cacheKey); found {
		return cachedForecast.(*models.ForecastResponse), nil
	}
	forecast, err := m.ForecastFunc(latitude, longitude, options...)
	if err == nil {
		m.Cache.Set(cacheKey, forecast, 15*mocktime.Minute)
	}
	return forecast, err
}

func (m *MockClient) TimeMachine(latitude, longitude float64, time mocktime.Time, options ...ForecastOption) (*models.ForecastResponse, error) {
	m.initCache()
	cacheKey := fmt.Sprintf("timemachine:%f:%f:%d:%v", latitude, longitude, time.Unix(), options)
	if cachedForecast, found := m.Cache.Get(cacheKey); found {
		return cachedForecast.(*models.ForecastResponse), nil
	}
	forecast, err := m.TimeMachineFunc(latitude, longitude, time, options...)
	if err == nil {
		m.Cache.Set(cacheKey, forecast, 1*mocktime.Hour)
	}
	return forecast, err
}

func (m *MockClient) UpdateRateLimiter(headers http.Header) {
	if m.UpdateRateLimiterFunc != nil {
		m.UpdateRateLimiterFunc(headers)
	}
}
