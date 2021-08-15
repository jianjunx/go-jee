package jee

import "net/http"

const gap = "_"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{map[string]HandlerFunc{}}
}

func (r *router) addRouter(method, palette string, handler HandlerFunc) {
	key := method + gap + palette
	r.handlers[key] = handler
}

func (r *router) handler(c *Context) {
	key := c.Method + gap + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}
