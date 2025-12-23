package main

import "fmt"

func main() {
	// map非线性安全
	// make 初始化
	/*var m map[string]string = make(map[string]string)
	fmt.Println(m == nil)
	fmt.Printf("%p\n", m)*/

	m := map[string]string{"中国": "北京", "美国": "华盛顿", "英国": "伦敦"}

	fmt.Println(m)

	m["中国"] = "南京"
	fmt.Println(m)

	m["德国"] = "柏林"
	fmt.Println(m)

	m["日本"] = "东京"
	fmt.Println(m)
	delete(m, "日本")
	fmt.Println(m)

	delete(m, "澳大利亚")
	fmt.Println(m)
	for key, value := range m {
		fmt.Println(key, value)
	}
	value, ok := m["中国re"]
	fmt.Println(value, ok)

	user := map[string]User{}
	user["user_1"] = User{"hello", 19, "user_1"}
	// 重置
	user["user_1"] = User{"hello", 19, "user_1_1"}

	fmt.Println(user)

}

type User struct {
	Name string
	Age  int
	Id   string
}
