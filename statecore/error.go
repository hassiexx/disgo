package statecore

import (
	"golang.org/x/xerrors"
)

// ErrNotFound is the error returned when an object does not exist in the state cache.
var ErrNotFound = xerrors.New("disgo: object not found in state")
