package hypster

import "net/http"

// Context defines context of HTTP request
type Context struct {
	// Response is http.ResponseWriter
	Response http.ResponseWriter
	// Request is instance of http.Request
	Request  *http.Request
	// Vars stores route variables
	Vars     map[string]string
	app      *AppBuilder
}

// Get returns service registered with given key
func (ctx *Context) Get(key string) interface{} {
	return ctx.app.GetService(key)
}
