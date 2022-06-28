package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	app "github.com/penkong/data4life"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(app.Conf.TokenSecret),
		ErrorHandler: jwtError,
	})
}
