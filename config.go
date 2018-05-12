package main

type Configuration struct {
	Server		ServerConfig
	Database 	[]DatabaseConfig
	Cache		[]CacheConfig
}

type ServerConfig struct {
	Port		int
	Debug		bool
}

type DatabaseConfig struct {
	Use			string
	Host		string
	Port		int
	User		string
	Password	string
	Database	string
}

type CacheConfig struct {
	Use			string
	Session		int
	Address		string
	Password	string
}