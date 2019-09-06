package types

import (
	"encoding/json"
)

// GuildChannel is a type for guild channels.
type GuildChannel map[string]*Channel

// GuildEmoji is a type for guild emojis.
type GuildEmoji map[string]*Emoji

// GuildMember is a type for guild members.
type GuildMember map[string]*Member

// GuildRole is a type for guild roles.
type GuildRole map[string]*Role

// GuildPresence is a type for guild presences.
type GuildPresence map[string]*Presence

// Guild is a struct for a guild.
type Guild struct {
	AFKChannelID                string        `json:"afk_channel_id"`
	AFKTimeout                  uint          `json:"afk_timeout"`
	ApplicationID               string        `json:"application_id"`
	Banner                      string        `json:"banner"`
	Channels                    GuildChannel  `json:"channels"`
	DefaultMessageNotifications uint          `json:"default_message_notifications"`
	Description                 string        `json:"description"`
	EmbedChannelID              string        `json:"embed_channel_id"`
	EmbedEnabled                bool          `json:"embed_enabled"`
	Emojis                      GuildRole     `json:"emoji"`
	ExplicitContentFilter       uint          `json:"explicit_content_filter"`
	Features                    []string      `json:"features"`
	Icon                        string        `json:"icon"`
	ID                          string        `json:"id"`
	JoinedAt                    string        `json:"joined_at"`
	Large                       bool          `json:"large"`
	MaxMembers                  uint          `json:"max_members"`
	MaxPresences                uint          `json:"max_presences"`
	MemberCount                 uint          `json:"member_count"`
	Members                     GuildMember   `json:"members"`
	MFALevel                    uint          `json:"mfa_level"`
	Name                        string        `json:"name"`
	Owner                       bool          `json:"owner"`
	OwnerID                     string        `json:"owner_id"`
	Permissions                 uint          `json:"permissions"`
	PreferredLocale             string        `json:"preferred_locale"`
	PremiumSubscriptionType     uint          `json:"premium_subscription_type"`
	PremiumTier                 uint          `json:"premium_tier"`
	Presences                   GuildPresence `json:"presences"`
	Region                      string        `json:"region"`
	Roles                       GuildRole     `json:"roles"`
	Splash                      string        `json:"splash"`
	SystemChannelID             string        `json:"system_channel_id"`
	Unavailable                 bool          `json:"available"`
	VanityURLCode               string        `json:"vanity_url_code"`
	VerificationLevel           uint          `json:"verification_level"`
	VoiceStates                 []*VoiceState `json:"voice_states"`
	WidgetChannelID             string        `json:"widget_channel_id"`
	WidgetEnabled               bool          `json:"widget_enabled"`
}

// UnmarshalJSON unmarshals guild channels into a map.
func (g *GuildChannel) UnmarshalJSON(data []byte) error {
	// Initialise map.
	*g = make(map[string]*Channel)

	// Dereference map.
	m := *g

	// Unmarshal channels.
	var channels []*Channel
	if err := json.Unmarshal(data, &channels); err != nil {
		return err
	}

	// Add channels to map, using the channel ID as the key.
	for _, channel := range channels {
		m[channel.ID] = channel
	}

	return nil
}

// UnmarshalJSON unmarshals guild emojis into a map.
func (g *GuildEmoji) UnmarshalJSON(data []byte) error {
	// Initialise map.
	*g = make(map[string]*Emoji)

	// Dereference map.
	m := *g

	// Unmarshal emojis.
	var emojis []*Emoji
	if err := json.Unmarshal(data, &emojis); err != nil {
		return err
	}

	// Add emojis to map, using the emoji ID as the key.
	for _, emoji := range emojis {
		m[emoji.ID] = emoji
	}

	return nil
}

// UnmarshalJSON unmarshals guild members into a map.
func (g *GuildMember) UnmarshalJSON(data []byte) error {
	// Initialise map.
	*g = make(map[string]*Member)

	// Dereference map.
	m := *g

	// Unmarshal members.
	var members []*Member
	if err := json.Unmarshal(data, &members); err != nil {
		return err
	}

	// Add members to map, using the user ID as the key.
	for _, member := range members {
		m[member.User.ID] = member
	}

	return nil
}

// UnmarshalJSON unmarshals guild presences into a map.
func (g *GuildPresence) UnmarshalJSON(data []byte) error {
	// Initialise map.
	*g = make(map[string]*Presence)

	// Dereference map.
	m := *g

	// Unmarshal presences.
	var presences []*Presence
	if err := json.Unmarshal(data, &presences); err != nil {
		return err
	}

	// Add presences to map, using the user ID as the key.
	for _, presence := range presences {
		m[presence.User.ID] = presence
	}

	return nil
}

// UnmarshalJSON unmarshals guild roles into a map.
func (g *GuildRole) UnmarshalJSON(data []byte) error {
	// Initialise map.
	*g = make(map[string]*Role)

	// Dereference map.
	m := *g

	// Unmarshal roles.
	var roles []*Role
	if err := json.Unmarshal(data, &roles); err != nil {
		return err
	}

	// Add roles to map, using the role ID as the key.
	for _, role := range roles {
		m[role.ID] = role
	}

	return nil
}
