package main

import (
	"log"

	//
	app "github.com/penkong/data4life"
	"github.com/penkong/data4life/util"
)

func main() {

	// Load up config files with viper
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	if err := app.Start(conf.ServerAddress); err != nil {
		log.Fatalf("app.Start: %v\n", err)
	}
}
