package gateway

import (
	"sync"
	"time"
)

// RateLimiter is the rate limiter for the gateway using the token bucket algorithm.
// 120 events can be sent per minute.
type rateLimiter struct {
	sync.Mutex
	tokens         uint
	maxTokens      uint
	nextRefill     time.Time
	refillInterval time.Duration
}

func newRateLimiter() *rateLimiter {
	return &rateLimiter{
		tokens:         120,
		maxTokens:      120,
		nextRefill:     time.Now().UTC().Add(1 * time.Minute),
		refillInterval: 1 * time.Minute,
	}
}

func (rl *rateLimiter) consume() {
	// Refill.
	rl.refill()

	// If there are tokens, deduct and return.
	if rl.tokens > 0 {
		rl.tokens--
		return
	}

	// Sleep until the next refill time.
	time.Sleep(rl.nextRefill.Sub(time.Now().UTC()))
}

func (rl *rateLimiter) refill() {
	// Get current time.
	currentTime := time.Now().UTC()

	// Check if current time is equal to or past the next refill time.
	if currentTime.Equal(rl.nextRefill) || currentTime.After(rl.nextRefill) {
		// Refill.
		rl.tokens = rl.maxTokens

		// Set next refill time.
		rl.nextRefill = time.Now().UTC().Add(rl.refillInterval)
	}
}
