package gateway

type GatewayOpcode uint

const (
	GatewayOpDispatch            GatewayOpcode = 0
	GatewayOpHeartbeat           GatewayOpcode = 1
	GatewayOpIdentify            GatewayOpcode = 2
	GatewayOpStatusUpdate        GatewayOpcode = 3
	GatewayOpVoiceStateUpdate    GatewayOpcode = 4
	GatewayOpResume              GatewayOpcode = 6
	GatewayOpReconnect           GatewayOpcode = 7
	GatewayOpRequestGuildMembers GatewayOpcode = 8
	GatewayOpInvalidSession      GatewayOpcode = 9
	GatewayOpHello               GatewayOpcode = 10
	GatewayOpHeartbeatACK        GatewayOpcode = 11
)
