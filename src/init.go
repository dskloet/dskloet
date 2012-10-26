package sample

import (
	"net/http"
)

func init() {
	http.HandleFunc("/hello", handleHello)
}
