package model

type Bill struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

var Bills = []Bill{
	{ID: 1, Name: "Isa", Surname: "Isaev", Balance: 52},
	{ID: 2, Name: "Suleyma", Surname: "Suleymanov", Balance: 77},
	{ID: 3, Name: "Mag", Surname: "Magov", Balance: 14},
}

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Bill    uint16 `json:"bill"`
}
