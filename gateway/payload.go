package gateway

import "encoding/json"

type heartbeatPayload struct {
	Op uint  `json:"op"`
	D  int64 `json:"d"`
}

type helloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type identifyPayload struct {
	Op uint         `json:"op"`
	D  identifyData `json:"d"`
}

type identifyData struct {
	Token              string             `json:"token"`
	Properties         identifyProperties `json:"properties"`
	Compress           bool               `json:"compress"`
	LargeThreshold     uint               `json:"large_threshold"`
	Shard              []uint             `json:"shard"`
	GuildSubscriptions bool               `json:"guild_subscriptions"`
}

type identifyProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type payload struct {
	Op uint            `json:"op"`
	D  json.RawMessage `json:"d"`
	S  uint            `json:"s"`
	T  string          `json:"t"`
}
