package auth

import (
	"strconv"
	"technical_test/api_server/dto"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(ctx *fiber.Ctx) error {
	var (
		err    error
		userID int64
	)
	userID, err = strconv.ParseInt(ctx.Params("user_id"), 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(&dto.ErrorResponse{ErrorMessage: "Not valid user ID"})
	}
	err = services.GetAuthServiceInstance().DeleteUser(userID)
	if err != nil {
		ce := HandleError(err)
		return ctx.Status(ce.Code).JSON(&dto.ErrorResponse{
			ErrorMessage: ce.Message,
		})
	}
	return ctx.Status(200).JSON(map[string]string{"message": "User sucessfully deleted"})
}
