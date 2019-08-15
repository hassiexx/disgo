package data

import snowflake "github.com/hassieswift621/discord-goflake"

// Permission is the type for permission constants.
type Permission uint

const (
	CreateInstantInvite Permission = 0x00000001
	KickMembers         Permission = 0x00000002
	BanMembers          Permission = 0x00000004
	Administrator       Permission = 0x00000008
	ManageChannels      Permission = 0x00000010
	ManageGuild         Permission = 0x00000020
	AddReactions        Permission = 0x00000040
	ViewAuditLog        Permission = 0x00000080
	ViewChannel         Permission = 0x00000400
	SendMessages        Permission = 0x00000800
	SendTTSMessages     Permission = 0x00001000
	ManageMessages      Permission = 0x00002000
	EmbedLinks          Permission = 0x00004000
	AttachFiles         Permission = 0x00008000
	ReadMessageHistory  Permission = 0x00010000
	MentionEveryone     Permission = 0x00020000
	UseExternalEmojis   Permission = 0x00040000
	Connect             Permission = 0x00100000
	Speak               Permission = 0x00200000
	MuteMembers         Permission = 0x00400000
	DeafenMembers       Permission = 0x00800000
	MoveMembers         Permission = 0x01000000
	UseVAD              Permission = 0x02000000
	PrioritySpeaker     Permission = 0x00000100
	ChangeNickname      Permission = 0x04000000
	ManageNicknames     Permission = 0x08000000
	ManageRoles         Permission = 0x10000000
	ManageWebhooks      Permission = 0x20000000
	ManageEmojis        Permission = 0x40000000
)

// PermissionOverwrite is an immutable struct for a permission overwrite.
type PermissionOverwrite struct {
	allow         PermissionBitSet
	deny          PermissionBitSet
	id            *snowflake.Snowflake
	overwriteType OverwriteType
}

// Allow gets the permission bit set for allowed permissions.
func (p *PermissionOverwrite) Allow() PermissionBitSet {
	return p.allow
}

// Deny gets the permission bit set for denied permissions.
func (p *PermissionOverwrite) Deny() PermissionBitSet {
	return p.deny
}

// ID gets the role ID or user ID depending on the type of overwrite.
func (p *PermissionOverwrite) ID() *snowflake.Snowflake {
	return p.id
}

// Type gets the overwrite type.
func (p *PermissionOverwrite) Type() OverwriteType {
	return p.overwriteType
}

// OverwriteType is the type for overwrite type constants.
type OverwriteType string

const (
	OverwriteTypeMember OverwriteType = "member"
	OverwriteTypeRole   OverwriteType = "role"
)

// PermissionBitSet is the type for a permission bit set.
type PermissionBitSet uint
