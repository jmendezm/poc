package auth

import (
	"encoding/json"
	"technical_test/api_server/dto"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	var (
		req    *dto.LoginRequest
		err    error
		connID string
	)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return nil
	}
	connID, err = services.GetAuthServiceInstance().Login(req.Email, req.Password)
	if err != nil {
		ce := HandleError(err)
		return ctx.Status(ce.Code).JSON(&dto.ErrorResponse{
			ErrorMessage: ce.Message,
		})
	}
	return ctx.Status(200).JSON(map[string]string{"connection_id": connID})
}
