package disgo

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
	copy(recipients, c.recipients)

	return recipients
}

// Type gets the channel type.
func (c *GroupDMChannel) Type() ChannelType {
	return c.channelType
}

// GuildChannel is the interface for a guild channel.
type GuildChannel interface {
	// GuildID gets the snowflake ID of the guild.
	GuildID() *snowflake.Snowflake
	// ID gets the snowflake ID of the channel.
	ID() *snowflake.Snowflake
	// Name gets the name of the channel.
	Name() string
	// ParentID gets the snowflake ID of the category if this channel is categorised.
	ParentID() *snowflake.Snowflake
	// PermissionOverwrites gets the permission overwrites for the channel.
	PermissionOverwrites() []*PermissionOverwrite
	// Position gets the position of the channel.
	Position() uint
	// Type gets the channel type.
	Type() ChannelType
}

// GuildChannelCategory is an immutable struct for a guild channel category.
type GuildChannelCategory struct {
	channelType          ChannelType
	guildID              *snowflake.Snowflake
	id                   *snowflake.Snowflake
	name                 string
	nsfw                 bool
	parentID             *snowflake.Snowflake
	permissionOverwrites []*PermissionOverwrite
	position             uint
}

// GuildID gets the snowflake ID of the guild.
func (c *GuildChannelCategory) GuildID() *snowflake.Snowflake {
	return c.guildID
}

// ID gets the snowflake ID of the category.
func (c *GuildChannelCategory) ID() *snowflake.Snowflake {
	return c.id
}

// IsNSFW returns whether the category is set as NSFW.
func (c *GuildChannelCategory) IsNSFW() bool {
	return c.nsfw
}

// Name gets the name of the category.
func (c *GuildChannelCategory) Name() string {
	return c.name
}

// ParentID gets the snowflake ID of the category.
// For now this will return nil as Discord does not support sub-categories.
// However, this has been left in as parent ID is returned with the response.
func (c *GuildChannelCategory) ParentID() *snowflake.Snowflake {
	return c.parentID
}

// PermissionOverwrites gets the permission overwrites for the category.
func (c *GuildChannelCategory) PermissionOverwrites() []*PermissionOverwrite {
	// Copy slice.
	overwrites := make([]*PermissionOverwrite, len(c.permissionOverwrites))
	copy(overwrites, c.permissionOverwrites)

	return overwrites
}

// Position gets the position of the category.
func (c *GuildChannelCategory) Position() uint {
	return c.position
}

// Type gets the channel type.
func (c *GuildChannelCategory) Type() ChannelType {
	return c.channelType
}

// GuildNewsChannel is an immutable struct for a guild news channel.
type GuildNewsChannel GuildTextChannel

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

// GuildID gets the snowflake ID of the guild.
func (c *GuildTextChannel) GuildID() *snowflake.Snowflake {
	return c.guildID
}

// ID gets the snowflake ID of the channel.
func (c *GuildTextChannel) ID() *snowflake.Snowflake {
	return c.id
}

// IsNSFW returns whether the channel is set as NSFW.
func (c *GuildTextChannel) IsNSFW() bool {
	return c.nsfw
}

// LastMessageID gets the snowflake ID of the last message sent.
// This may or may not point to a deleted message.
func (c *GuildTextChannel) LastMessageID() *snowflake.Snowflake {
	return c.lastMessageID
}

// LastPinTimestamp gets the timestamp of the last pinned message.
func (c *GuildTextChannel) LastPinTimestamp() time.Time {
	return c.lastPinTimestamp
}

// Name gets the name of the channel.
func (c *GuildTextChannel) Name() string {
	return c.name
}

// ParentID gets the snowflake ID of the category if this channel is categorised.
func (c *GuildTextChannel) ParentID() *snowflake.Snowflake {
	return c.parentID
}

