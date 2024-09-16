package pirateweather_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	client := pirateweather.NewClient(apiKey)

	require.NotNil(t, client)
	require.Equal(t, apiKey, client.APIKey)
	require.NotNil(t, client.HTTPClient)
	require.Equal(t, "https://api.pirateweather.net/forecast", client.BaseURL)
	require.NotNil(t, client.RateLimiter)
}

func TestClientWithCustomHTTPClient(t *testing.T) {
	apiKey := "test-api-key"
	customHTTPClient := &http.Client{
		Timeout: time.Second * 30,
	}

	client := pirateweather.NewClient(apiKey)
	client.HTTPClient = customHTTPClient

	require.Equal(t, customHTTPClient, client.HTTPClient)
}

func TestClientWithCustomBaseURL(t *testing.T) {
	apiKey := "test-api-key"
	customBaseURL := "https://custom.pirateweather.net/forecast"

	client := pirateweather.NewClient(apiKey)
	client.BaseURL = customBaseURL

	require.Equal(t, customBaseURL, client.BaseURL)
}
