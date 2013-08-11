package sample

import (
  "appengine/blobstore"
  "encoding/json"
  "fmt"
  "html/template"
  "net/http"
)

type HelloHandler Handler

type Name struct {
  First string
  Last  string
  PhotoUrl string
}

func (h *HelloHandler) handle() {
  tpl, err := template.ParseFiles("templates/hello.html")
  if err != nil {
    fmt.Fprintf(h.writer, "Error parsing template: %v", err)
    return
  }
  h.context.Debugf("Parsed template.")

  blobs, other, err := blobstore.ParseUpload(h.httpRequest)
  if err != nil {
    fmt.Fprintf(h.writer, "Error parsing blobs: %v", err)
    return
  }

  var params Name
  nameJson := other.Get("nameJson")
  err = json.Unmarshal([]byte(nameJson), &params)

  if photoBlob, present := blobs["photo"]; present {
    params.PhotoUrl = "/photo/" + string(photoBlob[0].BlobKey)
  }

  if err != nil {
    fmt.Fprintf(h.writer, "Error parsing JSON: %v [%v]", err, nameJson)
    return
  }
  h.context.Debugf("Parsed JSON: %v", params)

  err = tpl.Execute(h.writer, &params)
  if err != nil {
    fmt.Fprintf(h.writer, "Error executing template: %v", err)
    return
  }
  h.context.Debugf("Executed template.")

  err = h.data.store(params.First + " " + params.Last)
  if err != nil {
    fmt.Fprintf(h.writer, "Error storing entry: %v", err)
    return
  }
  h.context.Debugf("Stored data.")
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
  handler := HelloHandler(*NewHandler(writer, request))
  handler.handle()
}
