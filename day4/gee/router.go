package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots   map[string]*node
	handles map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		roots:   make(map[string]*node),
		handles: make(map[string]HandleFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break

			}
		}
	}
	return parts
}
func (router *Router) addRouter(method string, pattern string, handleFunc HandleFunc) {

	parts := parsePattern(pattern)
	key := method + "-" + pattern

	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parts, 0)

	router.handles[key] = handleFunc
}

func (router *Router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)

		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]

			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (router *Router) Handle(c *Context) {

	/*if handler, ok := router.handles[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 not found page:%s", c.Req.URL.Path)
	}*/
	n, params := router.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		router.handles[key](c)
	} else {
		c.String(http.StatusNotFound, "404 not found %s\n", c.Path)
	}

}
