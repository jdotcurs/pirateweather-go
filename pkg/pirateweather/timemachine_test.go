package pirateweather_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/stretchr/testify/assert"
)

func TestTimeMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/test-api-key/45.420000,-75.690000,1620000000", r.URL.Path)
		assert.Equal(t, "si", r.URL.Query().Get("units"))

		w.Header().Set("Ratelimit-Limit", "1000")
		w.Header().Set("Ratelimit-Remaining", "999")
		w.Header().Set("Ratelimit-Reset", "1620000000")

		response := map[string]interface{}{
			"latitude":  45.42,
			"longitude": -75.69,
			"timezone":  "America/Toronto",
			"currently": map[string]interface{}{
				"time":        1620000000,
				"temperature": 18.5,
			},
			"hourly": map[string]interface{}{
				"data": make([]interface{}, 24),
			},
		}

		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := pirateweather.NewClient("test-api-key")
	client.BaseURL = server.URL

	timeMachineTimestamp := time.Unix(1620000000, 0)
	forecast, err := client.TimeMachine(45.42, -75.69, timeMachineTimestamp, pirateweather.WithUnits("si"))

	assert.NoError(t, err)
	assert.NotNil(t, forecast)
	assert.Equal(t, 45.42, forecast.Latitude)
	assert.Equal(t, -75.69, forecast.Longitude)
	assert.Equal(t, "America/Toronto", forecast.Timezone)
	assert.NotNil(t, forecast.Currently)
	assert.Equal(t, int64(1620000000), forecast.Currently.Time)
	assert.Equal(t, 18.5, forecast.Currently.Temperature)
	assert.Equal(t, 24, len(forecast.Hourly.Data))
}
