package auth

import (
	"encoding/json"
	"strconv"
	"technical_test/api_server/dto"
	"technical_test/domain"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func UpdateUserInfo(ctx *fiber.Ctx) error {
	var (
		err       error
		userIDStr string
		userID    int64
		user      *domain.User
		req       *dto.UpdateUserRequest
	)
	userIDStr = ctx.Params("user_id", "")
	if userID, err = strconv.ParseInt(userIDStr, 10, 64); err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{ErrorMessage: "Bad user id"})
	}

	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return nil
	}
	user = &domain.User{
		UserId:         userID,
		AccountId:      req.AccountID,
		Identification: req.Identification,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		CompanyName:    req.CompanyName,
		Phone:          req.Phone,
		EmergencyPhone: req.EmergencyPhone,
		I18N:           req.I18n,
		Address:        req.Address,
	}
	user, err = services.GetAuthServiceInstance().UpdateUser(user)
	if err != nil {
		ce := HandleError(err)
		resp := dto.ErrorResponse{ErrorMessage: ce.Message}
		return ctx.Status(ce.Code).JSON(resp)
	}
	resp := dto.UserResponse{}
	resp.FromDomain(user)
	return ctx.JSON(resp)
}
