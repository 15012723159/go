package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := new(User)
	fmt.Println(user)
	fmt.Println(user == nil)
	fmt.Printf("%p\n", &user)
	user.Name = "hello"
	user.Age = 20

	user2 := user

	user2.Name = "hello2"
	fmt.Println(user)
	fmt.Println(user2)

	fmt.Println(user == user2)

	user3 := User{"hello2", 20}

	fmt.Println(user3 == *user2)
}
