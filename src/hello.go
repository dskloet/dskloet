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

func (h *HelloHandler) parseTemplate() (tpl *template.Template) {
  if (h.err != nil) { return }

  tpl, h.err = template.ParseFiles("templates/hello.html")
  if h.err != nil {
    fmt.Fprintf(h.writer, "Error parsing template: %v", h.err)
  } else {
    h.context.Debugf("Parsed template.")
  }
  return
}

func (h *HelloHandler) parseNameJson() (params Name) {
  if (h.err != nil) { return }

  nameJson := h.httpRequest.FormValue("nameJson")
  h.err = json.Unmarshal([]byte(nameJson), &params)
  if h.err != nil {
    fmt.Fprintf(h.writer, "Error parsing JSON: %v", h.err)
  } else {
    h.context.Debugf("Parsed JSON: %v", params)
  }
  return
}

func (h *HelloHandler) executeTemplate(tpl *template.Template, params Name) {
  if (h.err != nil) { return }

  h.err = tpl.Execute(h.writer, &params)
  if h.err != nil {
    fmt.Fprintf(h.writer, "Error executing template: %v", h.err)
  } else {
    h.context.Debugf("Executed template.")
  }
  return
}

func (h *HelloHandler) storeName(params Name) {
  if (h.err != nil) { return }

  h.err = h.data.store(params.First + " " + params.Last)
  if h.err != nil {
    fmt.Fprintf(h.writer, "Error storing entry: %v", h.err)
  } else {
    h.context.Debugf("Stored data.")
  }
  return
}

func (h *HelloHandler) handle() {
  tpl := h.parseTemplate()
  params := h.parseNameJson()
  h.executeTemplate(tpl, params)
  h.storeName(params)
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
  handler := HelloHandler(*NewHandler(writer, request))
  handler.handle()
}
