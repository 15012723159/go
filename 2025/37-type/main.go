package main

import "fmt"

// 结构体名称首字母大写 外面包可以访问
// 结构体属性 首字母大写 可以跨包访问
type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var user User

	user.Id = 1
	user.Age = 2
	user.Name = "hello world"

	fmt.Println(user)

	user = User{2, "fs", 18}

	fmt.Println(user)

	user = User{
		Age:  19,
		Name: "test",
	}

	fmt.Println(user)

}
