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
	GuildID              uint64                 `json:"guild_id,string"`
	Icon                 string                 `json:"icon"`
	ID                   uint64                 `json:"id,string"`
	LastMessageID        uint64                 `json:"last_message_id,string"`
	LastPinTimestamp     string                 `json:"last_pin_timestamp"`
	Name                 string                 `json:"name"`
	NSFW                 bool                   `json:"nsfw"`
	ParentID             uint64                 `json:"parent_id,string"`
	PermissionOverwrites []*PermissionOverwrite `json:"permission_overwrites"`
	Position             uint                   `json:"position"`
	RateLimitPerUser     uint                   `json:"rate_limit_per_user"`
	Recipients           []*User                `json:"recipients"`
	Topic                string                 `json:"topic"`
	Type                 uint                   `json:"type"`
	UserLimit            uint                   `json:"user_limit"`

	// For state caching.
	RecipientSet *UInt64HashSet
}
