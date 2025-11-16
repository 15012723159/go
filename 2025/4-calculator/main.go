package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	fmt.Println(a*b + c)

	c++

	fmt.Println(c)

	fmt.Println(!(c > b))
}
