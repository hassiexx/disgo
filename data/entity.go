package data

import (
	dgoflake "github.com/hassieswift621/discord-goflake"
)

// Entity is the interface for a Discord entity.
type Entity interface {
	ID() *dgoflake.Snowflake
}
