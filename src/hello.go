package sample

import (
  "bytes"
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

  buf := bytes.NewBuffer(make([]byte, 0, 0))
  err = tpl.Execute(buf, &params)
  if err != nil {
    fmt.Fprintf(h.writer, "Error executing template: %v", err)
    return
  }
  h.context.Debugf("Executed template: %v", buf)

  err = h.data.store(params.First + " " + params.Last)
  if err != nil {
    fmt.Fprintf(h.writer, "Error storing entry: %v", err)
    return
  }
  h.context.Debugf("Stored data.")

  h.writer.Write(buf.Bytes())
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
  handler := HelloHandler(*NewHandler(writer, request))
  handler.handle()
}
