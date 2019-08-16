package statecore

// PermissionOverwriteState is the struct used to store raw permission overwrite state data.
type PermissionOverwriteState struct {
	// Allowed permissions bit set.
	Allow uint `json:"allow"`
	// Denied permissions bit set.
	Deny uint `json:"deny"`
	// Role/User ID depending on the type of overwrite.
	ID string `json:"id"`
	// The type of overwrite. Either 'role' or 'member'.
	Type string `json:"type"`
}
