package main

import (
	"github.com/gin-gonic/gin"
	"github.com/GreyHood-Studio/web_server/router"
	"github.com/spf13/viper"
	"log"
	"fmt"
	"github.com/GreyHood-Studio/server_util/database"
)

func readConfigDB() {

}

func setupConfig() (Configuration, error) {
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

	return config, err
}

func setupLog() {

}

func setupRedis(configs []CacheConfig) {
	for _, config := range configs {
		if config.Use != "coordinate" && config.Use != "session" {
			return
		}
		database.ConnectRedis(config.Use, config.Session, config.Address, config.Password)
		fmt.Println(config.Use, "redis Successfully connected!")
	}
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

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	router.SetAPIRoute(r)
	return r
}

func setupConnection(dbs []DatabaseConfig, caches []CacheConfig){
	setupPostgres(dbs)
	setupRedis(caches)
}