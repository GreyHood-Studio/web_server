package main

import (
	"fmt"
	"github.com/GreyHood-Studio/web_server/router"
	"github.com/GreyHood-Studio/server_util/setup"
)

func main() {
	config := setup.NewConfig()
	fmt.Println(config)

	setup.ConnectDatabase(config.Database)
	r := setup.ConnectRouter(config.Cache)

	router.SetAPIRoute(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprint(":", config.Server.Port))
}
