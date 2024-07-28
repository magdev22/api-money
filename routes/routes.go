package routes

import (
	"api/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	userHandler := handlers.UserHandler{Db: db}
	//creating database and table
	router.POST("/database/:db", userHandler.CreateDatabaseUsers)
	router.POST("/table/:table", userHandler.CreateTableUsers)
	//user routes
	router.GET("/user", userHandler.GetAllUsers)
	router.GET("/user/:id", userHandler.GetUserById)
	router.POST("/user", userHandler.CreateUser)
	router.PUT("/updateuser/:id", userHandler.UpdateUser)
	router.DELETE("/deleteuser/:id", userHandler.DeleteUserById)
	router.POST("/transfer", userHandler.TransferBalanceHandler)
}
