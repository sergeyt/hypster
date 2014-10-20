package hypster

import "net/http"
import "github.com/gorilla/mux"

// AppBuilder provides fluent API to create RESTful web apps
type AppBuilder struct {
	router   *mux.Router
	services map[string]interface{}
}

// Handler is function to process HTTP requests
type Handler func(*Context) (interface{}, error)

// NewApp creates new instance of AppBuilder
func NewApp(services map[string]interface{}) *AppBuilder {
	return &AppBuilder{
		router:   mux.NewRouter(),
		services: services,
	}
}

// ServeHTTP implements http.Handler
func (app *AppBuilder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.router.ServeHTTP(w, req)
}

// GetService finds service from services map of the current web application.
func (app *AppBuilder) GetService(key string) interface{} {
	return app.services[key]
}

// Route returns RouteBuilder for given URL pattern
func (app *AppBuilder) Route(pattern string) *RouteBuilder {
	return &RouteBuilder{app, pattern}
}

// Shortcuts

// Head registers HEAD handler
func (app *AppBuilder) Head(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Head(handler)
	return app
}

// Options registers OPTIONS handler
func (app *AppBuilder) Options(pattern string, handler Handler) *AppBuilder {
	app.Route(pattern).Options(handler)
	return app
}

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
