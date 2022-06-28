package apirouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/controllers/auth"
)

func SetUpAuth(r fiber.Router) {
	authRoutes := r.Group("/auth")
	authRoutes.Post("/login", authCtrl.Login)
	authRoutes.Post("/register", authCtrl.Register)
	authRoutes.Post("/forget-pass", authCtrl.ForgetPassword)
}
