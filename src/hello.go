package sample

import (
  "encoding/json"
  "fmt"
  "html/template"
  "net/http"
)

type HelloHandler Handler

type Name struct {
  First string
  Last  string
}

func (h *HelloHandler) handle() {
  tpl, err := template.ParseFiles("templates/hello.html")
  if err != nil {
    fmt.Fprintf(h.writer, "Error parsing template: %v", err)
    return
  }
  h.context.Debugf("Parsed template.")

  var params Name
  nameJson := h.httpRequest.FormValue("nameJson")
  err = json.Unmarshal([]byte(nameJson), &params)
  if err != nil {
    fmt.Fprintf(h.writer, "Error parsing JSON: %v", err)
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
