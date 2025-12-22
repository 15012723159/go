package main

import "fmt"

func main() {
	demo("323")
}

func demo(i interface{}) {
	result, err := i.(int)
	if !err {
		fmt.Println("传值不为int类型")
		return
	}
	fmt.Println(result)
}
