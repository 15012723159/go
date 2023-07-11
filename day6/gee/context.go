package gee

import (
	json "encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter

	Req *http.Request

	Path string

	Method     string
	StatusCode int
	Params     map[string]string
	handlers   []HandleFunc
	index      int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {

	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}
func (context *Context) Next() {
	context.index++
	s := len(context.handlers)
	for ; context.index < s; context.index++ {
		context.handlers[context.index](context)
	}
}
func (context *Context) Param(key string) string {
	value, _ := context.Params[key]
	fmt.Println("key" + key + "value=" + value)
	return value
}

// 获取表单中参数的值

func (context *Context) PostForm(key string) interface{} {

	return context.Req.FormValue(key)
}

// 获取URL上面的参数
func (context *Context) Query(key string) interface{} {
	return context.Query(key)
}

//设置statusCode

func (context *Context) Status(code int) {
	context.StatusCode = code
	context.Writer.WriteHeader(code)
}

// 设置header头
func (context *Context) SetHeader(key, value string) {

	context.Writer.Header().Set(key, value)
}

// 设置响应的内容形势-1-返回string
func (context *Context) String(code int, format string, values ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 设置响应的内容形势-2-返回json

func (context *Context) Json(code int, obj interface{}) {
	context.SetHeader("Content-Type", "application/json")
	context.Status(code)
	encode := json.NewEncoder(context.Writer)

	if ok := encode.Encode(obj); ok != nil {
		http.Error(context.Writer, ok.Error(), 500)
	}
}

// 设置响应的内容形势-3-返回Data
func (context *Context) Data(code int, data []byte) {
	context.Status(code)
	context.Writer.Write(data)
}

// 设置响应的内容形势-4-html

func (context *Context) Html(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(code)
	context.Writer.Write([]byte(html))

}

func (context *Context) Fail(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(code)
	context.Writer.Write([]byte(html))
}
