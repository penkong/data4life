package apirouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/controllers/writer_to"
)

func SetupWriterTo(r fiber.Router) {
	userRoutes := r.Group("/write-to")
	userRoutes.Get("/token-pg", writerto.WriterToPG)
}
