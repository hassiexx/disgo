package statemem

import (
	"sync"

	"github.com/hassieswift621/disgo/statecore"
	"github.com/hassieswift621/disgo/types"
)

// State is a struct for state cache.
type State struct {
	channels             map[uint64]*types.Channel
	emojis               map[uint64]*types.Emoji
	guilds               map[uint64]*types.Guild
	members              map[uint64]map[uint64]*types.Member
	messages             map[uint64]*types.Message
	permissionOverwrites map[uint64]map[uint64]*types.PermissionOverwrite
	presences            map[uint64]map[uint64]*types.Presence
	roles                map[uint64]*types.Role
	self                 *types.User
	users                map[uint64]*types.User
	sync.RWMutex
}

// New creates a new instance of state.
func New() *State {
	return &State{
		channels:             make(map[uint64]*types.Channel),
		emojis:               make(map[uint64]*types.Emoji),
		guilds:               make(map[uint64]*types.Guild),
		members:              make(map[uint64]map[uint64]*types.Member),
		permissionOverwrites: make(map[uint64]map[uint64]*types.PermissionOverwrite),
		presences:            make(map[uint64]map[uint64]*types.Presence),
		roles:                make(map[uint64]*types.Role),
		users:                make(map[uint64]*types.User),
	}
}

// AddChannel adds a channel.
func (s *State) AddChannel(channel *types.Channel) {
	s.Lock()
	defer s.Unlock()

	// Check if the channel is a DM.
	if channel.Type == types.ChannelTypeDM || channel.Type == types.ChannelTypeGroupDM {
		// Extract recipients and add to users map and channel recipient hash set.
		recipients := channel.Recipients
		channel.Recipients = nil
		channel.RecipientSet = types.NewUInt64HashSet()

		for _, recipient := range recipients {
			user, exists := s.users[recipient.ID]
			if exists {
				// Update user.
				user.AvatarHash = recipient.AvatarHash
				user.Discriminator = recipient.Discriminator
				user.Username = recipient.Username
			} else {
				s.users[recipient.ID] = recipient
			}
			channel.RecipientSet.Add(recipient.ID)
		}

		// Add channel to map.
		s.channels[channel.ID] = channel

		return
	}

	// Extract permission overwrites from channel.
	overwrites := channel.PermissionOverwrites
	channel.PermissionOverwrites = nil

	// Add overwrites to map.
	for _, overwrite := range overwrites {
		overwriteMap, exists := s.permissionOverwrites[channel.ID]
		if !exists {
			overwriteMap = make(map[uint64]*types.PermissionOverwrite)
		}
		overwriteMap[overwrite.ID] = overwrite
		s.permissionOverwrites[channel.ID] = overwriteMap
	}

	// Add channel to map.
	s.channels[channel.ID] = channel

	// If the channel is a guild channel, add the channel ID to the guild channel hash set.
	if channel.GuildID != 0 {
		s.guilds[channel.GuildID].ChannelSet.Add(channel.ID)
	}
}

// AddGuildsReady adds unavailable guilds from the ready event.
func (s *State) AddGuildsReady(guilds []*types.Guild) {
	s.Lock()
	defer s.Unlock()

	for _, guild := range guilds {
		s.guilds[guild.ID] = &types.Guild{
			ID:          guild.ID,
			Unavailable: true,
		}
	}
}

// Channel gets a channel by its ID.
func (s *State) Channel(id uint64) (*types.Channel, error) {
	s.RLock()
	defer s.RUnlock()

	channel, exists := s.channels[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return channel, nil
}

// Emoji gets an emoji by its ID.
func (s *State) Emoji(id uint64) (*types.Emoji, error) {
	s.RLock()
	defer s.RUnlock()

	emoji, exists := s.emojis[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return emoji, nil
}

// Guild gets a guild by its ID.
func (s *State) Guild(id uint64) (*types.Guild, error) {
	s.RLock()
	defer s.RUnlock()

	guild, exists := s.guilds[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return guild, nil
}

// Member gets a guild member.
func (s *State) Member(guildID uint64, memberID uint64) (*types.Member, error) {
	s.RLock()
	defer s.RUnlock()

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
func (s *State) Message(id uint64) (*types.Message, error) {
	s.RLock()
	defer s.RUnlock()

	message, exists := s.messages[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return message, nil
}

// PermissionOverwrite gets a role or user permission overwrite for a channel.
func (s *State) PermissionOverwrite(channelID uint64, overwriteID uint64) (*types.PermissionOverwrite, error) {
	s.RLock()
	defer s.RUnlock()

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
func (s *State) Presence(guildID uint64, userID uint64) (*types.Presence, error) {
	s.RLock()
	defer s.RUnlock()

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
func (s *State) Role(id uint64) (*types.Role, error) {
	s.RLock()
	defer s.RUnlock()

	role, exists := s.roles[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return role, nil
}

// Self gets the bot user.
func (s *State) Self() (*types.User, error) {
	s.RLock()
	defer s.RUnlock()

	if s.self != nil {
		return nil, statecore.ErrNotFound
	}

	return s.self, nil
}

// SetSelf sets the bot user.
func (s *State) SetSelf(self *types.User) {
	s.Lock()
	defer s.Unlock()

	s.self = self
}

// User gets a user.
func (s *State) User(id uint64) (*types.User, error) {
	s.RLock()
	defer s.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, statecore.ErrNotFound
	}

	return user, nil
}
