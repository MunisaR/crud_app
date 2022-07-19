package main

import (
	"headfirstgo/crud/crud_app/models"
	"headfirstgo/crud/crud_app/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()

}
