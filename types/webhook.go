package types

// Webhook is a struct for a webhook.
type Webhook struct {
	Avatar    string `json:"avatar"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Token     string `json:"token"`
	User      *User  `json:"user"`
}
