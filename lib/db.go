package lib

import (
	"errors"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

type User struct {
	gorm.Model
	Name string
	Age  int
}

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	database = db
}

func GetDatabase() (*gorm.DB, error) {
	if database == nil {
		return nil, errors.New("database is not initialized")
	}
	return database, nil
}
