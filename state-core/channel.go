package statecore

// Channel is a struct for a channel state.
type Channel struct {
	ApplicationID    string
	Bitrate          uint
	GuildID          string
	Icon             string
	ID               string
	LastMessageID    string
	LastPinTimestamp string
	Name             string
	NSFW             bool
	ParentID         string
	Position         uint
	RateLimitPerUser uint
	Recipients       []*User
	Topic            string
	Type             uint
	UserLimit        uint
}
