package auth

import (
	"strconv"
	"technical_test/api_server/dto"
	"technical_test/domain"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(ctx *fiber.Ctx) error {
	var (
		err       error
		userIDStr string
		userID    int64
		user      *domain.User
	)
	userIDStr = ctx.Params("user_id", "")
	if userID, err = strconv.ParseInt(userIDStr, 10, 64); err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{ErrorMessage: "Bad user id"})
	}
	user, err = services.GetAuthServiceInstance().GetUserByID(userID)
	if err != nil {
		ce := HandleError(err)
		resp := dto.ErrorResponse{ErrorMessage: ce.Message}
		return ctx.Status(ce.Code).JSON(resp)
	}
	resp := dto.UserResponse{}
	resp.FromDomain(user)
	return ctx.JSON(resp)
}
