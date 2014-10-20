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
func (route *RouteBuilder) Head(handler Handler) *RouteBuilder {
	return route.register("HEAD", handler)
}

// Options registers OPTIONS handler
func (route *RouteBuilder) Options(handler Handler) *RouteBuilder {
	return route.register("OPTIONS", handler)
}

// Get registers GET handler
func (route *RouteBuilder) Get(handler Handler) *RouteBuilder {
	return route.register("GET", handler)
}

// Post registers POST handler
func (route *RouteBuilder) Post(handler Handler) *RouteBuilder {
	return route.register("POST", handler)
}

// Put registers PUT handler
func (route *RouteBuilder) Put(handler Handler) *RouteBuilder {
	return route.register("PUT", handler)
}

// Update registers UPDATE handler
func (route *RouteBuilder) Update(handler Handler) *RouteBuilder {
	return route.register("UPDATE", handler)
}

// Patch registers PATCH handler
func (route *RouteBuilder) Patch(handler Handler) *RouteBuilder {
	return route.register("PATCH", handler)
}

// Delete registers DELETE handler
func (route *RouteBuilder) Delete(handler Handler) *RouteBuilder {
	return route.register("DELETE", handler)
}

func (route *RouteBuilder) register(verb string, handler Handler) *RouteBuilder {
	route.app.router.HandleFunc(route.pattern, route.wrapHandler(handler)).Methods(verb)
	return route
}

func (route *RouteBuilder) wrapHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		ctx := &Context{w, req, vars, route.app}
		res, err := handler(ctx)

		var bytes []byte
		if err != nil {
			bytes, _ = json.Marshal(errorPayload{err.Error()})
		} else {
			bytes, _ = json.Marshal(res)
		}

		w.Write(bytes)
	}
}
