package main

import (
	"fmt"

	"github.com/RickHPotter/jwt/initialisers"
	"github.com/RickHPotter/jwt/routers"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDatabase()
	initialisers.SyncDatabase()
}

func main() {
	routers.LoadRoute()
	fmt.Println("Hello Again.")
}
