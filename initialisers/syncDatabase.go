package initialisers

import "github.com/RickHPotter/jwt/models"

func SyncDatabase() {
	// [1] + I guess the dst in DB.AutoMigrate(dst ...interface{})) stands for data structs
	DB.AutoMigrate(&(models.User{}))
}

// ! [1]
// ! AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
// ! It will change existing column's type if its size, precision, nullable changed.
// ! But it will not delete unused columns to protect your data.
