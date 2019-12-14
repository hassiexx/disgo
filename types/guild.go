package types

// Guild is a struct for a guild.
type Guild struct {
	AFKChannelID                uint64        `json:"afk_channel_id,string"`
	AFKTimeout                  uint          `json:"afk_timeout"`
	ApplicationID               uint64        `json:"application_id,string"`
	Banner                      string        `json:"banner"`
	Channels                    []*Channel    `json:"channels"`
	DefaultMessageNotifications uint          `json:"default_message_notifications"`
	Description                 string        `json:"description"`
	EmbedChannelID              uint64        `json:"embed_channel_id,string"`
	EmbedEnabled                bool          `json:"embed_enabled"`
	Emojis                      []*Emoji      `json:"emoji"`
	ExplicitContentFilter       uint          `json:"explicit_content_filter"`
	Features                    []string      `json:"features"`
	Icon                        string        `json:"icon"`
	ID                          uint64        `json:"id,string"`
	JoinedAt                    string        `json:"joined_at"`
	Large                       bool          `json:"large"`
	MaxMembers                  uint          `json:"max_members"`
	MaxPresences                uint          `json:"max_presences"`
	MemberCount                 uint          `json:"member_count"`
	Members                     []*Member     `json:"members"`
	MFALevel                    uint          `json:"mfa_level"`
	Name                        string        `json:"name"`
	Owner                       bool          `json:"owner"`
	OwnerID                     uint64        `json:"owner_id,string"`
	Permissions                 uint          `json:"permissions"`
	PreferredLocale             string        `json:"preferred_locale"`
	PremiumSubscriptionType     uint          `json:"premium_subscription_type"`
	PremiumTier                 uint          `json:"premium_tier"`
	Presences                   []*Presence   `json:"presences"`
	Region                      string        `json:"region"`
	Roles                       []*Role       `json:"roles"`
	Splash                      string        `json:"splash"`
	SystemChannelID             uint64        `json:"system_channel_id,string"`
	Unavailable                 bool          `json:"available"`
	VanityURLCode               string        `json:"vanity_url_code"`
	VerificationLevel           uint          `json:"verification_level"`
	VoiceStates                 []*VoiceState `json:"voice_states"`
	WidgetChannelID             uint64        `json:"widget_channel_id,string"`
	WidgetEnabled               bool          `json:"widget_enabled"`

	// For state caching.
	ChannelSet *UInt64HashSet
	EmojiSet   *UInt64HashSet
	RoleSet    *UInt64HashSet
}
