package main

import (
	"go-myGram/database"
	"go-myGram/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
