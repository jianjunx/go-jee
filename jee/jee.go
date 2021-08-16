package jee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engnie struct {
	router *router
}

func New() *Engnie {
	return &Engnie{router: newRouter()}
}

// 将路由添加到map中
func (e *Engnie) addRouter(method, palette string, handler HandlerFunc) {
	e.router.addRouter(method, palette, handler)
}

func (e *Engnie) GET(palette string, handler HandlerFunc) {
	e.addRouter("GET", palette, handler)
}

func (e *Engnie) POST(palette string, handler HandlerFunc) {
	e.addRouter("POST", palette, handler)
}

// 监听http服务的包装
func (e *Engnie) Run(addr string) error {
	// ListenAndServe第二个参数是Handler接口 这里Engnie已经实现了Handler
	return http.ListenAndServe(addr, e)
}

// Engnie结构体实现Handler接口ServeHTTP方法，当监听到请求时会自动调用该方法处理
func (e *Engnie) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// 创建上下文
	c := newContext(rw, req)
	// 调用路由处理函数
	e.router.handler(c)
}
