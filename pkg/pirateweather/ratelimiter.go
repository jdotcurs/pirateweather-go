package pirateweather

import (
	"sync"
	"time"
)

// RateLimiter represents a rate limiter for API requests
type RateLimiter struct {
	mu           sync.Mutex
	limit        int
	remaining    int
	resetTime    time.Time
	lastResetted time.Time
}

// NewRateLimiter creates a new RateLimiter with the given limit
func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		limit:        limit,
		remaining:    limit,
		resetTime:    time.Now().Add(time.Hour * 24),
		lastResetted: time.Now(),
	}
}

// Allow checks if a request is allowed based on the current rate limit
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.remaining = rl.limit
		rl.resetTime = now.Add(time.Hour * 24)
		rl.lastResetted = now
	}

	if rl.remaining > 0 {
		rl.remaining--
		return true
	}
	return false
}

// UpdateFromHeaders updates the rate limiter based on the API response headers
func (rl *RateLimiter) UpdateFromHeaders(limit, remaining int, reset time.Time) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.limit = limit
	rl.remaining = remaining
	rl.resetTime = reset
}
