package statecore

// User is a struct for a user state.
type User struct {
	AvatarHash    string
	Bot           bool
	Discriminator string
	Flags         uint
	ID            string
	Locale        string
	MFAEnabled    bool
	PremiumType   uint
	Username      string
}
