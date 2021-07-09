package main

import (
	"FirstGoApp/models"
	"fmt"
)

func main() {
	user := models.User{
		ID:        12,
		FirstName: "Roy",
		LastName:  "Rojas",
	}

	fmt.Println(user)
}
