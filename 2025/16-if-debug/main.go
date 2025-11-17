package main

import "fmt"

func main() {
	var score int = 65

	if score > 65 {
		fmt.Println("大于等于65")
	}

	if score <= 65 {
		fmt.Println("小于等于65")
	}
	if point := 65; point > 65 {
		fmt.Println("大于65")
	} else if point < 65 {
		fmt.Println("小于65")
	} else {
		fmt.Println("等于65")
	}
}
