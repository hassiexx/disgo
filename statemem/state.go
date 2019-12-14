package statemem

import (
	"sync"

	"github.com/hassieswift621/disgo/statecore"
	"github.com/hassieswift621/disgo/types"
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
	self                 *types.User
	users                map[string]*types.User
	sync.RWMutex
}

// New creates a new instance of state.
func New() *State {
	return &State{
		channels:             make(map[string]*types.Channel),
		emojis:               make(map[string]*types.Emoji),
		guilds:               make(map[string]*types.Guild),
		members:              make(map[string]map[string]*types.Member),
		permissionOverwrites: make(map[string]map[string]*types.PermissionOverwrite),
		presences:            make(map[string]map[string]*types.Presence),
		roles:                make(map[string]*types.Role),
		users:                make(map[string]*types.User),
	}
}

// AddChannel adds a channel.
func (s *State) AddChannel(channel *types.Channel) types.Channel {
	s.Lock()
	defer s.Unlock()

	// Check if the channel is a DM.
	if channel.Type == types.ChannelTypeDM || channel.Type == types.ChannelTypeGroupDM {
		// Extract recipients and add to users map and channel recipient hash set.
		recipients := channel.Recipients
		channel.Recipients = nil
		channel.RecipientSet = types.NewStringHashSet()

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

		return *channel
	}

	// Extract permission overwrites from channel.
	overwrites := channel.PermissionOverwrites
	channel.PermissionOverwrites = nil

	// Add overwrites to map.
	for _, overwrite := range overwrites {
		overwriteMap, exists := s.permissionOverwrites[channel.ID]
		if !exists {
			overwriteMap = make(map[string]*types.PermissionOverwrite)
		}
		overwriteMap[overwrite.ID] = overwrite
		s.permissionOverwrites[channel.ID] = overwriteMap
	}

	// Add channel to map.
	s.channels[channel.ID] = channel

	// If the channel is a guild channel, add the channel ID to the guild channel hash set.
	if channel.GuildID != "" {
		s.guilds[channel.GuildID].ChannelSet.Add(channel.ID)
	}

	return *channel
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
func (s *State) Channel(id string) (types.Channel, error) {
	s.RLock()
	defer s.RUnlock()

	channel, exists := s.channels[id]
	if !exists {
		return types.Channel{}, statecore.ErrNotFound
	}

	return *channel, nil
}

// Emoji gets an emoji by its ID.
func (s *State) Emoji(id string) (types.Emoji, error) {
	s.RLock()
	defer s.RUnlock()

	emoji, exists := s.emojis[id]
	if !exists {
		return types.Emoji{}, statecore.ErrNotFound
	}

	return *emoji, nil
}

// Guild gets a guild by its ID.
func (s *State) Guild(id string) (types.Guild, error) {
	s.RLock()
	defer s.RUnlock()

	guild, exists := s.guilds[id]
	if !exists {
		return types.Guild{}, statecore.ErrNotFound
	}

	return *guild, nil
}

// Member gets a guild member.
func (s *State) Member(guildID string, memberID string) (types.Member, error) {
	s.RLock()
	defer s.RUnlock()

	guild, exists := s.members[guildID]
	if !exists {
		return types.Member{}, statecore.ErrNotFound
	}

	member, exists := guild[memberID]
	if !exists {
		return types.Member{}, statecore.ErrNotFound
	}

	return *member, nil
}

// Message gets a message.
func (s *State) Message(id string) (types.Message, error) {
	s.RLock()
	defer s.RUnlock()

	message, exists := s.messages[id]
	if !exists {
		return types.Message{}, statecore.ErrNotFound
	}

	return *message, nil
}

// PermissionOverwrite gets a role or user permission overwrite for a channel.
func (s *State) PermissionOverwrite(channelID string, overwriteID string) (types.PermissionOverwrite, error) {
	s.RLock()
	defer s.RUnlock()

	channel, exists := s.permissionOverwrites[channelID]
	if !exists {
		return types.PermissionOverwrite{}, statecore.ErrNotFound
	}

	overwrite, exists := channel[overwriteID]
	if !exists {
		return types.PermissionOverwrite{}, statecore.ErrNotFound
	}

	return *overwrite, nil
}

// Presence gets a user's guild presence.
func (s *State) Presence(guildID string, userID string) (types.Presence, error) {
	s.RLock()
	defer s.RUnlock()

	guild, exists := s.presences[guildID]
	if !exists {
		return types.Presence{}, statecore.ErrNotFound
	}

	presence, exists := guild[userID]
	if !exists {
		return types.Presence{}, statecore.ErrNotFound
	}

	return *presence, nil
}

// Role gets a role.
func (s *State) Role(id string) (types.Role, error) {
	s.RLock()
	defer s.RUnlock()

	role, exists := s.roles[id]
	if !exists {
		return types.Role{}, statecore.ErrNotFound
	}

	return *role, nil
}

// Self gets the bot user.
func (s *State) Self() (types.User, error) {
	s.RLock()
	defer s.RUnlock()

	if s.self != nil {
		return types.User{}, statecore.ErrNotFound
	}

	return *s.self, nil
}

// SetSelf sets the bot user.
func (s *State) SetSelf(self *types.User) {
	s.Lock()
	defer s.Unlock()

	s.self = self
}

// User gets a user.
func (s *State) User(id string) (types.User, error) {
	s.RLock()
	defer s.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return types.User{}, statecore.ErrNotFound
	}

	return *user, nil
}
