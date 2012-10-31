package sample

import (
  "appengine"
  "appengine/user"
  "net/http"
)

type Handler struct {
  writer http.ResponseWriter
  httpRequest *http.Request
  context appengine.Context
  user *user.User
  data *DataManager
}

func NewHandler(writer http.ResponseWriter, request *http.Request) *Handler {
  context := appengine.NewContext(request)
  return &Handler{
    writer,
    request,
    context,
    user.Current(context),
    NewDataManager(context),
  }
}
