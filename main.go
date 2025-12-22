package main

import (
	"fmt"
	"log"

	"github.com/market-inventory/server"
)

func main() {
	app, conf := server.Setup()

	defer func() {
		log.Println("see you later...")
	}()

	listenArg := fmt.Sprintf(":%d", conf.API_PORT)

	log.Fatal(app.Listen(listenArg))
}
