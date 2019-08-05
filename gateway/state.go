package gateway

import (
	"sync"
	"time"
)

type heartbeatState struct {
	sync.Mutex
	Interval          time.Duration
	LastHeartbeatAck  time.Time
	LastHeartbeatTime time.Time
}
