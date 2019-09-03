package json

// Presence is a struct for a user's presence.
type Presence struct {
	GuildID string   `json:"guild_id"`
	Roles   []string `json:"roles"`
	Status  string   `json:"status"`
	User    *User    `json:"user"`
}

// Activity is a struct for a presence activity.
type Activity struct {
	ApplicationID string             `json:"application_id"`
	Assets        *ActivityAsset     `json:"assets"`
	Details       string             `json:"details"`
	Flags         uint               `json:"flags"`
	Instance      bool               `json:"instance"`
	Name          string             `json:"name"`
	Secrets       *ActivitySecret    `json:"secrets"`
	State         string             `json:"state"`
	Timestamps    *ActivityTimestamp `json:"timestamps"`
	Type          uint               `json:"type"`
	URL           string             `json:"url"`
}

// ActivityAsset is a struct for activity assets.
type ActivityAsset struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
}

// ActivityParty is a struct for a activity party.
type ActivityParty struct {
	ID   string `json:"id"`
	Size []uint `json:"size"`
}

// ActivitySecret is a struct for activity secrets.
type ActivitySecret struct {
	Join     string `json:"join"`
	Match    string `json:"match"`
	Spectate string `json:"spectate"`
}

// ActivityTimestamp is a struct for activity timestamps.
type ActivityTimestamp struct {
	Start uint `json:"start"`
	End   uint `json:"end"`
}
