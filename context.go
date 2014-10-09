package hypster

import (
	"encoding/json"
	"net/http"
)

// Context defines context of HTTP request
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	app      *AppBuilder
}

type errorPayload struct {
	error string
}

// Get returns service registered with given key
func (ctx *Context) Get(key string) interface{} {
	return ctx.app.GetService(key)
}

// TODO ReadJson API

// WriteJson sends given value as JSON
func (ctx *Context) WriteJson(value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		b, _ = json.Marshal(errorPayload{err.Error()})
	}
	ctx.Response.Write(b)
}
