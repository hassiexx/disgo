package gateway

import (
	"encoding/json"

	"github.com/hassieswift621/disgo/types"
	"golang.org/x/xerrors"
)

type gatewayEvent string

const (
	eventChannelCreate gatewayEvent = "CHANNEL_CREATE"
	eventChannelUpdate gatewayEvent = "CHANNEL_UPDATE"
	eventReady         gatewayEvent = "READY"
)

func (s *Session) channelCreate(data json.RawMessage) error {
	var channel *types.Channel

	// Unmarshal data.
	if err := unmarshalRaw(data, &channel); err != nil {
		return xerrors.Errorf("failed to unmarshal channel create data", err)
	}

	// Store channel in state.
	channel = s.state.AddChannel(channel)

	// Dispatch event.
	s.dispatcher.ChannelCreate(*channel)

	return nil
}

func (s *Session) channelUpdate(data json.RawMessage) error {
	var channel *types.Channel

	// Unmarshal data.
	if err := unmarshalRaw(data, &channel); err != nil {
		return xerrors.Errorf("failed to unmarshal channel update data", err)
	}

	// Store channel in state.
	channel = s.state.AddChannel(channel)

	// Dispatch event.
	s.dispatcher.ChannelUpdate(*channel)

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
