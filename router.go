// Fluent API for gorilla mux
package hypster

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Router struct {
  impl *mux.Router
  routes map[string]*RouteBuilder
}

type RouteBuilder struct {
  // handlers
  get func(*Context)
  post func(*Context)
  put func(*Context)
  update func(*Context)
  patch func(*Context)
  del func(*Context)
}

type Handler func(*Context)

// Router API

func NewRouter() *Router {
  return &Router{mux.NewRouter(), make(map[string]*RouteBuilder)}
}

// implement http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req* http.Request) {
  r.impl.ServeHTTP(w, req)
}

func (r *Router) Route(pattern string) *RouteBuilder {
  rb := r.routes[pattern]

  if rb == nil {
    rb := &RouteBuilder{}
    r.routes[pattern] = rb

    r.impl.HandleFunc(pattern, func(w http.ResponseWriter, req* http.Request) {
      ctx := &Context{w,req}
      switch (req.Method) {
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

func (r *Router) Get(pattern string, handler Handler) *Router {
  r.Route(pattern).Get(handler)
  return r
}

func (r *Router) Post(pattern string, handler Handler) *Router {
  r.Route(pattern).Post(handler)
  return r
}

func (r *Router) Put(pattern string, handler Handler) *Router {
  r.Route(pattern).Put(handler)
  return r
}

func (r *Router) Update(pattern string, handler Handler) *Router {
  r.Route(pattern).Update(handler)
  return r
}

func (r *Router) Patch(pattern string, handler Handler) *Router {
  r.Route(pattern).Patch(handler)
  return r
}

func (r *Router) Delete(pattern string, handler Handler) *Router {
  r.Route(pattern).Delete(handler)
  return r
}

// RouteBuilder API

// Registers GET handler
func (r *RouteBuilder) Get(handler Handler) *RouteBuilder {
  r.get = handler
  return r
}

// Registers POST handler
func (r *RouteBuilder) Post(handler Handler) *RouteBuilder {
  r.post = handler
  return r
}

// Registers PUT handler
func (r *RouteBuilder) Put(handler Handler) *RouteBuilder {
  r.put = handler
  return r
}

// Registers UPDATE handler
func (r *RouteBuilder) Update(handler Handler) *RouteBuilder {
  r.update = handler
  return r
}

// Registers PATCH handler
func (r *RouteBuilder) Patch(handler Handler) *RouteBuilder {
  r.patch = handler
  return r
}

// Registers DELETE handler
func (r *RouteBuilder) Delete(handler Handler) *RouteBuilder {
  r.del = handler
  return r
}
