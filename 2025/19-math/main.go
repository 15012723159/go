package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var a, b float64 = 12.3, 12.5

	fmt.Println(math.Floor(a))
	fmt.Println(math.Ceil(b))
	fmt.Println(math.Modf(a))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Round(a))
	fmt.Println(math.Round(b))
	fmt.Println("==============================")
	//随机数
	seed := int64(42)
	r1 := rand.New(rand.NewSource(seed))
	r2 := rand.New(rand.NewSource(seed))
	fmt.Println(r1.Intn(10), r1.Intn(10))
	fmt.Println(r2.Intn(10), r2.Intn(10))
}
