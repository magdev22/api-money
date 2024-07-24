package balance

import "fmt"

var pereveod bool
var summa int

type User struct {
	ID      int
	name    string
	surname string
	balance int
}

var Users = []User{
	{ID: 1, name: "Isa", surname: "Isaev", balance: 52},
	{ID: 2, name: "Suleyma", surname: "Suleymanov", balance: 77},
	{ID: 3, name: "Mag", surname: "Magov", balance: 14},
}

func BalanceOperation(firstuserId int, seconduserId int, pereveod bool, summa int) {
	var firstuser *User
	var seconduser *User
	for i := range Users {
		if Users[i].ID == firstuserId {
			firstuser = &Users[i]
		}
		if Users[i].ID == seconduserId {
			seconduser = &Users[i]
		}
	}
	if pereveod == true {
		firstuser.balance += summa
		seconduser.balance -= summa
	}
	fmt.Println(firstuser)
	fmt.Println(seconduser)
}

//сделать так, чтобы они переводили деньги друг другу, когда алекс переводит 500 у
// боба увеличивается счет на 500, а если боб переводит 500 у боба уменьшается счет на 500
