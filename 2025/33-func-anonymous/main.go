package main

import "fmt"

func main() {
	func() {
		fmt.Println("无参数-无返回值匿名函数")
	}()

	func() string {
		fmt.Println("hello world")
		return "hello world"
	}()

	name := "霍去病"
	func(name string) {
		fmt.Println(name)
	}(name)

	name = func(name string) string {
		fmt.Println(name + "有返回值")
		return name + " hello world"
	}(name)

	fmt.Println(name)
}
