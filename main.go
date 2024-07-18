package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Bill    int    `json:"bill"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/aa")
		if err != nil {
			c.JSON(500, gin.H{"error": "Database connection error"})
			return
		}
		defer db.Close()

		userJSON, err := GetAllUsers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error getting users"})
			return
		}

		c.Header("Content-Type", "application/json")
		c.String(200, string(userJSON))
	})

	r.Run(":8080")
}

func GetAllUsers(db *sql.DB) ([]byte, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Bill); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return usersJSON, nil
}

func updateUserById(db *sql.DB, name string, surname string, bill int, id int) error {
	result, err := db.Exec("UPDATE users SET name = ?, surname = ?, bill = ? WHERE id = ?", name, surname, bill, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("User with id %d updated. Rows affected: %d\n", id, rowsAffected)
	return nil
}

func deleteUserById(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("User with id %d deleted. Rows affected: %d\n", id, rowsAffected)
	return nil
}
