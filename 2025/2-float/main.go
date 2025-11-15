package main

import "fmt"

func main() {
	// 只有float64和float32

	/*	var i int = 10
		var f float64 = 2

		f = float64(i)
		fmt.Println(f, i)
		fmt.Printf("%f\n", f)
		fmt.Printf("%T\n", f)*/

	var a float32 = 1.6

	var b float64 = 1.5

	fmt.Println(float64(a) + b)
	fmt.Println(float32(b) + a)

}
