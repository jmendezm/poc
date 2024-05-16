package middlewares

import (
	"fmt"
	"technical_test/api_server/dto"
	"technical_test/domain"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func ConnectionChecker() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Filter returns true
		connID := c.Get("connection_id")
		if connID == "" {
			return c.Status(403).JSON(&dto.ErrorResponse{ErrorMessage: "Unauthorized request"})
		}
		conn, err := services.GetAuthServiceInstance().GetConnectionByID(connID)
		if err != nil {
			ce := HandleError(err)
			return c.Status(ce.Code).JSON(&dto.ErrorResponse{ErrorMessage: ce.Message})
		}
		if conn == nil || conn.Connected == "" {
			return c.Status(403).JSON(&dto.ErrorResponse{ErrorMessage: domain.ErrNotLoggedIn.Message})
		}
		fmt.Println("connection found | allowed to pass")

		return c.Next()
	}
}
