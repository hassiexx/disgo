package disgo

import snowflake "github.com/hassieswift621/discord-goflake"

type Webhook struct {
	avatarHash string
	channelID  *snowflake.Snowflake
	guildID    *snowflake.Snowflake
	id         *snowflake.Snowflake
	name       string
	token      string
	user       *User
}

// AvatarHash gets the webhook's default avatar hash.
func (w *Webhook) AvatarHash() string {
	return w.avatarHash
}

// ChannelID gets the snowflake ID of the channel which this webhook is for.
func (w *Webhook) ChannelID() *snowflake.Snowflake {
	return w.channelID
}

// GuildID gets the snowflake ID of the guild which this webhook is for.
func (w *Webhook) GuildID() *snowflake.Snowflake {
	return w.guildID
}

// ID gets the snowflake ID of the webhook.
func (w *Webhook) ID() *snowflake.Snowflake {
	return w.id
}

// Name gets the default name of the webhook.
func (w *Webhook) Name() string {
	return w.name
}

// Token gets the secure token of the webhook.
func (w *Webhook) Token() string {
	return w.token
}

// User gets the user which created this webhook.
// This is nil when the webhook is fetched using its token.
func (w *Webhook) User() *User {
	return w.user
}
