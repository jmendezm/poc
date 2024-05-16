package domain

import "time"

type Connection struct {
	ID           int64  `gorm:"<-:create;primaryKey"`
	ConnectionID string `gorm:"<-:create;index"`
	AccountId    int64
	UserId       int64     `gorm:"<-:create"`
	CreateDate   time.Time `gorm:"<-:create;autoCreateTime:milli"`
	LastUpdated  time.Time `gorm:"autoUpdateTime:milli"`
	RecordStatus RecordStatus
	Active       RecordStatus
	Connected    string
	Disconnected string
	UserData     string
	AccountData  string
	AuthMenu     string
	AuthKeys     string
	AuthGroups   string
}
