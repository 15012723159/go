package main

import "fmt"

func main() {
	//声明slice 和数组
	var slice []string
	var array [3]string

	fmt.Println(slice, array)

	fmt.Println(slice == nil)

	fmt.Printf("%p", slice)

	// 切片是引用类型
	names := []string{"hello", "world"}
	fmt.Println(names)
	names2 := names

	fmt.Println(names, names2)

	fmt.Printf("%p %p", names, names2)
}
