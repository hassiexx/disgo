package disgo

// Embed is an immutable struct for an embed.
type Embed struct {
	author      *EmbedAuthor
	colour      uint
	description string
	embedType   string
	fields      []*EmbedField
	footer      *EmbedFooter
	image       *EmbedImage
	provider    *EmbedProvider
	thumbnail   *EmbedThumbnail
	timestamp   Timestamp
	title       string
	url         string
	video       *EmbedVideo
}

// Author gets the embed author.
func (e *Embed) Author() *EmbedAuthor {
	return e.author
}

// Colour gets the embed colour.
func (e *Embed) Colour() uint {
	return e.colour
}

// Description gets the embed description.
func (e *Embed) Description() string {
	return e.description
}

// Fields gets the embed fields.
func (e *Embed) Fields() []*EmbedField {
	// Copy slice.
	fields := make([]*EmbedField, len(e.fields))
	copy(fields, e.fields)

	return fields
}

// Footer gets the embed footer.
func (e *Embed) Footer() *EmbedFooter {
	return e.footer
}

// Type gets the embed type.
func (e *Embed) Type() string {
	return e.embedType
}

// Image gets the embed image.
func (e *Embed) Image() *EmbedImage {
	return e.image
}

// Provider gets the embed provider.
func (e *Embed) Provider() *EmbedProvider {
	return e.provider
}

// Thumbnail gets the embed thumbnail.
func (e *Embed) Thumbnail() *EmbedThumbnail {
	return e.thumbnail
}

// Timestamp gets the embed timestamp.
func (e *Embed) Timestamp() Timestamp {
	return e.timestamp
}

// Title gets the embed title.
func (e *Embed) Title() string {
	return e.title
}

// URL gets the embed URL.
func (e *Embed) URL() string {
	return e.url
}

// Video gets the embed video.
func (e *Embed) Video() *EmbedVideo {
	return e.video
}

// EmbedAuthor is an immutable struct for an embed author.
type EmbedAuthor struct {
	iconURL  string
	name     string
	proxyURL string
	url      string
}

// IconURL gets the URL of the author icon.
func (a *EmbedAuthor) IconURL() string {
	return a.iconURL
}

// Name gets the name of the author.
func (a *EmbedAuthor) Name() string {
	return a.name
}

// ProxyURL gets the proxied URL of the author icon.
func (a *EmbedAuthor) ProxyURL() string {
	return a.proxyURL
}

// URL gets the author's URL.
func (a *EmbedAuthor) URL() string {
	return a.url
}

// EmbedField is an immutable struct for an embed field.
type EmbedField struct {
	inline bool
	name   string
	value  string
}

// IsInline returns whether the field should display inline.
func (f *EmbedField) IsInline() bool {
	return f.inline
}

// Name gets the name of the field.
func (f *EmbedField) Name() string {
	return f.name
}

// Value gets the value of the field.
func (f *EmbedField) Value() string {
	return f.value
}

// EmbedFooter is an immutable struct for an embed footer.
type EmbedFooter struct {
	iconURL  string
	proxyURL string
	text     string
}

// IconURL gets the URL of the footer icon.
func (f *EmbedFooter) IconURL() string {
	return f.iconURL
}

// ProxyURL gets the proxied URL of the footer icon.
func (f *EmbedFooter) ProxyURL() string {
	return f.proxyURL
}

// Text gets the footer text.
func (f *EmbedFooter) Text() string {
	return f.text
}

// EmbedImage is an immutable struct for an embed image.
type EmbedImage struct {
	height   uint
	proxyURL string
	url      string
	width    uint
}

// Height gets the height of the image.
func (i *EmbedImage) Height() uint {
	return i.height
}

// ProxyURL gets the proxied URL of the image.
func (i *EmbedImage) ProxyURL() string {
	return i.proxyURL
}

// URL gets the source URL of the image.
func (i *EmbedImage) URL() string {
	return i.url
}

// Width gets the width of the image.
func (i *EmbedImage) Width() uint {
	return i.width
}

// EmbedProvider is an immutable struct for an embed provider.
type EmbedProvider struct {
	name string
	url  string
}

// Name gets the name of the provider.
func (p *EmbedProvider) Name() string {
	return p.name
}

// URL gets the URL of the provider.
func (p *EmbedProvider) URL() string {
	return p.url
}

// EmbedThumbnail is an immutable struct for an embed thumbnail.
type EmbedThumbnail struct {
	height   uint
	proxyURL string
	url      string
	width    uint
}

// Height gets the height of the thumbnail.
func (t *EmbedThumbnail) Height() uint {
	return t.height
}

// ProxyURL gets the proxied URL of the thumbnail.
func (t *EmbedThumbnail) ProxyURL() string {
	return t.proxyURL
}

// Width gets the width of the thumbnail.
func (t *EmbedThumbnail) Width() uint {
	return t.width
}

// URL gets the source URL of the thumbail.
func (t *EmbedThumbnail) URL() string {
	return t.url
}

// EmbedVideo is an immutable struct for an embed video.
type EmbedVideo struct {
	height uint
	url    string
	width  uint
}

// Height gets the height of the video.
func (v *EmbedVideo) Height() uint {
	return v.height
}

// URL gets the source URL of the video.
func (v *EmbedVideo) URL() string {
	return v.url
}

// Width gets the width of the video.
func (v *EmbedVideo) Width() uint {
	return v.width
}
