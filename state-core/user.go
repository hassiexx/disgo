package statecore

// UserState is the struct used to store raw user state data.
type UserState struct {
	// The user's avatar hash.
	AvatarHash string `json:"avatar"`
	// Whether the user is a bot.
	Bot bool `json:"bot"`
	// The user's 4 digit discriminator.
	Discriminator string `json:"discriminator"`
	// The user's profile flags, also called badges.
	Flags uint `json:"flags"`
	// The user's snowflake ID.
	ID string `json:"id"`
	// The user's locale.
	Locale string `json:"locale"`
	// Whether the user has MFA enabled.
	MFAEnabled bool `json:"mfa_enabled"`
	// The user's Nitro subscription type.
	PremiumType uint `json:"premium_type"`
	// The user's username.
	Username string
}
