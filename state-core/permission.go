package statecore

// PermissionOverwrite is a struct for a permission overwrite state.
type PermissionOverwrite struct {
	Allow uint
	Deny  uint
	ID    string
	Type  string
}
