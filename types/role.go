package types

// Role is a struct for a role.
type Role struct {
	Colour      uint   `json:"color"`
	Hoist       bool   `json:"hoist"`
	ID          uint64 `json:"id,string"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentionable"`
	Name        string `json:"name"`
	Permissions uint   `json:"permissions"`
	Position    uint   `json:"position"`
}
