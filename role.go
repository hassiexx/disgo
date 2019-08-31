package disgo

import snowflake "github.com/hassieswift621/discord-goflake"

// Role is an immutable struct for a role.
type Role struct {
	colour      int
	hoist       bool
	id          *snowflake.Snowflake
	managed     bool
	mentionable bool
	name        string
	permissions PermissionBitSet
	position    uint
}

// Colour gets the colour of the role in its integer representation.
func (r *Role) Colour() int {
	return r.colour
}

// ID gets the snowflake ID of the role.
func (r *Role) ID() *snowflake.Snowflake {
	return r.id
}

// IsHoist returns whether the role is shown in the user list.
func (r *Role) IsHoist() bool {
	return r.hoist
}

// IsManaged returns whether the role is part of an integration.
func (r *Role) IsManaged() bool {
	return r.managed
}

// IsMentionable returns whether the role can be mentioned.
func (r *Role) IsMentionable() bool {
	return r.mentionable
}

// Name gets the name of the role.
func (r *Role) Name() string {
	return r.name
}

// Permissions gets the permissions bit set for the role.
func (r *Role) Permissions() PermissionBitSet {
	return r.permissions
}

// Position gets the position of the role.
func (r *Role) Position() uint {
	return r.position
}
