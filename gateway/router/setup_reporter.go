package apirouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/controllers/reporter"
)

func SetupReporter(r fiber.Router) {
	reporterRoutes := r.Group("/reporter")
	reporterRoutes.Get("/token-pg", reporter.ReporterPG)
}
