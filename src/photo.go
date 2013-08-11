package sample

import (
  "appengine"
  "appengine/blobstore"
  "net/http"
)

type PhotoHandler Handler

func (h *PhotoHandler) handle() {
  blobKey := h.httpRequest.URL.Path[len("/photo/"):]
  blobstore.Send(h.writer, appengine.BlobKey(blobKey))
}

func handlePhoto(writer http.ResponseWriter, request *http.Request) {
  handler := PhotoHandler(*NewHandler(writer, request))
  handler.handle()
}
