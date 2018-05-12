package main

import "fmt"

func main() {
	config, err := setupConfig()
	fmt.Println("setup config: ",config, err)

	setupConnection(config.Database, config.Cache)
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprint(":", config.Server.Port))
}
