package model


type GameServer struct {
	// [ip]:[port]
	Address			string
	MaxRoom			int
	CurrentRoom		int
}

type NPCServer struct {
	Address			string
}

var GameServers map[string]GameServer
var NPCServers map[string]NPCServer
