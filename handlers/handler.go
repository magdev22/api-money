package handlers

import (
	model "api/data"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Db *sql.DB
}

func (h *UserHandler) CreateDatabaseUsers(c *gin.Context) {
	dbName := c.Param("db")
	if dbName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database name parameter is required"})
		return
	}

	_, err := h.Db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Database created successfully"})
}

func (h *UserHandler) CreateTableUsers(c *gin.Context) {
	tablename := c.Param("table")

	rows, err := h.Db.Query("CREATE TABLE " + tablename + " (id INT AUTO_INCREMENT PRIMARY KEY,name VARCHAR(20),surname VARCHAR(20),bill INT(20))")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating users table"})
		return
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Bill); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning users"})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, gin.H{"message": "table created successfully"})

}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	rows, err := h.Db.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting users"})
		return
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Bill); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning users"})
			return
		}
		users = append(users, user)
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling users"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(usersJSON))
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	row := h.Db.QueryRow("SELECT * FROM users WHERE id = ?", userID)
	var user model.User
	err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.Bill)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	result, err := h.Db.Exec("INSERT INTO users (name, surname, bill) VALUES (?, ?, ?)", user.Name, user.Surname, user.Bill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user_id": userID})
}

func (h *UserHandler) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	_, err = h.Db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateUser model.User
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	_, err = h.Db.Exec("UPDATE users SET name = ?, surname = ?, bill = ? WHERE id = ?", updateUser.Name, updateUser.Surname, updateUser.Bill, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *UserHandler) TransferBalanceHandler(c *gin.Context) {
	firstUserID, _ := strconv.Atoi(c.PostForm("firstUserID"))
	secondUserID, _ := strconv.Atoi(c.PostForm("secondUserID"))
	summa, _ := strconv.Atoi(c.PostForm("summa"))
	perevod, _ := strconv.ParseBool(c.PostForm("perevod"))

	var firstUser model.User
	err := h.Db.QueryRow("SELECT * FROM users WHERE id = ?", firstUserID).Scan(&firstUser.ID, &firstUser.Name, &firstUser.Surname, &firstUser.Bill)
	if err != nil {
		c.JSON(404, gin.H{"error": "First user not found"})
		return
	}

	var secondUser model.User
	err = h.Db.QueryRow("SELECT * FROM users WHERE id = ?", secondUserID).Scan(&secondUser.ID, &secondUser.Name, &secondUser.Surname, &secondUser.Bill)
	if err != nil {
		c.JSON(404, gin.H{"error": "Second user not found"})
		return
	}

	if perevod {
		if firstUser.Bill >= summa && firstUser.ID != secondUser.ID {
			firstUser.Bill -= summa
			secondUser.Bill += summa

			h.UpdateUserBalanceInDB(&firstUser)
			h.UpdateUserBalanceInDB(&secondUser)

			c.JSON(200, gin.H{"message": "Transfer successful, your balance: " + strconv.Itoa(firstUser.Bill)})
		} else {
			c.JSON(400, gin.H{"error": "error transfer"})
		}
	} else {
		c.JSON(400, gin.H{"error": "Unsupported operation type"})
	}
}

func (h *UserHandler) UpdateUserBalanceInDB(user *model.User) {
	_, err := h.Db.Exec("UPDATE users SET bill = ? WHERE id = ?", user.Bill, user.ID)
	if err != nil {
		fmt.Println("Error updating user balance in the database:", err)
	}
}
