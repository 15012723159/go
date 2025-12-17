package main

import (
	"fmt"
	calc "forg/pkg/caculator"
)

func main() {
	fmt.Println("hello world")
	c := calc.Add2(1, 2)
	fmt.Println(c)
}
