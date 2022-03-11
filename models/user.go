package models

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Say() {
	fmt.Println(u.Name)
}
