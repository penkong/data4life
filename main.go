package main

import (
	_ "github.com/lib/pq"
	"log"
	//
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	//
	"github.com/penkong/data4life/api/router"
	"github.com/penkong/data4life/pkg/connectDb"
	"github.com/penkong/data4life/util"
)

func main() {

	// Load up config files with viper
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connectdb.Setup(&conf)

	app := fiber.New(fiber.Config{
		AppName:       conf.APPName,
		ServerHeader:  "X-GO",
		StrictRouting: false,
		CaseSensitive: true,
		Immutable:     true,
		BodyLimit:     8 * 1024 * 1024,
	})

	app.Use(cors.New())
	app.Use(logger.New())
	api.Setup(app)

	app.Listen(conf.ServerAddress)

}
