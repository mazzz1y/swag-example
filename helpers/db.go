package helpers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		panic(err)
	}

	return db
}
