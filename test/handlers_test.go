package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	// Создаем фейковую базу данных для теста
	db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	defer db.Close()

	// Создаем экземпляр хендлера UserHandler
	userHandler := UserHandler{db: db}

	// Создаем тестовый HTTP запрос
	req, _ := http.NewRequest("GET", "/user/", nil)
	w := httptest.NewRecorder()

	// Создаем контекст Gin для обработки запроса
	r := gin.Default()
	r.GET("/user/", userHandler.GetAllUsers)
	r.ServeHTTP(w, req)

	// Проверяем статус код ответа
	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200")

	// Проверяем формат ответа на соответствие JSON
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Expected Content-Type application/json")
}

func TestMain(t *testing.T) {
	// Тестирование функции main, которая инициализирует роутер и запускает сервер
	go main()
}
