package main

import (
	"api/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/aa")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	UserHandler := handlers.UserHandler{Db: db}

	router.GET("/user", UserHandler.GetAllUsers)
	router.GET("/user/:id", UserHandler.GetUserById)
	router.POST("/user", UserHandler.CreateUser)
	router.PUT("/updateuser/:id", UserHandler.UpdateUser)
	router.DELETE("/deleteuser/:id", UserHandler.DeleteUserById)
	router.Run(":8080")
}
