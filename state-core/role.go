package statecore

// Role is a struct for raw role state data.
type Role struct {
	Colour      uint   `json:"color"`
	Hoist       bool   `json:"hoist"`
	ID          string `json:"id"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentionable"`
	Name        string `json:"name"`
	Permissions uint   `json:"permissions"`
	Position    uint   `json:"position"`
}
