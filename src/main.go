package sample

import (
  "appengine/user"
  "bytes"
  "fmt"
  "html/template"
  "net/http"
)

type MainHandler Handler

func (h *MainHandler) handle() {
  tpl, err := template.ParseFiles("templates/index.html")
  if err != nil {
    fmt.Fprintf(h.writer, "Error parsing template: %v", err)
    return
  }

  var loginUrl string
  if h.user == nil {
    loginUrl, err = user.LoginURL(h.context, "/")
    if err != nil {
      fmt.Fprintf(h.writer, "Error generating login URL: %v", err)
      return
    }
  }

  type Params struct { LoginUrl string }
  buf := bytes.NewBuffer(make([]byte, 0, 0))

  err = tpl.Execute(buf, &Params{ LoginUrl: loginUrl })
  if err != nil {
    fmt.Fprintf(h.writer, "Error executing template: %v", err)
    return
  }

  h.writer.Write(buf.Bytes())
}

func handleMain(writer http.ResponseWriter, request *http.Request) {
  handler := MainHandler(*NewHandler(writer, request))
  handler.handle()
}
