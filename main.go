package main

import (
	"fmt"
	models2 "github.com/devzzk/go-blog/app/Models"
)

func main() {
	user1 := models2.User{
		Age:  18,
		Name: "小红",
	}
	var users [10]models2.User
	users[0] = user1
	fmt.Println(users)
}
