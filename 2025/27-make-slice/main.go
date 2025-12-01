package main

import "fmt"

func main() {
	arr := make([]int, 0)
	fmt.Println(arr)
	fmt.Printf("%p %p\n", &arr, arr)
	fmt.Println(cap(arr))

	//删除某个数组某个元素
	arr2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr3 := arr2[0:7]
	arr3 = append(arr3, arr2[8:]...)

	fmt.Println(arr3)
	fmt.Printf("%p %p\n", &arr2, arr3)

}
