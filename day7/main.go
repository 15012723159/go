package main

import (
	"gee"
	"net/http"
)

func main() {

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

	r.GET("/panic", func(context *gee.Context) {
		names := []string{"helloworld"}
		context.String(http.StatusOK, names[100])
	})
	r.Use(gee.Recovery())
	r.Run(":9999")
}
