package event

import "github.com/hassieswift621/disgo/types"

// Dispatcher is the interface for an event dispatcher.
type Dispatcher interface {
	// ChannelCreate event.
	ChannelCreate(channel types.Channel)

	// ChannelUpdate event.
	ChannelUpdate(channel types.Channel)

	// Ready event.
	Ready(shardID uint)
}
