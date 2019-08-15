package data

import (
	"time"

	snowflake "github.com/hassieswift621/discord-goflake"
)

// ChannelType is the type for channel type constants.
type ChannelType uint

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
)

// Channel is the interface for a channel.
type Channel interface {
	ID() *snowflake.Snowflake
	Type() ChannelType
}

// DMChannel is an immutable struct for a DM channel.
type DMChannel struct {
	channelType   ChannelType
	id            *snowflake.Snowflake
	lastMessageID *snowflake.Snowflake
	recipient     *User
}

// GroupDMChannel is an immutable struct for a group DM channel.
type GroupDMChannel struct {
	applicationID *snowflake.Snowflake
	channelType   ChannelType
	iconHash      string
	id            *snowflake.Snowflake
	ownerID       *snowflake.Snowflake
	lastMessageID *snowflake.Snowflake
	recipients    []*User
}

// GuildChannel is the interface for a guild channel.
type GuildChannel interface {
	Guild() *snowflake.Snowflake
	ID() *snowflake.Snowflake
	Name() string
	ParentID() *snowflake.Snowflake
	PermissionOverwrites()
	Position() uint
	Type() ChannelType
}

// GuildNewsChannel is an immutable struct for a guild news channel.
type GuildNewsChannel struct {
	channelType          ChannelType
	guildID              *snowflake.Snowflake
	id                   *snowflake.Snowflake
	lastMessageID        *snowflake.Snowflake
	lastPinTimestamp     time.Time
	name                 string
	nsfw                 bool
	parentID             *snowflake.Snowflake
	permissionOverwrites []struct{}
	position             uint
	rateLimitPerUser     uint
	topic                string
}

// GuildTextChannel is an immutable struct for a guild text channel.
type GuildTextChannel struct {
	channelType          ChannelType
	guildID              *snowflake.Snowflake
	id                   *snowflake.Snowflake
	lastMessageID        *snowflake.Snowflake
	lastPinTimestamp     time.Time
	name                 string
	nsfw                 bool
	parentID             *snowflake.Snowflake
	permissionOverwrites []struct{}
	position             uint
	rateLimitPerUser     uint
	topic                string
}

// GuildVoiceChannel is an immutable struct for a guild voice channel.
type GuildVoiceChannel struct {
	bitrate              uint
	channelType          ChannelType
	guildID              *snowflake.Snowflake
	id                   *snowflake.Snowflake
	name                 string
	nsfw                 bool
	parentID             *snowflake.Snowflake
	permissionOverwrites []struct{}
	position             uint
	userLimit            uint
}
