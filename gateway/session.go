package gateway

import (
	"context"
	"log"

	"nhooyr.io/websocket"
)

const gatewayURL = "wss://gateway.discord.gg/?v=6&encoding=json&compress=zlib-stream"

// Session is a single connection to the Discord Gateway.
type Session struct {
	cancel            func()
	ctx               context.Context
	heartbeatInterval int
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
		return err
	}

	// Store websocket.
	s.ws = ws

	// Set message limit to 1GB.
	s.ws.SetReadLimit(1073741824)

	// Handle hello payload.
	err = s.handleHello()
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) handleHello() error {
	// Read hello payload.
	msgType, msg, err := s.ws.Read(s.ctx)
	if err != nil {
		return err
	}
	log.Println(msg)

	// Parse payload.
	var data []byte

	// We should get a compressed zlib stream.
	if msgType == websocket.MessageBinary {
		// Decompress.
		data, err = decompress(msg)
		if err != nil {
			return err
		}
	} else {
		data = msg
	}

	return err
}
