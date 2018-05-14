package main

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/GreyHood-Studio/web_server/router"
	"github.com/spf13/viper"
	"github.com/GreyHood-Studio/server_util/database"
	"github.com/GreyHood-Studio/server_util/error"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func setupConfig() (Configuration) {
	var config Configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return config
}

func setupLog() {

}

func setupPostgres(configs []DatabaseConfig) {
	var psql string
	for _, config := range configs {
		if config.Use != "game" && config.Use != "user" {
			//error
			return
		}
		psql = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Database)
		database.ConnectPG(config.Use, psql)
		fmt.Println(config.Use, "db Successfully connected!")
	}
}

func setupRouter(caches []CacheConfig) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// cache layer
	for _, config := range caches {
		if config.Use == "cache" {
			fmt.Println(config.Use, "redis Successfully connected!")
		} else if config.Use == "session" {
			createRedisSession (r, config.Address, config.Password)
			fmt.Println(config.Use, "redis Successfully connected!")
		}
	}

	// setup route 설정 전에 넘겨야함
	router.SetAPIRoute(r)
	return r
}

func createRedisSession(r *gin.Engine, address string, password string) {
	fmt.Println(address, password)
	store, err := redis.NewStore(10, "tcp", address, "", []byte(password))

	error.CheckError(err, "redis connect error")

	r.Use(sessions.Sessions("session", store))
}

func setupDatabase(dbs []DatabaseConfig){
	// rdb
	setupPostgres(dbs)
}