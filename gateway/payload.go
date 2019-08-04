package gateway

import "encoding/json"

type payload struct {
	D  json.RawMessage `json:"d"`
	Op uint            `json:"op"`
	S  uint            `json:"s"`
	T  string          `json:"t"`
}

type helloPayload struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}
