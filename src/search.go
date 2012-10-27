package sample

import (
  "bytes"
  "fmt"
  "html/template"
  "net/http"
  "time"
)

type SearchTemplateParameters struct {
  Times []time.Time
}

func handleSearch(writer http.ResponseWriter, request *http.Request) {
  tpl, err := template.ParseFiles("templates/search.html")
  if err != nil {
    fmt.Fprintf(writer, "Error parsing template: %v", err)
    return
  }

  data := NewDataManager(request)
  name := request.FormValue("name")
  times, err := data.load(name)
  if err != nil {
    fmt.Fprintf(writer, "Error loading entries: %v", err)
    return
  }

  params := SearchTemplateParameters{times}
  buf := bytes.NewBuffer(make([]byte, 0, 0))
  err = tpl.Execute(buf, &params)
  if err != nil {
    fmt.Fprintf(writer, "Error executing template: %v", err)
    return
  }

  writer.Write(buf.Bytes())
}
