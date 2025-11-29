package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			goto hello
		}
		fmt.Println(i)
	}

hello:
	fmt.Println("hello world")
}
