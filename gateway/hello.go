package gateway

import (
	"context"
	"github.com/hassieswift621/disgo/json"
	"time"

	"golang.org/x/xerrors"
)

type helloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// Hello handles the hello payload.
func (s *Session) hello(ctx context.Context) error {
	// Read payload.
	payload, err := s.readPayload(ctx)
	if err != nil {
		return xerrors.Errorf("get hello payload: %w", err)
	}

	// We should get an opcode 10.
	if payload.Op != 10 {
		return xerrors.Errorf("get hello payload: got opcode %d instead", payload.Op)
	}

	// Unmarshal payload data.
	var helloData helloData
	if err = json.UnmarshalRaw(payload.D, &helloData); err != nil {
		return xerrors.Errorf("failed to unmarshal hello payload: %w", err)
	}

	// Store heartbeat interval.
	s.heartbeatState = &heartbeatState{
		Interval: time.Duration(helloData.HeartbeatInterval) * time.Millisecond,
	}

	return nil
}
