package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/penkong/data4life/util"
)

// Protected protect routes
func Protected() fiber.Handler {

	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(conf.TokenSecret),
		ErrorHandler: jwtError,
	})
}
