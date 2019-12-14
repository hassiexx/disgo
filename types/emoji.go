package types

// Emoji is a struct for an emoji.
type Emoji struct {
	Animated      bool     `json:"animated"`
	ID            uint64   `json:"id,string"`
	Managed       bool     `json:"managed"`
	Name          string   `json:"name"`
	RequireColons bool     `json:"require_colons"`
	Roles         []string `json:"roles"`
	User          *User    `json:"user"`
}
