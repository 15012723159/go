package main

import "net/http"

func main() {

	http.Handle("/", http.StripPrefix("/file", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":9998", nil)
}
