package gateway

import (
	"context"
	"encoding/json"
	"fmt"

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
	msgType, msg, err := s.ws.Read(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to read payload: %w", err)
	}

	// Get payload data.
	var data []byte
	if msgType == websocket.MessageBinary {
		// Decompress payload.
		fmt.Println(msg)
		data, err = decompress(msg)
		if err != nil {
			return nil, xerrors.Errorf("failed to decompress payload: %w", err)
		}
	} else {
		data = msg
	}

	// Unmarshal payload into json.
	var payload payload
	if err = unmarshal(data, &payload); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal payload: %w", err)
	}

	return &payload, nil
}

// SendPayload sends a payload on the websocket.
func (s *Session) sendPayload(ctx context.Context, payload interface{}) error {
	// Unmarshal payload.
	data, err := marshal(payload)
	if err != nil {
		return xerrors.Errorf("failed to marshal payload: %w", err)
	}

	// Send payload.
	err = s.ws.Write(ctx, websocket.MessageText, data)
	if err != nil {
		return xerrors.Errorf("failed to send payload: %w", err)
	}

	return nil
}
