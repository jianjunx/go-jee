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

func (e *Engnie) addRouter(method, palette string, handler HandlerFunc) {
	e.router.addRouter(method, palette, handler)
}

func (e *Engnie) GET(palette string, handler HandlerFunc) {
	e.addRouter("GET", palette, handler)
}

func (e *Engnie) POST(palette string, handler HandlerFunc) {
	e.addRouter("POST", palette, handler)
}

func (e *Engnie) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engnie) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := newContext(rw, req)
	e.router.handler(c)
}
