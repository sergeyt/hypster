package hypster

import "net/http"
import "encoding/json"
import "github.com/gorilla/mux"

// RouteBuilder holds route handlers
type RouteBuilder struct {
	app     *AppBuilder
	pattern string // current URL pattern
}

// Head registers HEAD handler
func (route *RouteBuilder) Head(handler interface{}) *RouteBuilder {
	return route.register("HEAD", handler)
}

// Options registers OPTIONS handler
func (route *RouteBuilder) Options(handler interface{}) *RouteBuilder {
	return route.register("OPTIONS", handler)
}

// Get registers GET handler
func (route *RouteBuilder) Get(handler interface{}) *RouteBuilder {
	return route.register("GET", handler)
}

// Post registers POST handler
func (route *RouteBuilder) Post(handler interface{}) *RouteBuilder {
	return route.register("POST", handler)
}

// Put registers PUT handler
func (route *RouteBuilder) Put(handler interface{}) *RouteBuilder {
	return route.register("PUT", handler)
}

// Update registers UPDATE handler
func (route *RouteBuilder) Update(handler interface{}) *RouteBuilder {
	return route.register("UPDATE", handler)
}

// Patch registers PATCH handler
func (route *RouteBuilder) Patch(handler interface{}) *RouteBuilder {
	return route.register("PATCH", handler)
}

// Delete registers DELETE handler
func (route *RouteBuilder) Delete(handler interface{}) *RouteBuilder {
	return route.register("DELETE", handler)
}

func (route *RouteBuilder) register(verb string, handler interface{}) *RouteBuilder {
	route.app.router.HandleFunc(route.pattern, route.wrapHandler(handler)).Methods(verb)
	return route
}

func (route *RouteBuilder) wrapHandler(h interface{}) http.HandlerFunc {
	switch h.(type) {
	case http.HandlerFunc:
		return h.(http.HandlerFunc)
	case func(w http.ResponseWriter, req *http.Request):
		return http.HandlerFunc(h.(func(w http.ResponseWriter, req *http.Request)))
	case Handler:
		h1 := h.(Handler)
		return func(w http.ResponseWriter, req *http.Request) {
			vars := mux.Vars(req)
			ctx := &Context{w, req, vars, route.app}
			res, err := h1(ctx)
			writeResult(w, res, err)
		}
	case func(*Context) interface{}, error:
		h2 := h.(func(*Context) (interface{}, error))
		return func(w http.ResponseWriter, req *http.Request) {
			vars := mux.Vars(req)
			ctx := &Context{w, req, vars, route.app}
			res, err := h2(ctx)
			writeResult(w, res, err)
		}
	default:
		panic("invalid handler")
	}
}

func writeResult(w http.ResponseWriter, res interface{}, err error) {
	var bytes []byte
	if err != nil {
		bytes, _ = json.Marshal(errorPayload{err.Error()})
	} else {
		bytes, _ = json.Marshal(res)
	}
	w.Write(bytes)
}
