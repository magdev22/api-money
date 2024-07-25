package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func Connect() (*sql.DB, error) {
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	ConnectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, "", DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("mysql", ConnectString)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
