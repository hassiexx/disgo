package types

// VoiceState is a struct for a voice state.
type VoiceState struct {
	ChannelID uint64  `json:"channel_id,string"`
	Deaf      bool    `json:"deaf"`
	GuildID   uint64  `json:"guild_id,string"`
	Member    *Member `json:"member"`
	Mute      bool    `json:"mute"`
	SelfDeaf  bool    `json:"self_deaf"`
	SelfMute  bool    `json:"self_mute"`
	SessionID string  `json:"session_id"`
	Suppress  bool    `json:"suppress"`
	UserID    uint64  `json:"user_id,string"`
}
