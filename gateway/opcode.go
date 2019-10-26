package gateway

type opcode uint

const (
	opcodeDispatch            opcode = 0
	opcodeHeartbeat           opcode = 1
	opcodeIdentify            opcode = 2
	opcodeStatusUpdate        opcode = 3
	opcodeVoiceStateUpdate    opcode = 4
	opcodeResume              opcode = 6
	opcodeReconnect           opcode = 7
	opcodeRequestGuildMembers opcode = 8
	opcodeInvalidSession      opcode = 9
	opcodeHello               opcode = 10
	opcodeHeartbeatACK        opcode = 11
)
