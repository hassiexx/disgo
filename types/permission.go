package types

// PermissionOverwrite is a struct for a permission overwrite.
type PermissionOverwrite struct {
	Allow uint   `json:"allow"`
	Deny  uint   `json:"deny"`
	ID    uint64 `json:"id,string"`
	Type  string `json:"type"`
}
