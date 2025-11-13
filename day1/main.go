package main

import "fmt"

var age int = 18

func main() {
	//fmt.Println("Hello, World!")
	// 输出结果
	//fmt.Printf("%x %d %o %b", 12, 18, 12, 18)
	// 获取转化的值
	//a := fmt.Sprintf("%X", 13)
	//fmt.Println("a=" + a)

	//fmt.Printf("%c", 65)

	//fmt.Printf("%s %q", "你好", "雷浩")
	//i := 18
	//fmt.Printf("%p", &i)
	//类型
	//fmt.Printf("%T", "false")
	//字符串%%
	//fmt.Printf("%d%%的青少年早恋", 19)
	//fmt.Printf("你好吗\n\t我很\t好")

	//长变量 可以在函数外 包内使用
	//var name string = "hello world"
	//fmt.Println(name)

	//短变量 只能在函数内部使用 用一个变量名 不能在同一区域出现
	name := "你好 世界"
	fmt.Println(name)
	fmt.Println(age)

	//5 一次声明多个变量
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
	//6 声明并赋值
	var d, e, f = 1, 2, false
	fmt.Println(d, e, f)
	//7 先声明后赋值，后短变量复制的方式 需要一个至少没有声明过的变量
	var (
		A = 1
		B = 2
		C = 3
	)
	A, B, C, D := 12, 13, 14, 15
	fmt.Println(A, B, C, D)

}
