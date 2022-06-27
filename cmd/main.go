package main

import (
	"log"

	_ "github.com/lib/pq"

	//
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

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

	repo := connectdb.Setup(&conf)

	app := fiber.New(fiber.Config{
		AppName:       conf.APPName,
		ServerHeader:  "X-GO",
		StrictRouting: false,
		CaseSensitive: true,
		Immutable:     true,
		BodyLimit:     8 * 1024 * 1024,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	apirouters.Setup(app, repo)

	app.Listen(conf.ServerAddress)

}
