package handlers

// //product

// import (
// 	model "api/data"
// 	"api/querries"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// var querri querries.Querries()

// func GetBills(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, )
// }

// func PostBills(c *gin.Context) {
// 	var NewBills model.Bill
// 	if err := c.BindJSON(&NewBills); err != nil {
// 		return
// 	}
// 	model.Bills = append(model.Bills, NewBills)
// 	c.IndentedJSON(http.StatusCreated, NewBills)
// }
// func GetBillBySurname(c *gin.Context) {
// 	surname := c.Param("surname")
// 	for _, a := range model.Bills {
// 		if a.Surname == surname {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
// }
// func DeleteBillBySurname(c *gin.Context) {
// 	surname := c.Param("surname")
// 	for i, a := range model.Bills {
// 		if a.Surname == surname {
// 			model.Bills = append(model.Bills[:i], model.Bills[i+1:]...)
// 			c.IndentedJSON(http.StatusNoContent, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
// }
// func UpdateBillBySurname(c *gin.Context) {
// 	var NewBill model.Bill
// 	surname := c.Param("surname")
// 	for i, a := range model.Bills {
// 		if a.Surname == surname {
// 			c.BindJSON(&NewBill)
// 			model.Bills[i] = NewBill
// 			c.IndentedJSON(http.StatusOK, NewBill)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
// }
