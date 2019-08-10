package gateway

import (
	"context"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"nhooyr.io/websocket"
)

type heartbeatPayload struct {
	Op uint   `json:"op"`
	D  uint64 `json:"d"`
}

// Heartbeat sends heartbeats at the specified interval.
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

// SendHeartbeat sends a heartbeat payload.
func (s *Session) sendHeartbeat(ctx context.Context) error {
	// Create payload.
	payload := heartbeatPayload{
		Op: uint(OpcodeHeartbeat),
		D:  atomic.LoadUint64(&s.sequence),
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
