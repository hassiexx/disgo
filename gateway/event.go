package gateway

import (
	"encoding/json"

	"github.com/hassieswift621/disgo/common/types"
	"golang.org/x/xerrors"
)

type event string

const (
	eventChannelCreate event = "CHANNEL_CREATE"
	eventReady         event = "READY"
)

func (s *Session) channelCreate(data json.RawMessage) error {
	var channel *types.Channel

	// Unmarshal data.
	if err := unmarshalRaw(data, &channel); err != nil {
		return xerrors.Errorf("failed to unmarshal channel create data", err)
	}

	// Store channel in state.
	s.state.AddChannel(channel)

	return nil
}

func (s *Session) ready(data json.RawMessage) error {
	var readyData readyData

	// Unmarshal data.
	if err := unmarshalRaw(data, &readyData); err != nil {
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
