package querries

// import (
// 	"database/sql"
// 	"fmt"
// )

// var user model.User
//     var db
// 	db := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/bills")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// func insertUser(dbConn) {
// 	insert, err := dbConn.Query("INSERT INTO `users`(`name`,`surname`,`bill`) VALUES('alex','magdev',52)")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer insert.Close()
// }

// func selectUsers(db, user) {
// 	res, err := db.Query("SELECT `name`, `surname,``bill` FROM `users`")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for res.Next() {
// 		err = res.Scan(&user.Name, &user.Surname, &user.Bill)
// 		if err != nil {
// 			panic(err)

// 		}
// 		fmt.Println(fmt.Sprintf("User:%s with bill %d", user.Surname, user.Bill))
// 	}
// }
