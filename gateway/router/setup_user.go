package apirouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/controllers/user"
	"github.com/penkong/data4life/gateway/middleware"
)

func SetUpUser(r fiber.Router) {
	userRoutes := r.Group("/user")
	userRoutes.Post("/current", userCtrl.GetCurrent)
	userRoutes.Get("/:id", userCtrl.GetUser)
	userRoutes.Patch("/:id", middleware.Protected(), userCtrl.UpdateUser)
	userRoutes.Delete("/:id", middleware.Protected(), userCtrl.DeleteUser)
}
