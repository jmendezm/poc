package auth

import (
	"encoding/json"
	"technical_test/api_server/dto"
	"technical_test/domain"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	var req *dto.CreateUserRequest
	var err error
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return nil
	}
	user := &domain.User{
		AccountId:      req.AccountID,
		Identification: req.Identification,
		Password:       req.Password,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		CompanyName:    req.CompanyName,
		Phone:          req.Phone,
		EmergencyPhone: req.EmergencyPhone,
		I18N:           req.I18n,
		Address:        req.Address,
	}
	user, err = services.GetAuthServiceInstance().CreateUser(user)
	if err != nil {
		ce := HandleError(err)
		resp := dto.ErrorResponse{ErrorMessage: ce.Message}
		return ctx.Status(ce.Code).JSON(resp)
	}
	resp := dto.UserResponse{
		AccountID:      user.AccountId,
		UserID:         user.UserId,
		Identification: user.Identification,
		CompanyName:    user.CompanyName,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Phone:          user.Phone,
		EmergencyPhone: user.EmergencyPhone,
		I18n:           user.I18N,
		Address:        user.Address,
	}
	return ctx.JSON(resp)
}
