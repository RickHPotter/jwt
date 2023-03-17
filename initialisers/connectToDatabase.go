package initialisers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// instead of *sql.DB usual pool, we use *gorm.DB
var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB") // data source name

	// gorm.Open() should be executed only once to avoid new pools
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to db.")
	}

}
