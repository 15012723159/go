package main

import (
	"fmt"
	"gee"
)

func main() {

	var a = 1
	var b = 2
	defer func() {
		fmt.Println(a + b)
	}()
	fmt.Println("main")

	r := gee.New()
	/*pathjoin := path.Join("hello", "world")
	fmt.Println("pathjoin=" + pathjoin)*/
	r.Static("/assets", "D:/goproject/7day/day6/static")

	r.GET("/hello/*name", func(context *gee.Context) {

		context.Json(200, map[string]string{
			"name": context.Param("name"),
			"age":  "18",
		})
	})

	r.Run(":9999")
}
