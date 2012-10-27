package sample

import (
  "bytes"
  "encoding/json"
  "fmt"
  "html/template"
  "net/http"
)

type Name struct {
  First string
  Last string
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
  tpl, err := template.ParseFiles("templates/hello.html")

  if err != nil {
    fmt.Fprintf(writer, "Error parsing template: %v", err)
    return
  }

  var params Name
  nameJson := request.FormValue("nameJson")
  err = json.Unmarshal([]byte(nameJson), &params)
  if err != nil {
    fmt.Fprintf(writer, "Error parsing JSON: %v", err)
    return
  }

  buf := bytes.NewBuffer(make([]byte, 0, 0))
  err = tpl.Execute(buf, &params)
  if err != nil {
    fmt.Fprintf(writer, "Error executing template: %v", err)
    return
  }

  data := NewDataManager(request)
  err = data.store(params.First + " " + params.Last)
  if err != nil {
    fmt.Fprintf(writer, "Error storing entry: %v", err)
    return
  }

  writer.Write(buf.Bytes())
}
