package types

// Guild is a struct for a guild.
type Guild struct {
	AFKChannelID                string        `json:"afk_channel_id"`
	AFKTimeout                  uint          `json:"afk_timeout"`
	ApplicationID               string        `json:"application_id"`
	Banner                      string        `json:"banner"`
	Channels                    []*Channel    `json:"channels"`
	DefaultMessageNotifications uint          `json:"default_message_notifications"`
	Description                 string        `json:"description"`
	EmbedChannelID              string        `json:"embed_channel_id"`
	EmbedEnabled                bool          `json:"embed_enabled"`
	Emojis                      []*Emoji      `json:"emoji"`
	ExplicitContentFilter       uint          `json:"explicit_content_filter"`
	Features                    []string      `json:"features"`
	Icon                        string        `json:"icon"`
	ID                          string        `json:"id"`
	JoinedAt                    string        `json:"joined_at"`
	Large                       bool          `json:"large"`
	MaxMembers                  uint          `json:"max_members"`
	MaxPresences                uint          `json:"max_presences"`
	MemberCount                 uint          `json:"member_count"`
	Members                     []*Member     `json:"members"`
	MFALevel                    uint          `json:"mfa_level"`
	Name                        string        `json:"name"`
	Owner                       bool          `json:"owner"`
	OwnerID                     string        `json:"owner_id"`
	Permissions                 uint          `json:"permissions"`
	PreferredLocale             string        `json:"preferred_locale"`
	PremiumSubscriptionType     uint          `json:"premium_subscription_type"`
	PremiumTier                 uint          `json:"premium_tier"`
	Presences                   []*Presence   `json:"presences"`
	Region                      string        `json:"region"`
	Roles                       []*Role       `json:"roles"`
	Splash                      string        `json:"splash"`
	SystemChannelID             string        `json:"system_channel_id"`
	Unavailable                 bool          `json:"available"`
	VanityURLCode               string        `json:"vanity_url_code"`
	VerificationLevel           uint          `json:"verification_level"`
	VoiceStates                 []*VoiceState `json:"voice_states"`
	WidgetChannelID             string        `json:"widget_channel_id"`
	WidgetEnabled               bool          `json:"widget_enabled"`
}
