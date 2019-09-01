package statecore

// Embed is a struct for an embed state.
type Embed struct {
	Author      *EmbedAuthor
	Colour      uint
	Description string
	Fields      []*EmbedField
	Footer      *EmbedFooter
	Image       *EmbedImage
	Provider    *EmbedProvider
	Thumbnail   *EmbedThumbnail
	Timestamp   string
	Title       string
	Type        string
	URL         string
	Video       *EmbedVideo
}

// EmbedAuthor is a struct for an embed author state.
type EmbedAuthor struct {
	IconURL  string
	Name     string
	ProxyURL string
	URL      string
}

// EmbedField is a struct for an embed field state.
type EmbedField struct {
	Inline bool
	Name   string
	Value  string
}

// EmbedFooter is a struct for an embed footer state.
type EmbedFooter struct {
	IconURL  string
	ProxyURL string
	Text     string
}

// EmbedImage is a struct for an embed image state.
type EmbedImage struct {
	Height   uint
	ProxyURL string
	URL      string
	Width    uint
}

// EmbedProvider is a struct for an embed provider state.
type EmbedProvider struct {
	Name string
	URL  string
}

// EmbedThumbnail is a struct for an embed thumbnail state.
type EmbedThumbnail struct {
	Height   uint
	ProxyURL string
	URL      string
	Width    uint
}

// EmbedVideo is a struct for an embed video state.
type EmbedVideo struct {
	Height uint
	URL    string
	Width  uint
}
