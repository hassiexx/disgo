package statecore

import (
	"github.com/hassieswift621/disgo/common/types"
)

// State is the interface for a state caching implementation.
type State interface {

	// AddGuildsReady adds unavailable guilds from the ready event.
	AddGuildsReady(guilds []*types.Guild)

	// Channel gets a channel.
	Channel(id string) (*types.Channel, error)

	// Emoji gets an emoji.
	Emoji(id string) (*types.Emoji, error)

	// Guild gets a guild.
	Guild(id string) (*types.Guild, error)

	// Member gets a guild member.
	Member(guildID string, memberID string) (*types.Member, error)

	// Message gets a message.
	Message(id string) (*types.Message, error)

	// PermissionOverwrite gets a role or user permission overwrite for a channel.
	PermissionOverwrite(channelID, overwriteID string) (*types.PermissionOverwrite, error)

	// Presence gets a user's guild presence.
	Presence(guildID string, userID string) (*types.Presence, error)

	// Role gets a role.
	Role(id string) (*types.Role, error)

	// Self gets the bot user.
	Self() (*types.User, error)

	// SetSelf sets the bot user.
	SetSelf(self *types.User)

	// User gets a user.
	User(id string) (*types.User, error)
}
