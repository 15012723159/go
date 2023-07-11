package gee

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(context *Context) {
		t := time.Now()
		context.Next()
		log.Printf("[%d] %s in %v\n", context.StatusCode, context.Req.URL, time.Since(t))
	}
}
