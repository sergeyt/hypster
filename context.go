package hypster

import (
  "net/http"
  "encoding/json"
)

type Context struct {
  Response http.ResponseWriter
  Request *http.Request
}

type errorPayload struct {
  error string
}

// TODO ReadJson API

func (ctx *Context) WriteJson(v interface{}) {
  b, err := json.Marshal(v)
  if err != nil {
    b, _ = json.Marshal(errorPayload{err.Error()})
  }
  ctx.Response.Write(b)
}
