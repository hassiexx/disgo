package gateway

import (
	"context"

	"golang.org/x/xerrors"
)

type resumePayload struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       uint64 `json:"seq"`
}

func (s *Session) resume(ctx context.Context) error {
	// Create payload.
	payload := resumePayload{
		Token:     s.token,
		SessionID: s.sessionID,
		Seq:       s.sequence,
	}

	// Send payload.
	if err := s.sendPayload(ctx, payload); err != nil {
		return xerrors.Errorf("failed to send resume payload: %w", err)
	}

	return nil
}
