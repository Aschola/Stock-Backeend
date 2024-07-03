package main

import (
	"project/db"
	"project/routes"
	"project/utils"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Set utils.DB to the initialized database instance
	utils.DB = db.GetDB()

	// Setup the router
	r := routes.SetupRouter()
	r.Run(":8080")
}
