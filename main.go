package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/penkong/data4life/router"
	"github.com/penkong/data4life/util"
)

func main() {

	// Load up config files with viper
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := gorm.Open(postgres.Open(conf.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot load db: ", err)
	}

	fmt.Println(db)

	app := fiber.New(fiber.Config{
		AppName:       conf.APPName,
		ServerHeader:  "X-GO",
		StrictRouting: false,
		CaseSensitive: true,
		Immutable:     true,
		BodyLimit:     8 * 1024 * 1024,
	})

	approuter.Route(app)

	app.Listen(conf.ServerAddress)

}
