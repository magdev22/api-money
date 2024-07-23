package database

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

const (
	DB_USER = "DB_USER"
	DB_NAME = "DB_NAME"
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
)

func Connect() (*sql.DB, error) {
	ConnectString := "root:@tcp(127.0.0.1:3306)/aa"
	fmt.Print(ConnectString)
	db, err := sql.Open("mysql", ConnectString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
