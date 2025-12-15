package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)
	r.Value = 0
	r.Next().Value = 1
	r.Next().Next().Value = 2
	r.Prev().Value = 4
	r.Prev().Prev().Value = 3
	r.Unlink(9)
	/*r1 := ring.New(1)
	r1.Value = 5
	r.Prev().Link(r1)*/
	r.Do(func(x interface{}) {
		fmt.Printf("%d ", x)
		fmt.Println()
	})
}
