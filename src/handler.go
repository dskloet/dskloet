package sample

import (
  "appengine"
  "appengine/user"
  "net/http"
)

type Handler struct {
  writer http.ResponseWriter
  httpRequest *http.Request
  err error
  context appengine.Context
  user *user.User
  data *DataManager
}

func NewHandler(writer http.ResponseWriter, request *http.Request) *Handler {
  context := appengine.NewContext(request)
  return &Handler{
    writer: writer,
    httpRequest: request,
    context: context,
    user: user.Current(context),
    data: NewDataManager(context),
  }
}
