package main

import "fmt"

// 整型
// 1.注意取值范围
// uintptr 大小不确定
// 类型转化
// 数学运算必须保持同一类型
// 类型转化
func main() {
	var a int8 = 8
	var b int16 = 16
	var c int32 = 32
	var d int64 = 64
	e := 64

	fmt.Println(a, b, c, d)
	fmt.Printf("%T\n", e)

	// 类型转化 类型(值)

	fmt.Println(int16(a) + b)

	s := fmt.Sprintf("%b", a)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	f := 0x71D5
	fmt.Printf("%c\n", f)
}
