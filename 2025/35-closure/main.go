package main

import "fmt"

func main() {
	res := test1()
	fmt.Println(res())
	fmt.Println(res())
	fmt.Println(res())

	res2 := test1()
	fmt.Println(res2())
}

func test1() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}
