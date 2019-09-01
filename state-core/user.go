package statecore

// User is a struct for raw user state data.
type User struct {
	AvatarHash    string `json:"avatar"`
	Bot           bool   `json:"bot"`
	Discriminator string `json:"discriminator"`
	Flags         uint   `json:"flags"`
	ID            string `json:"id"`
	Locale        string `json:"locale"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	PremiumType   uint   `json:"premium_type"`
	Username      string `json:"username"`
}
