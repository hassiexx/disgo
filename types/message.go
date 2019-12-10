package types

// Message is a struct for a message.
type Message struct {
	Activity         *MessageActivity    `json:"activity"`
	Application      *MessageApplication `json:"application"`
	Attachments      []*Attachment       `json:"attachments"`
	Author           *MessageAuthor      `json:"author"`
	ChannelID        string              `json:"channel_id"`
	Content          string              `json:"content"`
	EditedTimestamp  string              `json:"edited_timestamp"`
	Embeds           []*Embed            `json:"embeds"`
	Flags            uint                `json:"flags"`
	GuildID          string              `json:"guild_id"`
	ID               string              `json:"id"`
	Member           *Member             `json:"member"`
	MentionChannels  []*ChannelMention   `json:"mention_channels"`
	MentionEveryone  bool                `json:"mention_everyone"`
	MentionRoles     []string            `json:"mention_roles"`
	Mentions         []*UserMention      `json:"mentions"`
	MessageReference *MessageReference   `json:"message_reference"`
	Nonce            string              `json:"nonce"`
	Pinned           bool                `json:"pinned"`
	Reactions        []*Reaction         `json:"reactions"`
	Timestamp        string              `json:"timestamp"`
	TTS              bool                `json:"tts"`
	Type             uint                `json:"type"`
	WebhookID        string              `json:"webhook_id"`
}

// Attachment is a struct for a message attachment.
type Attachment struct {
	Filename string `json:"filename"`
	Height   uint   `json:"height"`
	ID       string `json:"id"`
	ProxyURL string `json:"proxy_url"`
	Size     uint   `json:"size"`
	URL      string `json:"url"`
	Width    uint   `json:"width"`
}

// ChannelMention is a struct for a channel mention in a message.
type ChannelMention struct {
	GuildID string `json:"guild_id"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    uint   `json:"type"`
}

// MessageActivity is a struct for a activity sent with rich-presence messages.
type MessageActivity struct {
	PartyID string `json:"party_id"`
	Type    uint   `json:"type"`
}

// MessageApplication is a struct for a application sent with rich-presence messages.
type MessageApplication struct {
	CoverImage  string `json:"cover_image"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

// MessageAuthor is a struct for a message author.
type MessageAuthor struct {
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
	Discriminator string `json:"discriminator"`
	ID            string `json:"id"`
	Username      string `json:"username"`
}

// Reaction is a struct for a message reaction.
type Reaction struct {
	Count uint   `json:"count"`
	Emoji *Emoji `json:"emoji"`
	Me    bool   `json:"me"`
}

// MessageReference is a struct for a message reference sent with a crossposted message.
type MessageReference struct {
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
	MessageID string `json:"message_id"`
}

// UserMention is a struct for a user mention in a message.
type UserMention struct {
	Avatar        string  `json:"avatar"`
	Bot           bool    `json:"bot"`
	Discriminator string  `json:"discriminator"`
	ID            string  `json:"id"`
	Username      string  `json:"username"`
	Member        *Member `json:"member"`
}
