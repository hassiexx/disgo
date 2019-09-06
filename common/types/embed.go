package types

// Embed is a struct for an embed.
type Embed struct {
	Author      *EmbedAuthor    `json:"author"`
	Colour      uint            `json:"color"`
	Description string          `json:"description"`
	Fields      []*EmbedField   `json:"fields"`
	Footer      *EmbedFooter    `json:"footer"`
	Image       *EmbedImage     `json:"image"`
	Provider    *EmbedProvider  `json:"provider"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail"`
	Timestamp   string          `json:"timestamp"`
	Title       string          `json:"title"`
	Type        string          `json:"type"`
	URL         string          `json:"url"`
	Video       *EmbedVideo     `json:"video"`
}

// EmbedAuthor is a struct for an embed author.
type EmbedAuthor struct {
	IconURL  string `json:"icon_url"`
	Name     string `json:"name"`
	ProxyURL string `json:"proxy_url"`
	URL      string `json:"url"`
}

// EmbedField is a struct for an embed field state.
type EmbedField struct {
	Inline bool   `json:"inline"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

// EmbedFooter is a struct for an embed footer.
type EmbedFooter struct {
	IconURL  string `json:"icon_url"`
	ProxyURL string `json:"proxy_url"`
	Text     string `json:"text"`
}

// EmbedImage is a struct for an embed image.
type EmbedImage struct {
	Height   uint   `json:"height"`
	ProxyURL string `json:"proxy_url"`
	URL      string `json:"url"`
	Width    uint   `json:"width"`
}

// EmbedProvider is a struct for an embed provider.
type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// EmbedThumbnail is a struct for an embed thumbnail.
type EmbedThumbnail struct {
	Height   uint   `json:"height"`
	ProxyURL string `json:"proxy_url"`
	URL      string `json:"url"`
	Width    uint   `json:"width"`
}

// EmbedVideo is a struct for an embed video.
type EmbedVideo struct {
	Height uint   `json:"height"`
	URL    string `json:"url"`
	Width  uint   `json:"width"`
}
