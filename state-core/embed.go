package statecore

// Embed is a struct for raw embed state data.
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

// EmbedAuthor is a struct for raw embed author state data.
type EmbedAuthor struct {
	IconURL  string `json:"icon_url"`
	Name     string `json:"name"`
	ProxyURL string `json:"proxy_url"`
	URL      string `json:"url"`
}

// EmbedField is a struct for raw embed field state data.
type EmbedField struct {
	Inline bool   `json:"inline"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

// EmbedFooter is a struct for raw embed footer state data.
type EmbedFooter struct {
	IconURL  string `json:"icon_url"`
	ProxyURL string `json:"proxy_url"`
	Text     string `json:"text"`
}

// EmbedImage is a struct for raw embed image state data.
type EmbedImage struct {
	Height   uint   `json:"height"`
	ProxyURL string `json:"proxy_url"`
	URL      string `json:"url"`
	Width    uint   `json:"width"`
}

// EmbedProvider is a struct for raw embed provider state data.
type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// EmbedThumbnail is a struct for raw embed thumbnail state data.
type EmbedThumbnail struct {
	Height   uint   `json:"height"`
	ProxyURL string `json:"proxy_url"`
	URL      string `json:"url"`
	Width    uint   `json:"width"`
}

// EmbedVideo is a struct for raw embed video state data.
type EmbedVideo struct {
	Height uint   `json:"height"`
	URL    string `json:"url"`
	Width  uint   `json:"width"`
}
