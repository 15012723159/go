package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "18"

	r, _ := strconv.ParseInt(s, 16, 64)

	fmt.Println(r)
	fmt.Printf("%T", r)

	// 把int 转化为string
	/*	var i int = 22
		s2 := strconv.FormatInt(int64(i), 16)
		fmt.Println(s2)*/

	/*i := 22

	str := strconv.Itoa(i)
	fmt.Println(str)
	fmt.Printf("%T", str)*/
	/*str := "1.5"
	f, _ := strconv.ParseFloat(str, 64)
	fmt.Println(f)*/
	f := 1.23456
	s3 := strconv.FormatFloat(f, 'f', 4, 64)
	fmt.Println(s3)
}
