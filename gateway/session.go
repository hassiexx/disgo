package gateway

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/xerrors"

	"nhooyr.io/websocket"
)

const gatewayURL = "wss://gateway.discord.gg/?v=6&encoding=json&compress=zlib-stream"

// Session is a single connection to the Discord Gateway.
type Session struct {
	cancel            func()
	ctx               context.Context
	heartbeatInterval time.Duration
	shardCount        uint
	shardID           uint
	token             string
	ws                *websocket.Conn
}

func NewSession(shardCount uint, shardID uint, token string) *Session {
	return &Session{
		shardCount: shardCount,
		shardID:    shardID,
		token:      token,
	}
}

func (s *Session) Open(ctx context.Context) error {
	// Wrap context to get a cancel func.
	s.ctx, s.cancel = context.WithCancel(ctx)

	// Attempt to handshake.
	ws, _, err := websocket.Dial(ctx, gatewayURL, websocket.DialOptions{})
	if err != nil {
		s.cancel()
		return xerrors.Errorf("failed to handshake: %v", err)
	}

	// Store websocket.
	s.ws = ws

	// Set message limit to 1GB.
	s.ws.SetReadLimit(1073741824)

	// Read hello payload.
	err = s.readHello()
	if err != nil {
		return xerrors.Errorf("failed to read hello payload: %v", err)
	}

	return nil
}

func (s *Session) readHello() error {
	// Read payload.
	payload, err := s.readPayload()
	if err != nil {
		return xerrors.Errorf("failed to get hello payload: %v", err)
	}

	// We should get an opcode 10.
	if payload.Op != 10 {
		return xerrors.Errorf("failed to get hello payload: %v", err)
	}

	// Unmarshal payload data.
	var helloPayload helloPayload
	if err = unmarshal(payload.D, &helloPayload); err != nil {
		return xerrors.Errorf("failed to unmarshal hello payload: %v", err)
	}

	// Store heatbeat interval.
	s.heartbeatInterval = time.Duration(helloPayload.HeartbeatInterval) * time.Millisecond

	return nil
}

func (s *Session) readPayload() (*payload, error) {
	// Read payload.
	msgType, msg, err := s.ws.Read(s.ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to read payload: %v", err)
	}

	// Get payload data.
	var data []byte
	if msgType == websocket.MessageBinary {
		// Decompress payload.
		data, err = decompress(msg)
		if err != nil {
			return nil, xerrors.Errorf("failed to decompress payload: %v", err)
		}
	} else {
		data = msg
	}

	// Unmarshal payload into json.
	var payload payload
	if err = unmarshal(data, &payload); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal payload: %v", err)
	}

	return &payload, nil
}
