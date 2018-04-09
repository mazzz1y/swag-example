package models

import (
	"github.com/dmirubtsov/swag-example/helpers"
)

var DB = helpers.DBConnect()

func MigrateDB() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Wallet{})
	DB.AutoMigrate(&Family{})
	DB.AutoMigrate(&UserFamilies{})

	DB.Model(&Wallet{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Wallet{}).AddForeignKey("family_id", "families(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Wallet{}).AddForeignKey("currency_id", "currencies(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Family{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.Model(&UserFamilies{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.Model(&UserFamilies{}).AddForeignKey("family_id", "families(id)", "RESTRICT", "RESTRICT")
}
