package tokengenerator

import "github.com/gofiber/fiber/v2"

func TokenCreator(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "i am token creator",
	})
}
