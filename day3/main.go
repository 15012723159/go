package main

import (
	"gee"
)

func main() {
	r := gee.New()
	//r.GET("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	//})
	//
	//r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
	//	for k, v := range req.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//	}
	//})
	//fmt.Println("hello world")

	r.GET("/", func(context *gee.Context) {
		context.String(200, "hello world")
	})
	r.GET("/json", func(context *gee.Context) {
		context.Json(200, map[string]string{
			"name": "qfx",
			"age":  "18",
		})
	})

	r.GET("/hello/:name", func(context *gee.Context) {

		context.Json(200, map[string]string{
			"name": context.Param("name"),
			"age":  "18",
		})
	})
	r.Run(":9999")
}
