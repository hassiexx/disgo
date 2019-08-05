package gateway

import (
	"sync"
	"time"
)

type heartbeatState struct {
	sync.Mutex
	Interval          time.Duration
	LastHeartbeatAck  time.Time
	LastHeartbeatSend time.Time
}
