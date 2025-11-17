package main

import "fmt"

func main() {
	switch s := 19; s {
	case 1:
		fmt.Println("hello world1")
	case 2:
		fmt.Println("hello world2")
	default:
		fmt.Println("hello world3")
	}

	switch month := 12; month {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		fmt.Println("春夏秋")
	case 10, 11, 12:
		fmt.Println("冬")
	default:
		fmt.Println("寒冬")
	}
	//fallthrough 穿透
	switch num := 10; num {
	default:
		fmt.Println("default")
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)

	}
	fmt.Println("=====================================")
	// break 中断 跳出循环 或者switch
	switch num := 10; num {
	default:
		fmt.Println("default")
		break
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}
}
