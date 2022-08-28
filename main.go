package main

import (
	"testAPI/models"
	"testAPI/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Product{})

	r := routes.SetupRoutes(db)
	r.Run()
}
