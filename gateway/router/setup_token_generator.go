package apirouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/controllers/token_generator"
)

func SetupTokenGenerator(r fiber.Router) {
	tokenGenerator := r.Group("/token")
	tokenGenerator.Get("/generator", tokengenerator.TokenCreator)
}
