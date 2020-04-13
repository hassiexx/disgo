package gateway

import (
	"context"
	"sync"
	"time"

	"github.com/hassieswift621/disgo/event"
	"github.com/hassieswift621/disgo/statecore"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"nhooyr.io/websocket"
)

// Gateway URL
const gatewayURL = "wss://gateway.discord.gg/?v=6&encoding=json"

// Session is a single connection to the Discord Gateway.
type Session struct {
	dispatcher     event.Dispatcher
	disconnect     chan bool
	done           chan struct{}
	heartbeatState *heartbeatState
	log            *zap.Logger
	sequence       uint
	sessionID      string
	shardCount     uint
	shardID        uint
	state          statecore.State
	token          string
	wg             sync.WaitGroup
	ws             *websocket.Conn
}

// NewSession creates a new session.
func NewSession(dispatcher event.Dispatcher, logger *zap.Logger, shardCount uint, shardID uint, token string) *Session {
	return &Session{
		dispatcher: dispatcher,
		log:        logger,
		shardCount: shardCount,
		shardID:    shardID,
		token:      token,
	}
}

// Open opens the session.
func (s *Session) Open(ctx context.Context) error {
	// Create channels.
	s.disconnect = make(chan bool)
	s.done = make(chan struct{})

	// Attempt to handshake.
	s.log.Debug("connecting to the gateway", zap.Uint("shard", s.shardID))
	ws, _, err := websocket.Dial(ctx, gatewayURL, &websocket.DialOptions{})
	if err != nil {
		close(s.done)
		return xerrors.Errorf("connect to gateway: %w", err)
	}

	// Store websocket.
	s.ws = ws

	// Set message limit to 1GB.
	s.ws.SetReadLimit(1073741824)

	// Read hello payload.
	s.log.Debug("receiving hello payload", zap.Uint("shard", s.shardID))
	if err = s.hello(ctx); err != nil {
		close(s.done)
		return xerrors.Errorf("read hello payload: %w", err)
	}

	// Start heartbeating.
	s.wg.Add(1)
	go s.heartbeat(ctx)

	// If we have a sequence number or a session ID, attempt to resume,
	// otherwise create a new session.
	if s.sequence != 0 || s.sessionID != "" {
		s.log.Debug("resuming session")
		identifyLimiter.acquire()
		if err = s.resume(ctx); err != nil {
			s.log.Debug("failed to resume session", zap.Uint("shard", s.shardID), zap.Error(err))
			s.sequence = 0
			s.sessionID = ""
		}
	}

	// If sequence is 0 attempt to identify.
	if s.sequence == 0 {
		s.log.Debug("identifying", zap.Uint("shard", s.shardID))
		identifyLimiter.acquire()
		if err = s.identify(ctx); err != nil {
			close(s.done)
			return xerrors.Errorf("identify: %w", err)
		}
	}

	// Connected, start handling connection and events.
	go s.handleConnection(ctx)

	return nil
}

// HandleConnection handles the connection by listening to the signals and handling events.
func (s *Session) handleConnection(ctx context.Context) {
	// Var to store whether to reconnect after disconnecting.
	var reconnect bool

	select {
	case <-s.done:
		// Received disconnect signal from external source.
		// Stop handling connection.
	case reconnect = <-s.disconnect:
	case <-ctx.Done():
	default:
		s.handleEvent(ctx)
	}

	// Signal goroutines to stop.
	close(s.done)

	// Wait for goroutines to stop.
	s.wg.Wait()

	// Close session.
	if err := s.ws.Close(websocket.StatusInternalError, ""); err != nil {
		s.log.Debug("failed to close websocket", zap.Uint("shard", s.shardID), zap.Error(err))
		return
	}

	if reconnect {
		// Reopen session.
		if err := s.Open(ctx); err != nil {
			s.log.Debug("failed to reconnect", zap.Uint("shard", s.shardID), zap.Error(err))
		}
	} else {
		// Reset session state.
		s.sequence = 0
		s.sessionID = ""
	}
}

func (s *Session) handleEvent(ctx context.Context) {
	payload, err := s.readPayload(ctx)
	if err != nil {
		s.log.Debug("failed to read payload", zap.Uint("shard", s.shardID), zap.Error(err))
		s.disconnect <- true
	} else {
		switch payload.Op {
		// Dispatch.
		case opcodeDispatch:
			// Store sequence number.
			s.sequence = payload.S

			// Var to store error from handling event.
			var err error

			// Check event type.
			switch payload.T {
			// Channel Create.
			case eventChannelCreate:
				s.log.Debug("handling channel create event")
				err = s.channelCreate(payload.D)

			// Channel Update.
			case eventChannelUpdate:
				s.log.Debug("handling channel update event")
				err = s.channelUpdate(payload.D)

			// Ready.
			case eventReady:
				s.log.Debug("handling ready event")
				err = s.ready(payload.D)
			}

			if err != nil {
				s.log.Debug("failed to handle event", zap.Uint("shard", s.shardID), zap.Error(err))
			}

		// Heartbeat.
		case opcodeHeartbeat:
			err := s.sendHeartbeat(ctx)
			if err != nil {
				s.log.Debug("failed to send heartbeat upon request", zap.Uint("shard", s.shardID),
					zap.Error(err))
				s.disconnect <- true
			}

		// Reconnect.
		case opcodeReconnect:
			s.disconnect <- true

		// Invalid session.
		case opcodeInvalidSession:
			s.sequence = 0
			s.sessionID = ""
			s.disconnect <- true

		// Heartbeat ACK.
		case opcodeHeartbeatACK:
			s.heartbeatState.Lock()
			s.heartbeatState.LastHeartbeatAck = time.Now()
			s.heartbeatState.Unlock()
		}
	}
}
