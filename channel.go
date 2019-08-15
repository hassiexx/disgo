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

// ID gets the snowflake ID of the channel.
func (c *DMChannel) ID() *snowflake.Snowflake {
	return c.id
}

// LastMessageID gets the snowflake ID of the last message sent.
// This may or may not point to a deleted message.
func (c *DMChannel) LastMessageID() *snowflake.Snowflake {
	return c.lastMessageID
}

// Recipient gets the recipient user of the DM.
func (c *DMChannel) Recipient() *User {
	return c.recipient
}

// Type gets the channel type.
func (c *DMChannel) Type() ChannelType {
	return c.channelType
}

// GroupDMChannel is an immutable struct for a group DM channel.
type GroupDMChannel struct {
	applicationID *snowflake.Snowflake
	channelType   ChannelType
	iconHash      string
	id            *snowflake.Snowflake
	lastMessageID *snowflake.Snowflake
	name          string
	ownerID       *snowflake.Snowflake
	recipients    []*User
}

// ApplicationID gets the snowflake ID of the application
// if this channel has been created by a bot.
func (c *GroupDMChannel) ApplicationID() *snowflake.Snowflake {
	return c.applicationID
}

// IconHash gets the icon hash of the channel.
func (c *GroupDMChannel) IconHash() string {
	return c.iconHash
}

// ID gets the snowflake ID of the channel.
func (c *GroupDMChannel) ID() *snowflake.Snowflake {
	return c.id
}

// LastMessageID gets the snowflake ID of the last message sent.
// This may or may not point to a deleted message.
func (c *GroupDMChannel) LastMessageID() *snowflake.Snowflake {
	return c.lastMessageID
}

// Name gets the name of the channel.
func (c *GroupDMChannel) Name() string {
	return c.name
}

// OwnerID gets the snowflake ID of the user who created the DM.
func (c *GroupDMChannel) OwnerID() *snowflake.Snowflake {
	return c.ownerID
}

// Recipients gets the recipient users of the Group DM.
func (c *GroupDMChannel) Recipients() []*User {
	// Copy slice.
	recipients := make([]*User, len(c.recipients))
	copy(c.recipients, recipients)

	return recipients
}

// Type gets the channel type.
func (c *GroupDMChannel) Type() ChannelType {
	return c.channelType
}

// GuildChannel is the interface for a guild channel.
type GuildChannel interface {
	GuildID() *snowflake.Snowflake
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

// GuildID gets the snowflake ID of the guild.
func (c *GuildNewsChannel) GuildID() *snowflake.Snowflake {
	return c.guildID
}

// ID gets the snowflake ID of the channel.
func (c *GuildNewsChannel) ID() *snowflake.Snowflake {
	return c.id
}

// IsNSFW returns whether the channel is set as NSFW.
func (c *GuildNewsChannel) IsNSFW() bool {
	return c.nsfw
}

// LastMessageID gets the snowflake ID of the last message sent.
// This may or may not point to a deleted message.
func (c *GuildNewsChannel) LastMessageID() *snowflake.Snowflake {
	return c.lastMessageID
}

// LastPinTimestamp gets the timestamp of the last pinned message.
func (c *GuildNewsChannel) LastPinTimestamp() time.Time {
	return c.lastPinTimestamp
}

// ParentID gets the snowflake ID of the category if this channel is categorised.
func (c *GuildNewsChannel) ParentID() *snowflake.Snowflake {
	return c.parentID
}

// Type gets the channel type.
func (c *GuildNewsChannel) Type() ChannelType {
	return c.channelType
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
	permissionOverwrites []*PermissionOverwrite
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
	permissionOverwrites []*PermissionOverwrite
	position             uint
	userLimit            uint
}
