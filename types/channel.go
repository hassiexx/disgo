package types

const (
	ChannelTypeGuildText uint = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
)

// Channel is a struct for a channel.
type Channel struct {
	ApplicationID        string                 `json:"application_id"`
	Bitrate              uint                   `json:"bitrate"`
	GuildID              string                 `json:"guild_id"`
	Icon                 string                 `json:"icon"`
	ID                   string                 `json:"id"`
	LastMessageID        string                 `json:"last_message_id"`
	LastPinTimestamp     string                 `json:"last_pin_timestamp"`
	Name                 string                 `json:"name"`
	NSFW                 bool                   `json:"nsfw"`
	ParentID             string                 `json:"parent_id"`
	PermissionOverwrites []*PermissionOverwrite `json:"permission_overwrites"`
	Position             uint                   `json:"position"`
	RateLimitPerUser     uint                   `json:"rate_limit_per_user"`
	Recipients           []*User                `json:"recipients"`
	Topic                string                 `json:"topic"`
	Type                 uint                   `json:"type"`
	UserLimit            uint                   `json:"user_limit"`

	// For state caching.
	RecipientSet *StringHashSet
}
