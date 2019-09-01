package statecore

// PermissionOverwrite is a struct for raw permission overwrite state data.
type PermissionOverwrite struct {
	// Allowed permissions bit set.
	Allow uint `json:"allow"`
	// Denied permissions bit set.
	Deny uint `json:"deny"`
	// Role/User ID depending on the type of overwrite.
	ID string `json:"id"`
	// The type of overwrite. Either 'role' or 'member'.
	Type string `json:"type"`
}
