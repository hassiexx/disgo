package gateway

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"nhooyr.io/websocket"
)

const gatewayURL = "wss://gateway.discord.gg/?v=6&encoding=json"

// Session is a single connection to the Discord Gateway.
type Session struct {
	disconnect     chan bool
	done           chan struct{}
	heartbeatState *heartbeatState
	log            *zap.Logger
	sequence       int64
	sessionID      string
	shardCount     uint
	shardID        uint
	token          string
	wg             sync.WaitGroup
	ws             *websocket.Conn
}

func NewSession(logger *zap.Logger, shardCount uint, shardID uint, token string) *Session {
	return &Session{
		log:        logger,
		shardCount: shardCount,
		shardID:    shardID,
		token:      token,
	}
}

func (s *Session) Open(ctx context.Context) error {
	// Create channels.
	s.disconnect = make(chan bool)
	s.done = make(chan struct{})

	// Initialise websocket headers.
	headers := http.Header{}
	headers.Add("Accept-Encoding", "zlib")

	// Attempt to handshake.
	s.log.Debug("attempting to connect to the gateway", zap.Uint("shard", s.shardID))
	ws, _, err := websocket.Dial(ctx, gatewayURL, websocket.DialOptions{HTTPHeader: headers})
	if err != nil {
		close(s.done)
		return xerrors.Errorf("failed to handshake: %w", err)
	}

	// Store websocket.
	s.ws = ws

	// Set message limit to 1GB.
	s.ws.SetReadLimit(1073741824)

	// Read hello payload.
	s.log.Debug("receiving hello payload", zap.Uint("shard", s.shardID))
	err = s.readHello(ctx)
	if err != nil {
		close(s.done)
		return xerrors.Errorf("failed to read hello payload: %w", err)
	}

	// Start heartbeating.
	s.wg.Add(1)
	go s.heartbeat(ctx)

	// If we don't have a sequence number and a session ID, create a new session,
	// otherwise resume.
	if s.sequence == 0 && s.sessionID == "" {
		err = s.identify(ctx)
		if err != nil {
			close(s.done)
			return xerrors.Errorf("failed to identify: %w", err)
		}

		// Start handling connection.
		go s.handleConnection(ctx)
	} else {
		// TODO: resume
	}

	return nil
}

func (s *Session) handleConnection(ctx context.Context) {
	// Var to store whether to reconnect after disconnecting.
	var reconnect bool

	select {
	case <-s.done:
		// Received disconnect signal from external source.
		// Stop handling connection.
		return
	case reconnect = <-s.disconnect:
	case <-ctx.Done():
	}

	// Signal goroutines to stop.
	close(s.done)

	// Wait for goroutines to stop.
	s.wg.Wait()

	// Close session.
	err := s.ws.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		s.log.Info("failed to close websocket", zap.Uint("shard", s.shardID))
		return
	}

	if reconnect {
		// Reopen session.
		err := s.Open(ctx)
		if err != nil {
			s.log.Info("failed to reconnect websocket", zap.Uint("shard", s.shardID))
		}
	} else {
		// Reset session state.
		s.sequence = 0
		s.sessionID = ""
	}
}

func (s *Session) heartbeat(ctx context.Context) {
	// Create ticker.
	ticker := time.NewTicker(s.heartbeatState.Interval)

	// Stop ticker and call done on wait group before returning.
	defer ticker.Stop()
	defer s.wg.Done()

	for {
		// Lock heartbeat state.
		s.heartbeatState.Lock()

		// If we have not received a heartbeat ack since the last heartbeat time,
		// we need to disconnect and attempt to resume.
		if s.heartbeatState.LastHeartbeatAck.Before(s.heartbeatState.LastHeartbeatSend) {
			// Log.
			s.log.Warn("did not receive a heartbeat ack, reconnecting", zap.Uint("shard", s.shardID))

			// Signal reconnect.
			s.disconnect <- true
		} else {
			// Send heartbeat.
			s.log.Debug("sending heartbeat", zap.Uint("shard", s.shardID))
			s.heartbeatState.LastHeartbeatSend = time.Now().UTC()
			err := s.sendHeartbeat(ctx)
			if err != nil {
				// Signal reconnect.
				s.disconnect <- true
			}
		}

		// Unlock heartbeat state.
		s.heartbeatState.Unlock()

		select {
		case <-ticker.C:
			// Send heartbeat.
		case <-s.done:
			// Stop heartbeating.
			return
		case <-ctx.Done():
			// Stop heartbeating.
			return
		}
	}
}

func (s *Session) readHello(ctx context.Context) error {
	// Read payload.
	payload, err := s.readPayload(ctx)
	if err != nil {
		return xerrors.Errorf("failed to get hello payload: %w", err)
	}

	// We should get an opcode 10.
	if payload.Op != 10 {
		return xerrors.Errorf("failed to get hello payload: got opcode %d instead", payload.Op)
	}

	// Unmarshal payload data.
	var helloData helloData
	if err = unmarshal(payload.D, &helloData); err != nil {
		return xerrors.Errorf("failed to unmarshal hello payload: %w", err)
	}

	// Store heatbeat interval.
	s.heartbeatState = &heartbeatState{
		Interval: time.Duration(helloData.HeartbeatInterval) * time.Millisecond,
	}

	return nil
}

func (s *Session) sendHeartbeat(ctx context.Context) error {
	// Create payload.
	payload := heartbeatPayload{
		Op: uint(OpcodeHeartbeat),
		D:  0,
	}

	// Marshal payload.
	data, err := marshal(payload)
	if err != nil {
		return xerrors.Errorf("failed to marshal heartbeat payload: %w", err)
	}

	// Send payload.
	if err = s.ws.Write(ctx, websocket.MessageText, data); err != nil {
		return xerrors.Errorf("failed to send heartbeat payload: %w", err)
	}

	return nil
}

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

	err := s.sendPayload(ctx, payload)
	if err != nil {
		return xerrors.Errorf("failed to send identify payload: %w", err)
	}

	return nil
}

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
