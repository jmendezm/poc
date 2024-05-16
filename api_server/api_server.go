package api_server

import (
	"fmt"
	"log"
	authHandlers "technical_test/api_server/handlers/auth"
	sitesHandlers "technical_test/api_server/handlers/sites"
	"technical_test/api_server/middlewares"
	conf "technical_test/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() {
	config := conf.Get()
	app := fiber.New()
	app.Use(recover.New())

	authGroup := app.Group("/auth")
	authGroup.Post("/", authHandlers.CreateUser)
	authGroup.Post("/login", authHandlers.Login)
	authGroup.Get("/login", authHandlers.GetLogin)
	authGroup.Get("/:user_id", authHandlers.GetUserInfo)
	authGroup.Put("/:user_id", authHandlers.UpdateUserInfo)
	authGroup.Delete("/:user_id", authHandlers.DeleteUser)

	sitesGroup := app.Group("/sites")
	sitesGroup.Use(middlewares.ConnectionChecker())
	sitesGroup.Get("/", sitesHandlers.GetSites)
	sitesGroup.Post("/", sitesHandlers.CreateSite)
	sitesGroup.Get("/:site_id", sitesHandlers.GetSiteByID)
	sitesGroup.Delete("/:site_id", sitesHandlers.DeleteSite)
	sitesGroup.Put("/:site_id", sitesHandlers.UpdateSite)

	if err := app.Listen(fmt.Sprintf(":%d", config.ApiListenPort)); err != nil {
		log.Fatal(err.Error())
	}
}
