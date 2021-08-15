package jee

import (
	"fmt"
	"net/http"
)

const gap = "_"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engnie struct {
	router map[string]HandlerFunc
}

func New() *Engnie {
	return &Engnie{router: make(map[string]HandlerFunc)}
}

func (e *Engnie) addRouter(method, palette string, handler HandlerFunc) {
	key := method + gap + palette
	e.router[key] = handler
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
	key := req.Method + gap + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(rw, req)
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "404 NOT FOUND: %s\n", req.URL)
	}
}
