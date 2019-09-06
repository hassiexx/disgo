package statecore

import (
	"github.com/hassieswift621/disgo/common/types"
)

// State is the interface for a state caching implementation.
type State interface {
	// Channel gets a channel by its ID.
	Channel(id string) (*types.Channel, error)
	// Emoji gets an emoji by its ID.
	Emoji(id string) (*types.Emoji, error)
	// Guild gets a guild by its ID.
	Guild(id string) (*types.Guild, error)
	// Member gets a guild member by their user ID.
	Member(id string) (*types.Member, error)
	// Presence gets a user presence by a user ID.
	Presence(id string) (*types.Presence, error)
	// Role gets a role by its ID.
	Role(id string) (*types.Role, error)
	// User gets a user by their ID.
	User(id string) (*types.User, error)
}
