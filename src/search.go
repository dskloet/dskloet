package sample

import (
  "encoding/json"
  "fmt"
  "net/http"
)

func handleSearch(writer http.ResponseWriter, request *http.Request) {
  data := NewDataManager(request)
  name := request.URL.Path[len("/search/"):]

  entries, err := data.load(name)
  if err != nil {
    fmt.Fprintf(writer, "Error loading entries: %v", err)
    return
  }

  resultJson, err := json.Marshal(entries)
  if err != nil {
    fmt.Fprintf(writer, "Error marshalling JSON: %v", err)
    return
  }

  writer.Write(resultJson)
}
