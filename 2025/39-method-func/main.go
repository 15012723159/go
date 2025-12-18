package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (user *User) Run() {
	user.Age -= 1
	fmt.Println("user run", user.Name, user.Age)
}
func main() {
	user := new(User)
	user.Name = "hello2"
	user.Age = 33
	user.Run()
	user.Run()
	user.Run()
}
