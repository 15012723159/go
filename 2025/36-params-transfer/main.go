package main

import "fmt"

func main() {
	a := 10
	b := 20
	arr := []int{2, 3}
	test(a, b, arr)
	fmt.Println(a, b)
	fmt.Println(arr)
}

func test(a, b int, arr []int) {
	a = 100
	b = 200
	arr[0] = 5
	arr[1] = 10
}
