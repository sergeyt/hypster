package hypster

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AppBuilder provides fluent API to create RESTful web apps
type AppBuilder struct {
	impl     *mux.Router
	routes   map[string]*RouteBuilder
	services map[string]interface{}
}

// Handler is function to process HTTP requests
type Handler func(*Context)

// NewApp creates new instance of AppBuilder
func NewApp(services map[string]interface{}) *AppBuilder {
	routes := make(map[string]*RouteBuilder)
	return &AppBuilder{mux.NewRouter(), routes, services}
}

// ServeHTTP implements http.Handler
func (app *AppBuilder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.impl.ServeHTTP(w, req)
}

// GetServices lookups service with given key
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
			ctx := &Context{w, req, app}
			switch req.Method {
			case "GET":
				rb.get(ctx)
			case "POST":
				rb.post(ctx)
			case "PUT":
				rb.post(ctx)
			case "UPDATE":
				rb.update(ctx)
			case "PATCH":
				rb.patch(ctx)
			case "DELETE":
				rb.del(ctx)
			}
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
