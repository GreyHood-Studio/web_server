package main

import "fmt"

func main() {
	config := setupConfig()
	fmt.Println(config)

	setupDatabase(config.Database)
	r := setupRouter(config.Cache)

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprint(":", config.Server.Port))
}
