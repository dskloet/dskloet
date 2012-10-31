package sample

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type SearchHandler Handler

type SearchResponse struct {
  Entries []Entry
}

func (h *SearchHandler) handle() {
  var response SearchResponse

  if h.user != nil {
    name := h.httpRequest.URL.Path[len("/search/"):]
    entries, err := h.data.load(name)
    if err != nil {
      fmt.Fprintf(h.writer, "Error loading entries: %v", err)
      return
    }
    response.Entries = entries
  }

  resultJson, err := json.Marshal(response)
  if err != nil {
    fmt.Fprintf(h.writer, "Error marshalling JSON: %v", err)
    return
  }
  h.writer.Write(resultJson)
}

func handleSearch(writer http.ResponseWriter, request *http.Request) {
  handler := SearchHandler(*NewHandler(writer, request))
  handler.handle()
}
