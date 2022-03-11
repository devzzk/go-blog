package main

import (
	"fmt"
	"github.com/devzzk/go-blog/models"
)

func main() {
	user1 := models.User{
		Age:  18,
		Name: "小红",
	}
	var users [10]models.User
	users[0] = user1
	fmt.Println(users)
}
