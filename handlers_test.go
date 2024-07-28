package main

import (
	model "api/data"
	conn "api/database"
	handlers "api/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var db, err = conn.Connect()

func TestCreateUser(t *testing.T) {
	userHandler := handlers.UserHandler{Db: db}

	newUser := model.User{Name: "John", Surname: "Doe", Bill: 100}
	newUserJSON, _ := json.Marshal(newUser)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(newUserJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/user", userHandler.CreateUser)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200")

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User created successfully", response["message"], "Expected message 'User created successfully'")
}

func TestGetAllUsers(t *testing.T) {
	userHandler := handlers.UserHandler{Db: db}

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/users", userHandler.GetAllUsers)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200")

	// Добавьте необходимые проверки для формата ответа и т.д.
}

func TestGetUserById(t *testing.T) {
	userHandler := handlers.UserHandler{Db: db}

	req, _ := http.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/user/:id", userHandler.GetUserById)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200")

	// Добавьте необходимые проверки для формата ответа и т.д.
}

func TestDeleteUserById(t *testing.T) {
	userHandler := handlers.UserHandler{Db: db}

	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/user/:id", userHandler.DeleteUserById)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200")

	// Добавьте необходимые проверки для формата ответа и т.д.
}

func TestUpdateUser(t *testing.T) {
	userHandler := handlers.UserHandler{Db: db}

	updateUser := model.User{Name: "Updated Name", Surname: "Updated Surname", Bill: 200}
	updateUserJSON, _ := json.Marshal(updateUser)

	req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(updateUserJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r := gin.Default()
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200")

	// Добавьте необходимые проверки для формата ответа и т.д.
}

func TestHandlers(t *testing.T) {
	t.Run("TestCreateUser", TestCreateUser)
	t.Run("TestGetAllUsers", TestGetAllUsers)
	t.Run("TestGetUserById", TestGetUserById)
	t.Run("TestUpdateUser", TestUpdateUser)
	t.Run("TestDeleteUserById", TestDeleteUserById)
}
