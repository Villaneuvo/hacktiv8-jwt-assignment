package main

import (
	"jwt-hacktiv/database"
	"jwt-hacktiv/router"
)

func main() {
	database.StartDB()
	router := router.StartApp()
	router.Run(":8080")
}
