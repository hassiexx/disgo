package event

import "github.com/hassieswift621/disgo/types"

// Handler is the interface for an event handler.
type Handler interface {
	// ChannelCreate event.
	ChannelCreate(channel types.Channel)

	// ChannelUpdate event.
	ChannelUpdate(channel types.Channel)

	// Ready event.
	Ready(shardID uint)
}
