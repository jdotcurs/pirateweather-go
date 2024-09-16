package utils

import (
	"fmt"
	"time"
)

// FormatTime formats a Unix timestamp to a human-readable string
func FormatTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

// ConvertTemperature converts temperature between Celsius and Fahrenheit
func ConvertTemperature(temp float64, fromUnit, toUnit string) (float64, error) {
	if fromUnit == toUnit {
		return temp, nil
	}

	switch {
	case fromUnit == "C" && toUnit == "F":
		return (temp * 9 / 5) + 32, nil
	case fromUnit == "F" && toUnit == "C":
		return (temp - 32) * 5 / 9, nil
	default:
		return 0, fmt.Errorf("unsupported temperature conversion: %s to %s", fromUnit, toUnit)
	}
}

// ConvertUnit converts between different units of measurement
func ConvertUnit(value float64, fromUnit, toUnit string) (float64, error) {
	switch {
	case fromUnit == "km" && toUnit == "mi":
		return value * 0.621371, nil
	case fromUnit == "mi" && toUnit == "km":
		return value * 1.60934, nil
	case fromUnit == "m/s" && toUnit == "km/h":
		return value * 3.6, nil
	case fromUnit == "km/h" && toUnit == "m/s":
		return value / 3.6, nil
	case fromUnit == "km/h" && toUnit == "mph":
		return value * 0.621371, nil
	case fromUnit == "mph" && toUnit == "km/h":
		return value * 1.60934, nil
	case fromUnit == "hPa" && toUnit == "inHg":
		return value * 0.02953, nil
	case fromUnit == "inHg" && toUnit == "hPa":
		return value / 0.02953, nil
	default:
		return 0, fmt.Errorf("unsupported unit conversion: %s to %s", fromUnit, toUnit)
	}
}
