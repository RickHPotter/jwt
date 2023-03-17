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

// * run with `compiledaemon --command="./jwt"` after [1]
func main() {
	routers.LoadRouters()
}

/* [1]
// !

cd go
cd src/github.com/RickHPotter
mkdir jwt
cd jwt

got mod init

go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v4
go get github.com/joho/godotenv
go get github.com/githubnemo/CompileDaemon

go install github.com/githubnemo/CompileDaemon

code .

// !
*/
