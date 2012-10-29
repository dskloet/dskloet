package sample

import (
  "appengine"
  "appengine/user"
  "encoding/json"
  "fmt"
  "net/http"
)

type SearchResponse struct {
  LoginUrl string
  Entries []Entry
}

func handleSearch(writer http.ResponseWriter, request *http.Request) {
  context := appengine.NewContext(request)
  data := NewDataManager(context)
  u := user.Current(context)

  var response SearchResponse

  if u == nil {
    response.LoginUrl, _ = user.LoginURL(context, "/thanks.html")
  } else {
    name := request.URL.Path[len("/search/"):]
    entries, err := data.load(name)
    if err != nil {
      fmt.Fprintf(writer, "Error loading entries: %v", err)
      return
    }
    response.Entries = entries
  }

  resultJson, err := json.Marshal(response)
  if err != nil {
    fmt.Fprintf(writer, "Error marshalling JSON: %v", err)
    return
  }
  writer.Write(resultJson)
}
