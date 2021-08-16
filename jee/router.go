package jee

import "net/http"

const gap = "_"

type router struct {
	// 路由处理函数的集合，key是Method + _ + Path,
	handlers map[string]HandlerFunc
}

// 创建路由实例的包装
func newRouter() *router {
	return &router{map[string]HandlerFunc{}}
}

// 将路由添加到map，method 请求类型(GET/POST...),palette路由规则，handler 处理方法
func (r *router) addRouter(method, palette string, handler HandlerFunc) {
	key := method + gap + palette
	r.handlers[key] = handler
}

// 关键，根据Method和Path找到对应路由的处理方法
func (r *router) handler(c *Context) {
	key := c.Method + gap + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}
