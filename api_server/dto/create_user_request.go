package dto

type CreateUserRequest struct {
	AccountID      int64  `json:"account_id"`
	Identification string `json:"identification"`
	Password       string `json:"password"`
	CompanyName    string `json:"company_name"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	EmergencyPhone string `json:"emergency_phone"`
	I18n           string `json:"i18n"`
	Address        string `json:"address"`
}
