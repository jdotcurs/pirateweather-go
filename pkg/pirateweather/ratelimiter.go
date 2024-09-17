package pirateweather

import (
	"sync"
	"time"
)

// RateLimiter represents a rate limiter for API requests
type RateLimiter struct {
	mu           sync.Mutex
	limit        int
	tokens       float64
	lastRefilled time.Time
	refillRate   float64
}

// NewRateLimiter creates a new RateLimiter with the given limit
func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		limit:        limit,
		tokens:       float64(limit),
		lastRefilled: time.Now(),
		refillRate:   float64(limit) / (30 * 24 * 60 * 60), // Tokens per second for a month
	}
}

// Allow checks if a request is allowed based on the current rate limit
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefilled).Seconds()
	rl.tokens = min(float64(rl.limit), rl.tokens+elapsed*rl.refillRate)
	rl.lastRefilled = now

	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}
	return false
}

// UpdateFromHeaders updates the rate limiter based on the API response headers
func (rl *RateLimiter) UpdateFromHeaders(limit, remaining int, reset time.Time) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.limit = limit
	rl.tokens = float64(remaining)
	rl.lastRefilled = time.Now()
	rl.refillRate = float64(limit) / time.Until(reset).Seconds()
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
