package geocoding

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GeocodingResult struct {
	DisplayName string `json:"display_name"`
	Address     struct {
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
	} `json:"address"`
}

func ReverseGeocode(latitude, longitude float64) (*GeocodingResult, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", latitude, longitude)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "PirateWeatherGoSDK/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var result GeocodingResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &result, nil
}

type ForwardGeocodingResult struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

func ForwardGeocode(address string) (*ForwardGeocodingResult, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s&limit=1", url.QueryEscape(address))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "PirateWeatherGoSDK/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var results []ForwardGeocodingResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results found for the given address")
	}

	return &results[0], nil
}
