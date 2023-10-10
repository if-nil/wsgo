package server

import "net/http"

type Context struct {
	Method string
}

func NewContext(r *http.Request) *Context {
	c := &Context{
		Method: r.Method,
	}
	return c
}
