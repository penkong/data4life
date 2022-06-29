package writerto

import "github.com/gofiber/fiber/v2"

func WriterToPG(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "I am writer to PG",
	})
}
