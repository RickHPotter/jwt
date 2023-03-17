package main

import (
	"github.com/RickHPotter/jwt/initialisers"
	"github.com/RickHPotter/jwt/routers"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDatabase()
	initialisers.SyncDatabase()
}

// * run with `compiledaemon --command="./jwt"`
func main() {
	routers.LoadRoute()
}
