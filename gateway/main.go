package app

import (
	"log"
	"strings"
	//
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	//
	"github.com/penkong/data4life/gateway/pkg/connect_db"
	"github.com/penkong/data4life/gateway/router"
	"github.com/penkong/data4life/gateway/util"
)

var app *fiber.App
var Conf util.Config
var Conf_err error

// Init function , do bootstrap parts here .
func init() {

	// load env vars with viper as utility
	Conf, Conf_err = util.LoadConfig(".")
	if Conf_err != nil {
		log.Fatal("cannot load config:", Conf_err)
	}

	// database connections
	connectdb.Setup(&Conf)

	// make instance of fiber .
	app = fiber.New(fiber.Config{
		AppName:       Conf.APPName,
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
	apirouters.Setup(app)
}

// clean port and allow to listen on CMD
func Start(addr string) error {
	if strings.IndexByte(addr, ':') == -1 {
		addr = ":" + addr
	}

	return app.Listen(addr)
}
