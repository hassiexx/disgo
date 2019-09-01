package json

// PermissionOverwrite is a struct for a permission overwrite.
type PermissionOverwrite struct {
	Allow uint   `json:"allow"`
	Deny  uint   `json:"deny"`
	ID    string `json:"id"`
	Type  string `json:"type"`
}
