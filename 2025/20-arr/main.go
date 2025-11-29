package main

import "fmt"

func main() {
	// 声名方式

	//1-完整写法
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// 2-短变量写法
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//方式3 长度大于初始值个数 只给前三个 其他的为默认值
	arr3 := [4]int{1, 2, 3}
	fmt.Println(arr3)

	//方式4 赋值时不写长度 数组长度根据元素个数确定
	arr4 := [...]int{1, 2, 3}
	fmt.Println(arr4)
}
