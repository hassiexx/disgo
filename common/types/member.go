package types

// Member is a struct for a member.
type Member struct {
	Deaf         bool     `json:"deaf"`
	JoinedAt     string   `json:"joined_at"`
	Mute         bool     `json:"mute"`
	Nick         string   `json:"nick"`
	PremiumSince string   `json:"premium_since"`
	Roles        []string `json:"roles"`
	User         *User    `json:"user"`
}
