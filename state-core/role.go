package statecore

// Role is a struct for a role state.
type Role struct {
	Colour      uint
	Hoist       bool
	ID          string
	Managed     bool
	Mentionable bool
	Name        string
	Permissions uint
	Position    uint
}
