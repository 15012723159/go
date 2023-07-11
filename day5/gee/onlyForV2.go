package gee

import (
	"log"
	"time"
)

func OnlyForV2() HandleFunc {
	return func(context *Context) {
		t := time.Now()
		context.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s int %v for v2", context.StatusCode, context.Req.URL, time.Since(t))
	}
}
