package statecore

import (
	"github.com/hassieswift621/disgo/json"
)

// State is the interface for a state caching implementation.
type State interface {
	// Channel gets a channel by its ID.
	Channel(id string) (*json.Channel, error)
	// Emoji gets an emoji by its ID.
	Emoji(id string) (*json.Emoji, error)
	// Guild gets a guild by its ID.
	Guild(id string) (*json.Guild, error)
	// Member gets a guild member by their user ID.
	Member(id string) (*json.Member, error)
	// Presence gets a user presence by a user ID.
	Presence(id string) (*json.Presence, error)
	// Role gets a role by its ID.
	Role(id string) (*json.Role, error)
	// User gets a user by their ID.
	User(id string) (*json.User, error)
}
