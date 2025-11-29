package main

import "fmt"

func main() {

abc:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				break abc
			}
			fmt.Println(i, j)
		}
	}
}

// break 后接标签 直接跳转至标签
