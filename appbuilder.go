package hypster

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// AppBuilder provides fluent API to create RESTful web apps
type AppBuilder struct {
	impl     *mux.Router
	routes   map[string]*RouteBuilder
	services map[string]interface{}
}

// Handler is function to process HTTP requests
type Handler func(*Context) (interface{}, error)

// TODO add error type
type errorPayload struct {
	error string
}

// NewApp creates new instance of AppBuilder
func NewApp(services map[string]interface{}) *AppBuilder {
	routes := make(map[string]*RouteBuilder)
	return &AppBuilder{mux.NewRouter(), routes, services}
}

// ServeHTTP implements http.Handler
func (app *AppBuilder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.impl.ServeHTTP(w, req)
}

// GetService finds service from services map of the current web application.
func (app *AppBuilder) GetService(key string) interface{} {
	return app.services[key]
}

// Route returns RouteBuilder for given URL pattern
func (app *AppBuilder) Route(pattern string) *RouteBuilder {
	rb := app.routes[pattern]

	if rb == nil {
		rb := &RouteBuilder{app: app}
		app.routes[pattern] = rb

		app.impl.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
			vars := mux.Vars(req)
			ctx := &Context{w, req, vars, app}

			var res interface{}
			var err error
			var bytes []byte

			switch req.Method {
			case "GET":
				res, err = rb.get(ctx)
			case "POST":
				res, err = rb.post(ctx)
			case "PUT":
				res, err = rb.post(ctx)
			case "UPDATE":
				res, err = rb.update(ctx)
			case "PATCH":
				res, err = rb.patch(ctx)
			case "DELETE":
				res, err = rb.del(ctx)
			}

			if err != nil {
				bytes, _ = json.Marshal(errorPayload{err.Error()})
			} else {
				bytes, _ = json.Marshal(res)
			}

			w.Write(bytes)
		})
	}

	return rb
}

// Shortcuts

// Get registers GET handler
func (app *AppBuilder) Get(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Get(handler)
	return app
}

// Post register Post handler
func (app *AppBuilder) Post(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Post(handler)
	return app
}

// Put registers PUT handler
func (app *AppBuilder) Put(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Put(handler)
	return app
}

// Update registers UPDATE handler
func (app *AppBuilder) Update(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Update(handler)
	return app
}

// Patch registers PATCH handler
func (app *AppBuilder) Patch(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Patch(handler)
	return app
}

// Delete registers DELETE handler
func (app *AppBuilder) Delete(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Delete(handler)
	return app
}
