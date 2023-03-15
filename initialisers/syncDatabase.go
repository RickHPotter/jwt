package initialisers

import "github.com/RickHPotter/jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&(models.User{}))
}
