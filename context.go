package hypster

import "net/http"

// Context defines context of HTTP request
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	app      *AppBuilder
}

// Get returns service registered with given key
func (ctx *Context) Get(key string) interface{} {
	return ctx.app.GetService(key)
}
