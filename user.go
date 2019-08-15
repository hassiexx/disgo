package data

import snowflake "github.com/hassieswift621/discord-goflake"

// User is an immutable struct for a user.
type User struct {
	avatarHash    string
	bot           bool
	discriminator string
	flags         Flags
	id            *snowflake.Snowflake
	locale        string
	mfaEnabled    bool
	premiumType   PremiumType
	username      string
}

// AvatarHash gets the user's avatar hash.
func (u *User) AvatarHash() string {
	return u.avatarHash
}

// Bot returns whether the user is a bot.
func (u *User) Bot() bool {
	return u.bot
}

// Discriminator gets the user's 4 digit discriminator.
func (u *User) Discriminator() string {
	return u.discriminator
}

// Flags gets the user's profile flags, also called badges.
func (u *User) Flags() Flags {
	return u.flags
}

// ID gets the user's unqiue snowflake ID.
func (u *User) ID() *snowflake.Snowflake {
	return u.id
}

// Locale gets the user's Discord client locale.
func (u *User) Locale() string {
	return u.locale
}

// MFAEnabled returns whether the user has enabled multi-factor authentication.
func (u *User) MFAEnabled() bool {
	return u.mfaEnabled
}

// PremiumType gets the user's premium type which is the Nitro subscription type.
func (u *User) PremiumType() PremiumType {
	return u.premiumType
}

// Username gets the user's username.
func (u *User) Username() string {
	return u.username
}

// Flag is the type for user profile flag (or badge) constants.
type Flag uint

const (
	// None.
	FlagNone Flag = 0
	// Discord Employee.
	FlagDiscordEmployee Flag = 1 << 0
	// Discord Partner.
	FlagDiscordPartner Flag = 1 << 1
	// Hypesquad Events.
	FlagHypesquadEvents Flag = 1 << 2
	// Bug Hunter.
	FlagBugHunter Flag = 1 << 3
	// Hypesquad House of Bravery.
	FlagHouseBravery Flag = 1 << 6
	// Hypesquad House of Brilliance.
	FlagHouseBrilliance Flag = 1 << 7
	// Hypesquad House of Balance.
	FlagHouseBalance Flag = 1 << 8
	// Discord Nitro Early Supporter
	FlagEarlySupporter Flag = 1 << 9
	// Team User.
	FlagTeamUser Flag = 1 << 10
)

// Flags is the type for user profile flags, also called badges.
type Flags uint

// HasFlag checks whether the user's flag set contains the specified flag.
func (f Flags) HasFlag(flag Flag) bool {
	return uint(f)&uint(flag) == uint(flag)
}

// PremiumType is the type for Discord Nitro subscription type constants.
type PremiumType uint

const (
	// None.
	PremiumTypeNone PremiumType = iota
	// Nitro Classic.
	PremiumTypeNitroClassic
	// Nitro.
	PremiumTypeNitro
)
