package main

import (
	handlers "api/handlers"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func TestGetAllUsers(t *testing.T) {
	// Создаем фейковую базу данных для теста
	db, _ := sql.Open("mysql", "DB_USER:DB_PASS@:@tcp(DB_HOST:DB_PORT)/DB_NAME")
	defer db.Close()

	// Создаем экземпляр хендлера UserHandler
	userHandler := handlers.UserHandler{Db: db}

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
