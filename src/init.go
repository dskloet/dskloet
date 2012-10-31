package sample

import (
  "net/http"
)

func init() {
  http.HandleFunc("/", handleMain)
  http.HandleFunc("/hello", handleHello)
  http.HandleFunc("/search/", handleSearch)
}
