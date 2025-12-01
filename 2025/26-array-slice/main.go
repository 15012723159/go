package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := arr[0:3] //包前不包后
	fmt.Println(arr, arr2)

	fmt.Printf("%p %p\n", &arr, arr2)

	arr2[2] = 10
	fmt.Println(arr, arr2)

	arr2 = append(arr2, []int{1, 2, 3}...)
	fmt.Println(arr, arr2)
	fmt.Printf("%p %p\n", &arr, arr2)
}
