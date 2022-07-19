package main

import (
	"headfirstgo/crud/models"
	"headfirstgo/crud/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()

}