// PermissionOverwrites gets the permission overwrites for the channel.
func (c *GuildTextChannel) PermissionOverwrites() []*PermissionOverwrite {
	// Copy slice.
	overwrites := make([]*PermissionOverwrite, len(c.permissionOverwrites))
	copy(overwrites, c.permissionOverwrites)

	return overwrites
}

// Position gets the position of the channel.
func (c *GuildTextChannel) Position() uint {
	return c.position
}

// RateLimitPerUser gets the rate limit for the channel in seconds.
func (c *GuildTextChannel) RateLimitPerUser() uint {
	return c.rateLimitPerUser
}

// Topic gets the channel topic.
func (c *GuildTextChannel) Topic() string {
	return c.topic
}

// Type gets the channel type.
func (c *GuildTextChannel) Type() ChannelType {
	return c.channelType
}

// GuildStoreChannel is an immutable struct for a guild store channel.
type GuildStoreChannel struct {
	channelType          ChannelType
	guildID              *snowflake.Snowflake
	id                   *snowflake.Snowflake
	name                 string
	nsfw                 bool
	parentID             *snowflake.Snowflake
	permissionOverwrites []*PermissionOverwrite
	position             uint
}

// GuildID gets the snowflake ID of the guild.
func (c *GuildStoreChannel) GuildID() *snowflake.Snowflake {
	return c.guildID
}

// ID gets the snowflake ID of the channel.
func (c *GuildStoreChannel) ID() *snowflake.Snowflake {
	return c.id
}

// IsNSFW returns whether the channel is set as NSFW.
func (c *GuildStoreChannel) IsNSFW() bool {
	return c.nsfw
}

// Name gets the name of the channel.
func (c *GuildStoreChannel) Name() string {
	return c.name
}

// ParentID gets the snowflake ID of the category if this channel is categorised.
func (c *GuildStoreChannel) ParentID() *snowflake.Snowflake {
	return c.parentID
}

// PermissionOverwrites gets the permission overwrites for the channel.
func (c *GuildStoreChannel) PermissionOverwrites() []*PermissionOverwrite {
	// Copy slice.
	overwrites := make([]*PermissionOverwrite, len(c.permissionOverwrites))
	copy(overwrites, c.permissionOverwrites)

	return overwrites
}

// Position gets the position of the channel.
func (c *GuildStoreChannel) Position() uint {
	return c.position
}

// Type gets the channel type.
func (c *GuildStoreChannel) Type() ChannelType {
	return c.channelType
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

// Bitrate gets the bitrate of the channel in bits.
func (c *GuildVoiceChannel) Bitrate() uint {
	return c.bitrate
}

// GuildID gets the snowflake ID of the guild.
func (c *GuildVoiceChannel) GuildID() *snowflake.Snowflake {
	return c.guildID
}

// ID gets the snowflake ID of the channel.
func (c *GuildVoiceChannel) ID() *snowflake.Snowflake {
	return c.id
}

// IsNSFW returns whether the channel is set as NSFW.
func (c *GuildVoiceChannel) IsNSFW() bool {
	return c.nsfw
}

// Name gets the name of the channel.
func (c *GuildVoiceChannel) Name() string {
	return c.name
}

// ParentID gets the snowflake ID of the category if this channel is categorised.
func (c *GuildVoiceChannel) ParentID() *snowflake.Snowflake {
	return c.parentID
}

// PermissionOverwrites gets the permission overwrites for the channel.
func (c *GuildVoiceChannel) PermissionOverwrites() []*PermissionOverwrite {
	// Copy slice.
	overwrites := make([]*PermissionOverwrite, len(c.permissionOverwrites))
	copy(overwrites, c.permissionOverwrites)

	return overwrites
}

// Position gets the position of the channel.
func (c *GuildVoiceChannel) Position() uint {
	return c.position
}

// Type gets the channel type.
func (c *GuildVoiceChannel) Type() ChannelType {
	return c.channelType
}

// UserLimit gets the user limit of the channel.
func (c *GuildVoiceChannel) UserLimit() uint {
	return c.userLimit
}
