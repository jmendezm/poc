package dto

import (
	"technical_test/domain"

	"github.com/nleeper/goment"
)

type UserResponse struct {
	AccountID      int64  `json:"account_id"`
	UserID         int64  `json:"user_id"`
	Identification string `json:"identification"`
	CompanyName    string `json:"company_name"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	EmergencyPhone string `json:"emergency_phone"`
	I18n           string `json:"i18n"`
	Address        string `json:"address"`
	CreatedDate    string `json:"create_date"`
	LastUpdateDate string `json:"last_update_date"`
}

func (ur *UserResponse) FromDomain(user *domain.User) {
	ur.AccountID = user.AccountId
	ur.UserID = user.UserId
	ur.Identification = user.Identification
	ur.CompanyName = user.CompanyName
	ur.FirstName = user.FirstName
	ur.LastName = user.LastName
	ur.Email = user.Email
	ur.Phone = user.Phone
	ur.EmergencyPhone = user.EmergencyPhone
	ur.I18n = user.I18N
	ur.Address = user.Address
	t, _ := goment.New(user.CreateDate)
	ur.CreatedDate = t.Format("L LTS Z")
	t, _ = goment.New(user.LastUpdated)
	ur.LastUpdateDate = t.Format("L LTS Z")
}
