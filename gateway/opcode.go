package gateway

const (
	opcodeDispatch            uint = 0
	opcodeHeartbeat           uint = 1
	opcodeIdentify            uint = 2
	opcodeStatusUpdate        uint = 3
	opcodeVoiceStateUpdate    uint = 4
	opcodeResume              uint = 6
	opcodeReconnect           uint = 7
	opcodeRequestGuildMembers uint = 8
	opcodeInvalidSession      uint = 9
	opcodeHello               uint = 10
	opcodeHeartbeatACK        uint = 11
)
