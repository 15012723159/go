package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{1, 2, 3}
	copy(arr2, arr[3:6])

	fmt.Println(arr, arr2)

	newSlice := make([]int, 0)
	n := 5
	copy(newSlice, arr[:n])
	newSlice = append(newSlice, arr[n+1:]...)
	fmt.Println(newSlice)
}
