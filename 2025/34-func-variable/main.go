package main

import "fmt"

func main() {
	/*var a func()
	a = func() {
		fmt.Println("a")
	}
	a()*/
	/*mydo(func(name string) {
		fmt.Println(name)
	})*/

	result := a()
	result2 := result()
	fmt.Println(result2)
}

/*func mydo(arg func(name string)) {
	fmt.Println("执行了mydo")
	arg("testing")
}*/

func a() func() int {
	return func() int {
		return 100
	}

}
