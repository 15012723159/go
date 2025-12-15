package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 初始化
	// 取头
	// 取尾
	mylist := list.New()

	//fmt.Println(mylist.Len())
	//fmt.Printf("%p", mylist)

	mylist.PushBack("a")
	mylist.PushFront("b")
	mylist.InsertAfter("c", mylist.Back())
	mylist.InsertBefore("d", mylist.Front())
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	n := 3
	var curr *list.Element
	if n == 1 {
		curr = mylist.Front()
	} else if n == mylist.Len() {
		curr = mylist.Back()
	} else {
		curr = mylist.Front()
		for i := 0; i < n-1; i++ {
			curr = curr.Next()
		}
	}

	fmt.Println(curr.Value)
}
