package app

import (
	"log"
	"strings"
	//
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	//
	"github.com/penkong/data4life/api/router"
	"github.com/penkong/data4life/pkg/connect_db"
	"github.com/penkong/data4life/util"
)

var app *fiber.App

// Init function , do bootstrap parts here .
func init() {

	// load env vars with viper as utility
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// database connections
	repo := connectdb.Setup(&conf)

	// make instance of fiber .
	app = fiber.New(fiber.Config{
		AppName:       conf.APPName,
		ServerHeader:  "X-GO",
		StrictRouting: false,
		CaseSensitive: true,
		Immutable:     true,
		BodyLimit:     8 * 1024 * 1024,
	})

	// middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// routers
	apirouters.Setup(app, repo)
}

// clean port and allow to listen on CMD
func Start(addr string) error {
	if strings.IndexByte(addr, ':') == -1 {
		addr = ":" + addr
	}

	return app.Listen(addr)
}
