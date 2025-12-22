package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Content string
}

type Pet struct {
	Name    string
	Age     int
	Content string
}

func (u User) Work() string {

	return u.Name + "工作内容是" + u.Content
}

type Teacher struct {
	user      User
	classroom string
}

func main() {
	teacher := Teacher{User{"王五", 20, "教书"}, "3年2班"}
	fmt.Println(teacher.user.Name, teacher.user.Age, teacher.classroom)

	fmt.Println(teacher.user.Work())
}
