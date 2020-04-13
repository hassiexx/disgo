package gateway

import (
	"context"

	"github.com/hassieswift621/disgo/json"
	"golang.org/x/xerrors"
	"nhooyr.io/websocket"
)

type payload struct {
	Op uint            `json:"op"`
	D  json.RawMessage `json:"d"`
	S  uint            `json:"s"`
	T  string          `json:"t"`
}

// ReadPayload reads a payload from the websocket.
func (s *Session) readPayload(ctx context.Context) (*payload, error) {
	// Read payload.
	msgType, r, err := s.ws.Reader(ctx)
	if err != nil {
		return nil, xerrors.Errorf("read payload: %w", err)
	}

	// Get payload.
	var payload payload
	if msgType == websocket.MessageBinary {
		// Decompress and unmarshal.

		if err = json.UnmarshalZlib(r, &payload); err != nil {
			return nil, xerrors.Errorf("decompress payload: %w", err)
		}
	} else {
		if err = json.Unmarshal(r, &payload); err != nil {
			return nil, xerrors.Errorf("unmarshal payload: %w", err)
		}
	}

	return &payload, nil
}

// SendPayload sends a payload on the websocket.
func (s *Session) sendPayload(ctx context.Context, payload interface{}) error {
	// Get writer.
	w, err := s.ws.Writer(ctx, websocket.MessageText)
	if err != nil {
		return xerrors.Errorf("failed to get websocket writer: %w", err)
	}

	// Marshal payload.
	if err = json.Marshal(w, payload); err != nil {
		_ = w.Close()
		return xerrors.Errorf("marshal payload: %w", err)
	}

	// Send payload by closing the writer.
	if err = w.Close(); err != nil {
		return xerrors.Errorf("send payload: %w", err)
	}

	return nil
}
