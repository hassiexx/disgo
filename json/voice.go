package json

// VoiceState is a struct for a voice state.
type VoiceState struct {
	ChannelID string  `json:"channel_id"`
	Deaf      bool    `json:"deaf"`
	GuildID   string  `json:"guild_id"`
	Member    *Member `json:"member"`
	Mute      bool    `json:"mute"`
	SelfDeaf  bool    `json:"self_deaf"`
	SelfMute  bool    `json:"self_mute"`
	SessionID string  `json:"session_id"`
	Suppress  bool    `json:"suppress"`
	UserID    string  `json:"user_id"`
}
