package main

import "fmt"

type Live interface {
	run()
	eat()
}
type Person struct {
	Name string
}

type Cat struct {
	Name string
}

func (person *Person) run() {
	fmt.Println(person.Name + "running")
}

func (person *Person) eat() {
	fmt.Println(person.Name + "eat")
}

func (cat Cat) run() {
	fmt.Println(cat.Name + "running")
}

func (cat *Cat) eat() {
	fmt.Println(cat.Name + "eat")
}

func allRun(live Live) {
	live.run()
}

func allEat(live Live) {
	live.eat()
}
func main() {

	//用户有结构体引用类型实现了interface的方法 那么就要用引用类型实现所有的方法 否则 实例调用就会报错 需要再初始化赋值的时候 带上&

	cat := &Cat{"胖胖"}
	allRun(cat)
	allEat(cat)

	person := Person{"瑜瑜"}
	allRun(&person)
	allEat(&person)
}
