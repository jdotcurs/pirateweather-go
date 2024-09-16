package pirateweather_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/stretchr/testify/assert"
)

func TestForecast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/test-api-key/45.420000,-75.690000", r.URL.Path)
		assert.Equal(t, "si", r.URL.Query().Get("units"))
		assert.Equal(t, "minutely", r.URL.Query().Get("exclude"))
		assert.Equal(t, "hourly", r.URL.Query().Get("extend"))
		assert.Equal(t, "2", r.URL.Query().Get("version"))

		w.Header().Set("Ratelimit-Limit", "1000")
		w.Header().Set("Ratelimit-Remaining", "999")
		w.Header().Set("Ratelimit-Reset", "1620000000")

		response := map[string]interface{}{
			"latitude":  45.42,
			"longitude": -75.69,
			"timezone":  "America/Toronto",
			"currently": map[string]interface{}{
				"time":        1620000000,
				"temperature": 20.5,
				"fireIndex":   5.0,
			},
			"hourly": map[string]interface{}{
				"data": make([]interface{}, 168),
			},
			"alerts": []map[string]interface{}{
				{
					"title":       "Test Alert",
					"description": "This is a test alert",
				},
			},
		}

		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := pirateweather.NewClient("test-api-key")
	client.BaseURL = server.URL

	forecast, err := client.Forecast(45.42, -75.69,
		pirateweather.WithUnits("si"),
		pirateweather.WithExclude([]string{"minutely"}),
		pirateweather.WithExtend("hourly"),
		pirateweather.WithVersion(2),
	)

	assert.NoError(t, err)
	assert.NotNil(t, forecast)
	assert.Equal(t, 45.42, forecast.Latitude)
	assert.Equal(t, -75.69, forecast.Longitude)
	assert.Equal(t, "America/Toronto", forecast.Timezone)
	assert.NotNil(t, forecast.Currently)
	assert.Equal(t, int64(1620000000), forecast.Currently.Time)
	assert.Equal(t, 20.5, forecast.Currently.Temperature)
	assert.Equal(t, 5.0, forecast.Currently.FireIndex)
	assert.Equal(t, 168, len(forecast.Hourly.Data))
	assert.Equal(t, 1, len(forecast.Alerts))
	assert.Equal(t, "Test Alert", forecast.Alerts[0].Title)
}
