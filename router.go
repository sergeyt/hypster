// Fluent API for gorilla mux
package hypster

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Router struct {
  impl *mux.Router
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

type Handler func(ctx *Context)

func NewRouter() *Router {
  return &Router{mux.NewRouter()}
}

func (r *Router) Route(url string) *RouteBuilder {
  rb := &RouteBuilder{}
  r.impl.HandleFunc(url, func(w http.ResponseWriter, req* http.Request) {
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
  return rb
}

// Registers GET handler
func (r *RouteBuilder) Get(h Handler) *RouteBuilder {
  r.get = h
  return r
}

// Registers POST handler
func (r *RouteBuilder) Post(h Handler) *RouteBuilder {
  r.post = h
  return r
}

// Registers PUT handler
func (r *RouteBuilder) Put(h Handler) *RouteBuilder {
  r.put = h
  return r
}

// Registers UPDATE handler
func (r *RouteBuilder) Update(h Handler) *RouteBuilder {
  r.update = h
  return r
}

// Registers DELETE handler
func (r *RouteBuilder) Delete(h Handler) *RouteBuilder {
  r.del = h
  return r
}
