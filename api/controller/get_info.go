package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	connectdb "github.com/penkong/data4life/pkg/connectDb"
)

func GetInfo(c *fiber.Ctx) error {
	fmt.Println(connectdb.Store)
	return c.SendString("Hello, World!")
}
