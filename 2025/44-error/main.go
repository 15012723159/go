package main

import "fmt"

func main() {
	error := demo(1, 2)

	fmt.Println(error)
}

func demo(a, b int) error {

	return fmt.Errorf("a和b相加的结果是=%d", a+b)
}
