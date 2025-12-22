package main

import "fmt"

func main() {

	//defer 采用栈结构 先进后出
	fmt.Println("开始执行操作")
	defer func() {
		fmt.Println("结束了所有的程序1")
	}()
	defer func() {
		fmt.Println("结束了所有的程序2")
	}()
	defer func() {
		fmt.Println("结束了所有的程序3")
	}()
	fmt.Println("执行结束")

	fmt.Println(demo())

	fmt.Println(demo2())
}

// defer 和return同时存在时 先return 赋值后 再执行defer 后再跳出函数

func demo() int {
	i := 0

	defer func() {
		i = i + 2
	}()
	return i
}

func demo2() (z int) {
	i := 0
	z = 0
	defer func() {
		z = i + 2
	}()

	return z
}
