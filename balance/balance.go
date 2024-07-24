// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// type User struct {
// 	ID      int    `json:"id"`
// 	Name    string `json:"name"`
// 	Surname string `json:"surname"`
// 	Balance int    `json:"balance"`
// }

// var db *sql.DB

// func TransferBalanceHandler(c *gin.Context) {
// 	firstUserID, _ := strconv.Atoi(c.PostForm("firstUserID"))
// 	secondUserID, _ := strconv.Atoi(c.PostForm("secondUserID"))
// 	summa, _ := strconv.Atoi(c.PostForm("summa"))
// 	perevod, _ := strconv.ParseBool(c.PostForm("perevod"))

// 	firstUser := GetUserFromDB(firstUserID)
// 	secondUser := GetUserFromDB(secondUserID)

// 	if perevod {
// 		if firstUser.Balance >= summa {
// 			firstUser.Balance -= summa
// 			secondUser.Balance += summa

// 			UpdateUserBalanceInDB(firstUser)
// 			UpdateUserBalanceInDB(secondUser)

// 			c.JSON(200, gin.H{"message": "Transfer successful"})
// 		} else {
// 			c.JSON(400, gin.H{"error": "Insufficient balance"})
// 		}
// 	} else {
// 		c.JSON(400, gin.H{"error": "Unsupported operation type"})
// 	}
// }

// func UpdateUserBalanceInDB(user *User) {
// 	_, err := db.Exec("UPDATE users SET balance = ? WHERE id = ?", user.Balance, user.ID)
// 	if err != nil {
// 		fmt.Println("Error updating user balance in the database:", err)
// 	}
// }