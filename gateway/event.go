package gateway

import (
	"encoding/json"
	"github.com/hassieswift621/disgo/common/types"
	"golang.org/x/xerrors"
)

type event string

const (
	eventReady event = "READY"
)

func (s *Session) ready(data json.RawMessage) error {
	var readyData readyData

	// Unmarshal data.
	if err := unmarshalRaw(data, readyData); err != nil {
		return xerrors.Errorf("failed to unmarshal ready data", err)
	}

	// Store session ID.
	s.sessionID = readyData.SessionID

	// Store bot user in state.
	s.state.SetSelf(readyData.User)

	// Store guilds in state.
	s.state.AddGuildsReady(readyData.Guild)

	return nil
}

type readyData struct {
	Guild     []*types.Guild `json:"guilds"`
	SessionID string         `json:"session_id"`
	User      *types.User    `json:"user"`
}
