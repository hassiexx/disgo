package disgo

import snowflake "github.com/hassieswift621/discord-goflake"

// Member is an immutable struct for a guild member.
type Member struct {
	deaf         bool
	joinedAt     Timestamp
	mute         bool
	nick         string
	premiumSince Timestamp
	roles        []*snowflake.Snowflake
}

// IsDeafened returns whether the member is deafened in voice channels.
func (m *Member) IsDeafened() bool {
	return m.deaf
}

// IsMuted returns whether the member is muted in voice channels.
func (m *Member) IsMuted() bool {
	return m.mute
}

// JoinedAt gets the timestamp when the member joined the guild.
func (m *Member) JoinedAt() Timestamp {
	return m.joinedAt
}

// Nick gets the member's nick.
func (m *Member) Nick() string {
	return m.nick
}

// PremiumSince gets the timestamp when the member used Nitro boost on the guild.
func (m *Member) PremiumSince() Timestamp {
	return m.premiumSince
}

// Roles gets the member's roles.
func (m *Member) Roles() []*Role {
	// TODO
	return nil
}

// RoleIDs gets the member's roles as snowflake IDs.
func (m *Member) RoleIDs() []*snowflake.Snowflake {
	// Copy slice.
	roleIDs := make([]*snowflake.Snowflake, len(m.roles))
	copy(roleIDs, m.roles)

	return roleIDs
}
