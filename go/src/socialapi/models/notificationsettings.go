package models

import "time"

type NotificationSettings struct {
	// unique idetifier of the notification setting
	Id int64 `json:"id,string"`

	// Id of the channel
	ChannelId int64 `json:"channelId,string"       sql:"NOT NULL"`

	// Creator of the notification settting
	AccountId int64 `json:"accountId,string"               sql:"NOT NULL"`

	// Holds dektop setting type
	DesktopSetting string `json:"desktopSetting"	sql:"NOT NULL"`

	// Holds mobile setting type
	MobileSetting string `json:"mobileSetting"			sql:"NOT NULL"`

	// Holds the data if channel is muted or not
	IsMuted bool `json:"isMuted"`

	// Holds data that getting notification when @channel is written
	// If user doesn't want to get notification
	// when written to channel @channel, @here or name of the user.
	// User uses 'suppress' feature and doesn't get notification
	IsSuppressed bool `json:"isSuppressed"`

	// Creation date of the notification settings
	CreatedAt time.Time `json:"createdAt"          sql:"NOT NULL"`

	// Modification date of the notification settings
	UpdatedAt time.Time `json:"updatedAt"          sql:"NOT NULL"`
}
