package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	AccountId      int64
	UserId         int64     `gorm:"<-:create;primaryKey"`
	CreateDate     time.Time `gorm:"<-:create;autoCreateTime"`
	LastUpdated    time.Time
	RecordStatus   RecordStatus
	Active         RecordStatus
	Identification string
	Password       string `gorm:"<-:create;autoUpdateTime:milli"`
	CompanyName    string
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	EmergencyPhone string
	I18N           string `gorm:"column:i18n"`
	Address        string
	AuthMenu       string
	AuthKeys       string
	AuthGroups     string
}

func (u *User) GenerateUserData() string {
	data := map[string]interface{}{
		"first_name":      u.FirstName,
		"last_name":       u.LastName,
		"company_name":    u.CompanyName,
		"email":           u.Email,
		"phone":           u.Phone,
		"emergency_phone": u.EmergencyPhone,
		"i18n":            u.I18N,
		"address":         u.Address,
		"identification":  u.Identification,
	}
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
