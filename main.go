package main

import (
	"fmt"
	"gojee/jee"
	"net/http"
)

func main() {
	fmt.Println("12")
	r := jee.New()
	r.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hello")
	})

	r.Run(":3000")
}
