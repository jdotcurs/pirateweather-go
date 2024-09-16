package pirateweather

import (
	"fmt"
)

// APIError represents an error returned by the Pirate Weather API
type APIError struct {
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: %s", e.Message)
}

// RateLimitError represents a rate limit error
type RateLimitError struct {
	Message string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("Rate Limit Error: %s", e.Message)
}

// JSONError represents an error that occurred while parsing JSON
type JSONError struct {
	Message string
}

func (e *JSONError) Error() string {
	return fmt.Sprintf("JSON Error: %s", e.Message)
}
