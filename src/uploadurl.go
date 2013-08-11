package sample

import (
  "appengine/blobstore"
  "fmt"
  "net/http"
)

type UploadUrlHandler Handler

func (h *UploadUrlHandler) handle() {
  continueUrl := "/hello"
  blobUploadUrl, err := blobstore.UploadURL(h.context, continueUrl, nil)
  if err != nil {
    h.context.Errorf("Error making upload URL: %v", err)
    return
  }

  fmt.Fprintf(h.writer, "%v", blobUploadUrl)
}

func handleUploadUrl(writer http.ResponseWriter, request *http.Request) {
  handler := UploadUrlHandler(*NewHandler(writer, request))
  handler.handle()
}
