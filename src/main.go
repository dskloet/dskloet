package sample

import (
  "appengine/user"
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

  err = tpl.Execute(h.writer, &Params{ LoginUrl: loginUrl })
  if err != nil {
    fmt.Fprintf(h.writer, "Error executing template: %v", err)
    return
  }
}

func handleMain(writer http.ResponseWriter, request *http.Request) {
  handler := MainHandler(*NewHandler(writer, request))
  handler.handle()
}
