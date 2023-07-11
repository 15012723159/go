package gee

import "net/http"

type Router struct {
	handles map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		handles: make(map[string]HandleFunc),
	}
}
func (router *Router) addRouter(method string, pattern string, handleFunc HandleFunc) {
	key := method + "-" + pattern
	router.handles[key] = handleFunc
}

func (router *Router) Handle(c *Context) {
	key := c.Req.Method + "-" + c.Req.URL.Path

	if handler, ok := router.handles[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 not found page:%s", c.Req.URL.Path)
	}

}
