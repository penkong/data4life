package apirouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/controllers/user"
)

func SetUpUser(r fiber.Router) {
	userRoutes := r.Group("/user")
	userRoutes.Post("/current", userCtrl.GetCurrent)
}
