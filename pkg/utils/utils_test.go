package utils_test

import (
	"testing"

	"github.com/jdotcurs/pirateweather-go/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestFormatTime(t *testing.T) {
	timestamp := int64(1620000000)
	expected := "2021-05-03T12:00:00+12:00"

	result := utils.FormatTime(timestamp)
	require.Equal(t, expected, result)
}

func TestConvertTemperature(t *testing.T) {
	testCases := []struct {
		temp     float64
		fromUnit string
		toUnit   string
		expected float64
	}{
		{0, "C", "F", 32},
		{32, "F", "C", 0},
		{100, "C", "C", 100},
		{100, "F", "F", 100},
	}

	for _, tc := range testCases {
		result, err := utils.ConvertTemperature(tc.temp, tc.fromUnit, tc.toUnit)
		require.NoError(t, err)
		require.InDelta(t, tc.expected, result, 0.01)
	}

	_, err := utils.ConvertTemperature(0, "K", "C")
	require.Error(t, err)
}

func TestConvertUnit(t *testing.T) {
	testCases := []struct {
		value    float64
		fromUnit string
		toUnit   string
		expected float64
	}{
		{1, "km", "mi", 0.621371},
		{1, "mi", "km", 1.60934},
		{1, "m/s", "km/h", 3.6},
		{3.6, "km/h", "m/s", 1},
		{100, "km/h", "mph", 62.1371},
		{62.1371, "mph", "km/h", 100},
		{1000, "hPa", "inHg", 29.53},
		{29.53, "inHg", "hPa", 1000},
	}

	for _, tc := range testCases {
		result, err := utils.ConvertUnit(tc.value, tc.fromUnit, tc.toUnit)
		require.NoError(t, err)
		require.InDelta(t, tc.expected, result, 0.01)
	}

	_, err := utils.ConvertUnit(1, "kg", "lb")
	require.Error(t, err)
}
