package main

import (
	"fmt"
	"log"

	"github.com/market-inventory/server"
)

func main() {
	app, conf, db := server.Setup()

	defer func() {
		sqlDB, err := db.DB()

		if err != nil {
			panic("Error on app ending (0)")
		}

		if err := sqlDB.Close(); err != nil {
			panic("Error on app ending (1)")
		}
	}()

	listenArg := fmt.Sprintf(":%d", conf.API_PORT)

	log.Fatal(app.Listen(listenArg))
}
