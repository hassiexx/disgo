package gateway

import (
	"encoding/json"

	"github.com/hassieswift621/disgo/types"
	"golang.org/x/xerrors"
)

const (
	eventChannelCreate string = "CHANNEL_CREATE"
	eventChannelUpdate string = "CHANNEL_UPDATE"
	eventReady         string = "READY"
)

func (s *Session) channelCreate(data json.RawMessage) error {
	var channel *types.Channel

	// Unmarshal data.
	if err := unmarshalRaw(data, &channel); err != nil {
		return xerrors.Errorf("failed to unmarshal channel create data", err)
	}

	// Store channel in state.
	channelState := s.state.AddChannel(channel)

	// Dispatch event.
	s.dispatcher.ChannelCreate(channelState)

	return nil
}

func (s *Session) channelUpdate(data json.RawMessage) error {
	var channel *types.Channel

	// Unmarshal data.
	if err := unmarshalRaw(data, &channel); err != nil {
		return xerrors.Errorf("failed to unmarshal channel update data", err)
	}

	// Store channel in state.
	channelState := s.state.AddChannel(channel)

	// Dispatch event.
	s.dispatcher.ChannelUpdate(channelState)

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

	// Dispatch event.
	s.dispatcher.Ready(s.shardID)

	return nil
}

type readyData struct {
	Guild     []*types.Guild `json:"guilds"`
	SessionID string         `json:"session_id"`
	User      *types.User    `json:"user"`
}
