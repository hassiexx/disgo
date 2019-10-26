package statemem

import (
	"github.com/hassieswift621/disgo/common/types"
	"github.com/hassieswift621/disgo/statecore"
)

// State is a struct for state cache.
type State struct {
	channels             map[string]*types.Channel
	emojis               map[string]*types.Emoji
	guilds               map[string]*types.Guild
	members              map[string]map[string]*types.Member
	messages             map[string]*types.Message
	permissionOverwrites map[string]map[string]*types.PermissionOverwrite
	presences            map[string]map[string]*types.Presence
	roles                map[string]*types.Role
	users                map[string]*types.User
}

// Channel gets a channel by its ID.
func (s *State) Channel(id string) (*types.Channel, error) {
	channel, exists := s.channels[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return channel, nil
}

// Emoji gets an emoji by its ID.
func (s *State) Emoji(id string) (*types.Emoji, error) {
	emoji, exists := s.emojis[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return emoji, nil
}

// Guild gets a guild by its ID.
func (s *State) Guild(id string) (*types.Guild, error) {
	guild, exists := s.guilds[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return guild, nil
}

// Member gets a guild member.
func (s *State) Member(guildID string, memberID string) (*types.Member, error) {
	guild, exists := s.members[guildID]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	member, exists := guild[memberID]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return member, nil
}

// Message gets a message.
func (s *State) Message(id string) (*types.Message, error) {
	message, exists := s.messages[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return message, nil
}

// PermissionOverwrite gets a role or user permission overwrite for a channel.
func (s *State) PermissionOverwrite(channelID string, overwriteID string) (*types.PermissionOverwrite, error) {
	channel, exists := s.permissionOverwrites[channelID]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	overwrite, exists := channel[overwriteID]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return overwrite, nil
}

// Presence gets a user's guild presence.
func (s *State) Presence(guildID string, userID string) (*types.Presence, error) {
	guild, exists := s.presences[guildID]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	presence, exists := guild[userID]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return presence, nil
}

// Role gets a role.
func (s *State) Role(id string) (*types.Role, error) {
	role, exists := s.roles[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return role, nil
}

// User gets a user.
func (s *State) User(id string) (*types.User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return user, nil
}
