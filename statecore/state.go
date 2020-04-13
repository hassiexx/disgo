package statecore

import (
	"github.com/hassieswift621/disgo/types"
)

// State is the interface for a state caching implementation.
type State interface {
	// AddChannel adds a channel.
	AddChannel(channel *types.Channel)

	// AddGuildsReady adds unavailable guilds from the ready event.
	AddGuildsReady(guilds []*types.Guild)

	// Channel gets a channel.
	Channel(id uint64) (*types.Channel, error)

	// Emoji gets an emoji.
	Emoji(id uint64) (*types.Emoji, error)

	// Guild gets a guild.
	Guild(id uint64) (*types.Guild, error)

	// Member gets a guild member.
	Member(guildID uint64, memberID uint64) (*types.Member, error)

	// Message gets a message.
	Message(id uint64) (*types.Message, error)

	// PermissionOverwrite gets a role or user permission overwrite for a channel.
	PermissionOverwrite(channelID, overwriteID uint64) (*types.PermissionOverwrite, error)

	// Presence gets a user's guild presence.
	Presence(guildID uint64, userID uint64) (*types.Presence, error)

	// Role gets a role.
	Role(id uint64) (*types.Role, error)

	// Self gets the bot user.
	Self() (*types.User, error)

	// SetSelf sets the bot user.
	SetSelf(self *types.User)

	// User gets a user.
	User(id uint64) (*types.User, error)
}
