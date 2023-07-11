package gee

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	*Router
}

func (engine *Engine) addRouter(method string, pattern string, handleFunc HandleFunc) {
	engine.Router.addRouter(method, pattern, handleFunc)
}

func New() *Engine {
	return &Engine{Router: NewRouter()}
}

func (engine *Engine) GET(pattren string, handleFunc HandleFunc) {
	engine.addRouter("GET", pattren, handleFunc)
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
	context := NewContext(w, r)
	engine.Router.Handle(context)

}
func (engine *Engine) Run(ip string) error {

	return http.ListenAndServe(ip, engine)

}
