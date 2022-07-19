package main

import (
	"headfirstgo/crud_app/models"
	"headfirstgo/crud_app/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()

}
