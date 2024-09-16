package pirateweather

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu           sync.Mutex
	limit        int
	remaining    int
	resetTime    time.Time
	lastResetted time.Time
}

func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		limit:        limit,
		remaining:    limit,
		resetTime:    time.Now().Add(time.Hour * 24),
		lastResetted: time.Now(),
	}
}

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

func (rl *RateLimiter) UpdateFromHeaders(limit, remaining int, reset time.Time) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.limit = limit
	rl.remaining = remaining
	rl.resetTime = reset
}
