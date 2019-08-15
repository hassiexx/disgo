package statecore

// PermissionOverwriteState is the struct used to store raw permission overwrite state data.
type PermissionOverwriteState struct {
	Allow uint   `json:"allow"`
	Deny  uint   `json:"deny"`
	ID    string `json:"id"`
	Type  string `json:"type"`
}
