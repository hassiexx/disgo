package json

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
