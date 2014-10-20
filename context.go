package hypster

import "encoding/json"
import "net/http"

// Context defines context of HTTP request
type Context struct {
	// Response is http.ResponseWriter
	Response http.ResponseWriter
	// Request is instance of http.Request
	Request *http.Request
	// Vars stores route variables
	Vars map[string]string
	app  *AppBuilder
}

// TODO add error type
type errorPayload struct {
	error string
}

// GetService returns service registered with given key
func (ctx *Context) GetService(key string) interface{} {
	return ctx.app.GetService(key)
}

// WriteError writes error encoded as JSON {error: "message"}
func (ctx *Context) WriteError(err error) {
	bytes, _ := json.Marshal(errorPayload{err.Error()})
	ctx.Response.Write(bytes)
}

// Read decodes JSON input into given value
func (ctx *Context) Read(out interface{}) error {
	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(out)
	return err
}
