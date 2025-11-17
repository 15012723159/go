package main

import "fmt"

func main() {
	var a string = "hello"

	b := &a

	*b = "hello world"
	fmt.Println(a, *b)

	c := *&a

	fmt.Println(c)
}
