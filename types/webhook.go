package types

// Webhook is a struct for a webhook.
type Webhook struct {
	Avatar    string `json:"avatar"`
	ChannelID uint64 `json:"channel_id,string"`
	GuildID   uint64 `json:"guild_id,string"`
	ID        uint64 `json:"id,string"`
	Name      string `json:"name"`
	Token     string `json:"token"`
	User      *User  `json:"user"`
}
