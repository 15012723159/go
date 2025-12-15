package main

import "fmt"

func main() {
	demo()
	demo2("hello world")
	name := demo3("hello world")

	fmt.Println(name)

	a, b, _ := demo4(1, 2, 3)
	fmt.Println(a)
	fmt.Println(b)

	demo5("霍去病", "李陵")
}

func demo() {
	fmt.Println("demo")
}

func demo2(name string) {
	fmt.Println("demo2" + name)
}

func demo3(name_param string) (name_return string) {
	name_return = name_param
	return
}

func demo4(a, b, c int) (c1, b1, a1 int) {

	return c, b, a
}

func demo5(names ...string) {
	for key, value := range names {
		fmt.Println(key, value)
	}
}
