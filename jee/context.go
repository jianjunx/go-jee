package jee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// 上下文
type Context struct {
	Writer     http.ResponseWriter // 响应
	Req        *http.Request       // 请求
	Path       string              // 请求路径
	Method     string              // 请求方法
	StatusCode int                 // 状态码
}

func newContext(rw http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: rw,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 获取表单值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 获取Query的值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置响应头
func (c *Context) SetHader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// string类型返回的方法
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHader("Content-Type", "text/plan")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON类型返回的方法
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 数据返回的方法
func (c *Context) Data(code int, datas []byte) {
	c.Status(code)
	c.Writer.Write(datas)
}

// HTML类型返回值的方法
func (c *Context) HTML(code int, html string) {
	c.SetHader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
