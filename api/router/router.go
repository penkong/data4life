package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/api/controller"
)

func Setup(app *fiber.App) {

	s := app.Group("/api/v1")
	s.Get("/", controller.GetInfo)
}
