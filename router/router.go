package approuter

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
