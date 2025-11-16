package main

import "fmt"

func main() {

	var a bool = true
	var b bool = false
	c := true
	fmt.Println(a, b, c)

	d := 5 < 3

	fmt.Println(d)
	fmt.Printf("%T\n", d)
}
