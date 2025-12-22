package main

import "fmt"

func main() {

	demo1()

}

func demo1() {
	fmt.Println("demo1的上半部分")
	defer func() {
		recover()
	}()
	demo2()
	fmt.Println("demo1的下半部分")
}
func demo2() {

	fmt.Println("demo2的上半部分")
	demo3()
	fmt.Println("demo2的下半部分")
}

func demo3() {

	fmt.Println("demo3的上半部分")
	panic("panic报错了")
	fmt.Println("demo3的下半部分")
}
