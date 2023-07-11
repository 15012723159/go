package main

import "fmt"

func main() {

	/*fmt.Println("before panic")

	panic("hello world")
	defer func() {
		fmt.Println("hello world2")
	}()
	fmt.Println("after panic")*/
	test_recover()
	fmt.Println("after recover")
}

func test_recover() {
	defer func() {
		fmt.Println("before panic")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}

	fmt.Println(arr[4])

	fmt.Println("hello world")
}
