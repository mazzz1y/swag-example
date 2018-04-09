package main

import (
	"github.com/dmirubtsov/swag-example/models"
)

func main() {
	models.MigrateDB()
	router := SetupRouter()
	router.Run(":8080")
}
