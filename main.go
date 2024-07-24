package main

import (
	database "api/db"
	"api/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	db, err := database.Connect()
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	routes.SetupRoutes(router, db)

	router.Run(":8080")

}
