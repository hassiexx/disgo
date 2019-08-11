package gateway

import (
	"context"
	"runtime"

	"golang.org/x/xerrors"
)

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
		Op: uint(OpcodeIdentify),
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
		return xerrors.Errorf("failed to send identify payload: %w", err)
	}

	return nil
}
