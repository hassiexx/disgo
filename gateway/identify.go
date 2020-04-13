package gateway

import (
	"context"
	"runtime"
	"sync"
	"time"

	"golang.org/x/xerrors"
)

var identifyLimiter = &identifyRateLimiter{}

type identifyPayload struct {
	Op uint         `json:"op"`
	D  identifyData `json:"d"`
}

type identifyData struct {
	Token              string             `json:"token"`
	Properties         identifyProperties `json:"properties"`
	Compress           bool               `json:"compress"`
	LargeThreshold     uint               `json:"large_threshold"`
	Shard              []uint             `json:"shard"`
	GuildSubscriptions bool               `json:"guild_subscriptions"`
}

type identifyProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

// Identify handles the identify payload.
func (s *Session) identify(ctx context.Context) error {
	// Create payload.
	payload := identifyPayload{
		Op: opcodeIdentify,
		D: identifyData{
			Token: s.token,
			Properties: identifyProperties{
				OS:      runtime.GOOS,
				Browser: "Disgo",
				Device:  "Disgo",
			},
			Compress:           true,
			LargeThreshold:     250,
			Shard:              []uint{s.shardID, s.shardCount},
			GuildSubscriptions: true,
		},
	}

	if err := s.sendPayload(ctx, payload); err != nil {
		return xerrors.Errorf("send identify payload: %w", err)
	}

	return nil
}

// IdentifyRateLimiter is the identify rate limiter.
// The API allows one identify every 5 seconds.
type identifyRateLimiter struct {
	lastIdentify time.Time
	sync.Mutex
}

// Acquire gets a token from the rate limiter.
// If a token is not available, it will sleep for the time difference.
func (rl *identifyRateLimiter) acquire() {
	rl.Lock()

	duration := time.Now().Sub(rl.lastIdentify).Seconds()

	if duration >= 5 {
		return
	}

	time.Sleep(time.Duration(5-duration) * time.Second)

	rl.lastIdentify = time.Now()

	rl.Unlock()
}
