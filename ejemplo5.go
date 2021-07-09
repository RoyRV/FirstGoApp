package main

import (
	"fmt"
)

func main() {
	//declarar una clase
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var _user user
	_user.ID = 13
	_user.FirstName = "Roy"
	_user.LastName = "Rojas"
	fmt.Println(_user)
	fmt.Println(_user.LastName)
	//crear un objeto e inicializarlo

	_user2 := user{
		ID:        3,
		LastName:  "Rojas",
		FirstName: "Sammy",
	}
	fmt.Println(_user2)
}
