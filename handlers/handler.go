package handlers

import (
	model "api/data"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Db *sql.DB
}

var querryString = "CREATE TABLE users (id INT AUTO_INCREMENT PRIMARY KEY,name VARCHAR(20),surname VARCHAR(20),bill INT(20))"

func (h *UserHandler) CreateTableUsers(c *gin.Context) {
	rows, err := h.Db.Query(querryString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating database"})
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
