## Jee Web Framework

Jee 是一个 Golang 的 web 框架

## 快速开始

```go
package main

import (
	"fmt"
	"gojee/jee"
	"net/http"
)

func main() {
	fmt.Println("12")
	r := jee.New()
	r.GET("/", func(c *jee.Context) {
		name := c.Query("name")
		c.HTML(http.StatusOK, "<h1>"+name+"</h1>")
	})
	r.POST("/hello", func(c *jee.Context) {
		c.JSON(http.StatusOK, jee.H{
			"username": c.PostForm("username"),
			"passwd":   c.PostForm("passwd"),
		})
	})
	r.Run(":3000")
}
```
