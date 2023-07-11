package gee

import (
	"net/http"
	"strings"
)

type HandleFunc func(*Context)

type Engine struct {
	*Router
	*RouterGroup
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc
	parent      *RouterGroup
	engine      *Engine //全局共享一个engine
}

func (engine *Engine) addRouter(method string, pattern string, handleFunc HandleFunc) {
	engine.Router.addRouter(method, pattern, handleFunc)
}

func (gorup *RouterGroup) addRouter(method string, comp string, handleFunc HandleFunc) {
	pattern := gorup.prefix + comp
	gorup.engine.Router.addRouter(method, pattern, handleFunc)
}
func New() *Engine {
	//return &Engine{Router: NewRouter()}
	engine := &Engine{Router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
func (engine *Engine) GET(pattren string, handleFunc HandleFunc) {
	engine.addRouter("GET", pattren, handleFunc)
}

func (gorup *RouterGroup) Get(pattern string, handleFunc HandleFunc) {
	gorup.addRouter("GET", pattern, handleFunc)
}

func (gorup *RouterGroup) Post(pattern string, handleFunc HandleFunc) {
	gorup.addRouter("POST", pattern, handleFunc)
}

func (engine *Engine) POST(pattren string, handleFunc HandleFunc) {
	engine.addRouter("POST", pattren, handleFunc)
}

func (engine *Engine) DELETE(pattren string, handleFunc HandleFunc) {
	engine.addRouter("DELETE", pattren, handleFunc)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.route[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "%s", "没有发现服务")
	}*/
	var middlewares []HandleFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context := NewContext(w, r)
	context.handlers = middlewares
	engine.Router.Handle(context)

}
func (engine *Engine) Run(ip string) error {

	return http.ListenAndServe(ip, engine)

}

/**
Use MiddleWare //使用中间键
*/

func (group *RouterGroup) Use(mideleWare ...HandleFunc) {
	group.middlewares = append(group.middlewares, mideleWare...)
}
