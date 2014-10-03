package hypster

import (
  "net/http"
)

type Context struct {
  Response http.ResponseWriter
  Request *http.Request
}

// TODO WriteJson, ReadJson API

func (ctx *Context) WriteJson(value interface{}) {

}
